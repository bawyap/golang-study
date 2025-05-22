package main

import "math/big"

type AccountBalance struct {
	amount   *big.Float
	Currency Currency
}

func NewAccountBalance(amount *big.Float, currency Currency) AccountBalance {
	return AccountBalance{amount: amount, Currency: currency}
}
