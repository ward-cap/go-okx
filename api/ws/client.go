package ws

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/ward-cap/go-okx"
	"github.com/ward-cap/go-okx/events"
	"go.uber.org/zap"
)

// ClientWs is the websocket api client
//
// https://www.okex.com/docs-v5/en/#websocket-api
type ClientWs struct {
	context.Context

	Cancel              context.CancelFunc
	StructuredEventChan chan interface{}
	RawEventChan        chan []byte
	ErrWsChan           chan error
	ErrEventChan        chan *events.Error
	SubscribeChan       chan *events.Subscribe
	UnsubscribeCh       chan *events.Unsubscribe
	LoginChan           chan *events.Login
	SuccessChan         chan *events.Success
	sendChan            map[bool]chan []byte
	url                 map[bool]okex.BaseURL
	conn                map[bool]*websocket.Conn
	apiKey              string
	secretKey           []byte
	passphrase          string
	lastTransmit        map[bool]*time.Time
	mu                  map[bool]*sync.RWMutex
	AuthRequested       *time.Time
	Authorized          bool
	Private             *Private
	Public              *Public
	Trade               *Trade
	logger              *zap.SugaredLogger
}

const (
	redialTick = 2 * time.Second
	writeWait  = 3 * time.Second
	pongWait   = 30 * time.Second
	PingPeriod = pongWait * 8 / 10
)

// NewClient returns a pointer to a fresh ClientWs
func NewClient(
	ctx context.Context,
	apiKey, secretKey, passphrase string,
	url map[bool]okex.BaseURL,
	logger *zap.SugaredLogger,
) *ClientWs {
	ctx, cancel := context.WithCancel(ctx)
	c := &ClientWs{
		Context: ctx,
		Cancel:  cancel,

		logger:       logger,
		apiKey:       apiKey,
		secretKey:    []byte(secretKey),
		passphrase:   passphrase,
		url:          url,
		sendChan:     map[bool]chan []byte{true: make(chan []byte, 3), false: make(chan []byte, 3)},
		conn:         make(map[bool]*websocket.Conn),
		lastTransmit: make(map[bool]*time.Time),
		mu:           map[bool]*sync.RWMutex{true: {}, false: {}},
	}
	c.Private = NewPrivate(c)
	c.Public = NewPublic(c)
	c.Trade = NewTrade(c)
	return c
}

// Connect into the server
//
// https://www.okex.com/docs-v5/en/#websocket-api-connect
func (c *ClientWs) Connect(p bool) error {
	if c.conn[p] != nil {
		return nil
	}
	return c.dial(p)

}

//func (c *ClientWs) Context() context.Context {
//	return c.ctx
//}

// Login
//
// https://www.okex.com/docs-v5/en/#websocket-api-login
func (c *ClientWs) Login() error {
	if c.Authorized {
		return nil
	}
	if c.AuthRequested != nil && time.Since(*c.AuthRequested).Seconds() < 30 {
		return nil
	}
	now := time.Now()
	c.AuthRequested = &now
	method := http.MethodGet
	path := "/users/self/verify"
	ts, sign := c.sign(method, path)
	args := []map[string]string{
		{
			"apiKey":     c.apiKey,
			"passphrase": c.passphrase,
			"timestamp":  ts,
			"sign":       sign,
		},
	}
	return c.Send(true, okex.LoginOperation, args)
}

// Subscribe
// Users can choose to subscribe to one or more channels, and the total length of multiple channels cannot exceed 4096 bytes.
//
// https://www.okex.com/docs-v5/en/#websocket-api-subscribe
func (c *ClientWs) Subscribe(p bool, ch []okex.ChannelName, args ...map[string]string) error {
	chCount := max(len(ch), 1)
	tmpArgs := make([]map[string]string, chCount*len(args))

	n := 0
	for i := 0; i < chCount; i++ {
		for _, arg := range args {
			tmpArgs[n] = make(map[string]string)
			for k, v := range arg {
				tmpArgs[n][k] = v
			}
			if len(ch) > 0 {
				tmpArgs[n]["channel"] = string(ch[i])
			}
			n++
		}
	}

	return c.Send(p, okex.SubscribeOperation, tmpArgs)
}

// Unsubscribe into channel(s)
//
// https://www.okex.com/docs-v5/en/#websocket-api-unsubscribe
func (c *ClientWs) Unsubscribe(p bool, ch []okex.ChannelName, args map[string]string) error {
	tmpArgs := make([]map[string]string, len(ch))
	for i, name := range ch {
		tmpArgs[i] = make(map[string]string)
		tmpArgs[i]["channel"] = string(name)
		for k, v := range args {
			tmpArgs[i][k] = v
		}
	}
	return c.Send(p, okex.UnsubscribeOperation, tmpArgs)
}

// Send message through either connections
func (c *ClientWs) Send(p bool, op okex.Operation, args []map[string]string, extras ...map[string]string) error {
	if op != okex.LoginOperation {
		err := c.Connect(p)
		if err == nil {
			if p {
				err = c.WaitForAuthorization()
				if err != nil {
					return err
				}
			}
		} else {
			return err
		}
	}

	data := map[string]interface{}{
		"op":   op,
		"args": args,
	}
	for _, extra := range extras {
		for k, v := range extra {
			data[k] = v
		}
	}
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	c.sendChan[p] <- j
	return nil
}

// WaitForAuthorization waits for the auth response and try to log in if it was needed
func (c *ClientWs) WaitForAuthorization() error {
	if c.Authorized {
		return nil
	}
	if err := c.Login(); err != nil {
		return err
	}
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for range ticker.C {
		if c.Authorized {
			return nil
		}
	}
	return nil
}

func (c *ClientWs) dial(p bool) error {
	c.mu[p].Lock()
	conn, _, err := websocket.Dial(c, string(c.url[p]), nil)
	if err != nil {
		c.mu[p].Unlock()
		return fmt.Errorf("dial error: %w", err)
	}
	conn.SetReadLimit(32768 * 10) // Increase read limit for large order books if necessary
	c.conn[p] = conn
	c.mu[p].Unlock()

	go c.receiver(p)
	go c.sender(p)

	return nil
}

func (c *ClientWs) sender(p bool) {
	ticker := time.NewTicker(time.Millisecond * 300)
	defer ticker.Stop()
	for {
		select {
		case data := <-c.sendChan[p]:
			c.mu[p].RLock()
			conn := c.conn[p]
			if conn == nil {
				c.mu[p].RUnlock()
				if c.logger != nil {
					c.logger.Warn("connection is nil")
				}
			}

			writeCtx, cancel := context.WithTimeout(c, writeWait)
			err := conn.Write(writeCtx, websocket.MessageText, data)
			cancel()

			if err != nil {
				c.mu[p].RUnlock()
				if c.logger != nil {
					c.logger.Error(err)
				}
				continue
			}
			now := time.Now()
			c.lastTransmit[p] = &now
			c.mu[p].RUnlock()
		case <-ticker.C:
			c.mu[p].RLock()
			conn := c.conn[p]
			lastTransmit := c.lastTransmit[p]
			c.mu[p].RUnlock()
			if conn != nil && (lastTransmit == nil || (lastTransmit != nil && time.Since(*lastTransmit) > PingPeriod)) {
				// Using select to prevent blocking if sendChan is full
				select {
				case c.sendChan[p] <- []byte("ping"):
					if c.logger != nil {
						c.logger.Info("send ping")
					}
				default:
					if c.logger != nil {
						c.logger.Info("can't send ping")
					}
				}
			}
		case <-c.Done():
			if c.logger != nil {
				c.logger.Warn("connection is closed")
			}
			return
		}
	}
}

func (c *ClientWs) receiver(p bool) {
	defer func() {
		c.mu[p].Lock()
		if c.conn[p] != nil {
			// Close with normal status, error handling is done by caller/logger
			_ = c.conn[p].Close(websocket.StatusNormalClosure, "receiver closing")
			c.conn[p] = nil
		}
		c.mu[p].Unlock()
	}()

	for {
		c.mu[p].RLock()
		conn := c.conn[p]
		c.mu[p].RUnlock()
		if conn == nil {
			if c.logger != nil {
				c.logger.Warnf("connection is nil")
			}
			return
		}

		// Emulate SetReadDeadline using context timeout
		readCtx, cancel := context.WithTimeout(c, pongWait)
		mt, data, err := conn.Read(readCtx)
		cancel()

		if err != nil {
			if e := c.ErrWsChan; e != nil {
				e <- err
			}
			c.Cancel()
			break
		}

		now := time.Now()
		c.mu[p].Lock()
		c.lastTransmit[p] = &now
		c.mu[p].Unlock()

		if mt == websocket.MessageText {
			switch string(data) {
			case "pong":
				if c.logger != nil {
					c.logger.Info("got pong")
				}
				continue

			case "ping":
				if c.logger != nil {
					c.logger.Info("got ping")
				}
				go func() {
					c.sendChan[p] <- []byte("pong")
				}()
				continue
			}

			e := &events.Basic{}
			if err := json.Unmarshal(data, &e); err != nil {
				// Not a structured event; forward raw and continue instead of dropping the connection.
				if c.RawEventChan != nil {
					c.RawEventChan <- data
				}
				continue
			}

			go c.process(data, e)
		}
	}
}

func (c *ClientWs) sign(method, path string) (ts, signature string) {
	ts = strconv.FormatInt(time.Now().Unix(), 10)

	mac := hmac.New(sha256.New, c.secretKey)
	mac.Write([]byte(ts + method + path))

	signature = base64.StdEncoding.EncodeToString(mac.Sum(nil))
	return
}

func (c *ClientWs) process(data []byte, e *events.Basic) bool {
	switch e.Event {
	case "error":
		e := events.Error{}
		_ = json.Unmarshal(data, &e)
		if c.ErrEventChan != nil {
			c.ErrEventChan <- &e
		}
		return true
	case "subscribe":
		e := events.Subscribe{}
		_ = json.Unmarshal(data, &e)
		if c.SubscribeChan != nil {
			c.SubscribeChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	case "unsubscribe":
		e := events.Unsubscribe{}
		_ = json.Unmarshal(data, &e)
		if c.UnsubscribeCh != nil {
			c.UnsubscribeCh <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	case "login":
		if time.Since(*c.AuthRequested).Seconds() > 30 {
			c.AuthRequested = nil
			_ = c.Login()
			break
		}
		c.Authorized = true
		e := events.Login{}
		_ = json.Unmarshal(data, &e)
		if c.LoginChan != nil {
			c.LoginChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	}
	if c.Private.Process(data, e) {
		return true
	}
	if c.Public.Process(data, e) {
		return true
	}
	if e.ID != "" {
		if e.Code != 0 {
			ee := *e
			ee.Event = "error"
			return c.process(data, &ee)
		}
		e := events.Success{}
		_ = json.Unmarshal(data, &e)
		if c.SuccessChan != nil {
			c.SuccessChan <- &e
		}
		if c.StructuredEventChan != nil {
			c.StructuredEventChan <- e
		}
		return true
	}
	if c.RawEventChan != nil {
		c.RawEventChan <- data
	}

	return false
}
