package utils

import "math/big"

func FromWei(value *big.Int) *big.Int {
	tmp := big.NewInt(1)
	tmp.SetString("1000000000000000000", 10)
	return value.Div(value, tmp)
}
