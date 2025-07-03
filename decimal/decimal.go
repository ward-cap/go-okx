package decimal

import (
	orig "github.com/shopspring/decimal"
)

type NullDecimalV2 orig.NullDecimal

//goland:noinspection GoMixedReceiverTypes
func (d *NullDecimalV2) UnmarshalJSON(decimalBytes []byte) error {
	if string(decimalBytes) == "null" || string(decimalBytes) == `""` {
		d.Valid = false
		return nil
	}
	d.Valid = true
	return d.Decimal.UnmarshalJSON(decimalBytes)
}

func (d NullDecimalV2) MarshalJSON() ([]byte, error) {
	return d.N().MarshalJSON()
}

func (d NullDecimalV2) MarshalText() ([]byte, error) {
	return d.N().MarshalText()
}

//goland:noinspection GoMixedReceiverTypes
func (d NullDecimalV2) N() orig.NullDecimal {
	return orig.NullDecimal(d)
}
