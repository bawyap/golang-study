package main

import (
	"math/big"
	"time"
)

type TransactionStatus string

type TransactionType string

const (
	INITIATED TransactionStatus = "Initiated"
	PENDING   TransactionStatus = "Pending"
	SUCCESS   TransactionStatus = "Success"
	FAILED    TransactionStatus = "Failed"
	REFUNDED  TransactionStatus = "Refunded"
)

const (
	CREDIT       TransactionType = "Credit"
	DEBIT        TransactionType = "Debit"
	ATM_WITHDRAW TransactionType = "ATM_WITHDRAW"
)

type Transaction struct {
	TrnsId     string
	SenderId   string
	ReceiverId string
	Amount     *big.Float
	Currency   Currency
	TrnsSts    TransactionStatus
	TimeStamp  time.Time
}
