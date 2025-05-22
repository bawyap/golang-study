package main

import (
	"fmt"
	"math/big"
)

type Account struct {
	AccountId      string
	AccountHolder  User
	AccountNumber  string
	AccountBalance AccountBalance
	TransHistory   []Transaction
}

func NewAccount(accountId string, accountHolder User, accountNumber string, accountBalance AccountBalance) Account {
	return Account{AccountId: accountId, AccountNumber: accountNumber, AccountHolder: accountHolder, AccountBalance: accountBalance}
}

func (account *Account) Credit(amount *big.Float) {
	accountBalance := account.AccountBalance
	balanceAmount := accountBalance.amount.Add(accountBalance.amount, amount)
	fmt.Printf("Deposited Amount: %s and CurrentBalance: %s\n", amount.String(), balanceAmount.String())
	account.updateTransHistory(CREDIT)
	fmt.Printf(" AccountId: %s and CurrentBalance: %s\n", account.AccountId, balanceAmount.String())

}

func (account *Account) getCurrentBalance() *big.Float {
	accountBalance := account.AccountBalance
	return accountBalance.amount
}

func (account *Account) Debit(amount *big.Float) {
	accountBalance := account.AccountBalance
	balanceAmount := accountBalance.amount.Sub(accountBalance.amount, amount)
	fmt.Printf("Debited Amount: %s and CurrentBalance: %s\n", amount.String(), balanceAmount.String())
	fmt.Printf(" AccountId: %s and CurrentBalance: %s\n", account.AccountId, balanceAmount.String())
}

func (account *Account) updateTransHistory(transType TransactionType) {
}
