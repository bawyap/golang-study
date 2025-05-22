package main

import (
	"fmt"
	"math/big"
	"time"
)

type DigitalWallet struct {
	Users              map[string]digitalWalletSystem.User
	Accounts           map[string]digitalWalletSystem.Account
	WalletTransHistory map[string]digitalWalletSystem.Transaction
}

func NewDigitalWallet(users map[string]digitalWalletSystem.User, accounts map[string]digitalWalletSystem.Account, transactionHistory map[string]digitalWalletSystem.Transaction) *DigitalWallet {
	return &DigitalWallet{Users: users, Accounts: accounts, WalletTransHistory: transactionHistory}

}

func (wallet *DigitalWallet) CreateUser(user digitalWalletSystem.User) error {
	if userVal, exists := wallet.Users[user.UserId]; exists {
		return fmt.Errorf("user with ID %s already exists", userVal.UserId)
	}
	wallet.Users[user.UserId] = user
	return nil
}

func (wallet *DigitalWallet) CreateAccount(account digitalWalletSystem.Account) error {
	if accountVal, exists := wallet.Accounts[account.AccountId]; exists {
		return fmt.Errorf("user with account ID %s already exists", accountVal.AccountId)
	}
	wallet.Accounts[account.AccountId] = account
	return nil
}

func (wallet *DigitalWallet) TransferFund(debitAccount digitalWalletSystem.Account, creditAccount digitalWalletSystem.Account, amount *big.Float, currency digitalWalletSystem.Currency) error {
	canDebitAmount, err := HasSufficientBalance(debitAccount, amount)
	if canDebitAmount {
		debitAccount.Debit(amount)
		creditAccount.Credit(amount)
		newTransaction := wallet.createTransaction(debitAccount, creditAccount, amount, currency)
		trnsErr := wallet.addTransactionToTransactionHistory(newTransaction)
		if trnsErr != nil {
			return trnsErr
		}
		return nil
	} else {
		return err
	}
}

func HasSufficientBalance(debitAccount digitalWalletSystem.Account, amount *big.Float) (bool, error) {
	debitAccBalance := debitAccount.getCurrentBalance()
	if amount.Cmp(debitAccBalance) == -1 || amount.Cmp(debitAccBalance) == 0 {
		return true, nil
	} else {
		return false, fmt.Errorf("Transaction failed: Insufficient fund %s", debitAccount.AccountId)
	}
}

func (wallet *DigitalWallet) createTransaction(debitAccount digitalWalletSystem.Account, creditAccount digitalWalletSystem.Account, amount *big.Float, currency digitalWalletSystem.Currency) digitalWalletSystem.Transaction {
	transactionNumber := getTransactionNumber()
	newTransaction := digitalWalletSystem.Transaction{
		TrnsId:     transactionNumber,
		SenderId:   debitAccount.AccountId,
		ReceiverId: creditAccount.AccountId,
		Amount:     amount,
		Currency:   currency,
		TrnsSts:    digitalWalletSystem.SUCCESS,
		TimeStamp:  time.Now(),
	}
	debitAccount.Credit(amount)
	fmt.Printf("Transaction successful: %+v\n", newTransaction.TrnsId)
	return newTransaction
}

func (wallet *DigitalWallet) addTransactionToTransactionHistory(newTransaction digitalWalletSystem.Transaction) error {
	if _, exists := wallet.WalletTransHistory[newTransaction.TrnsId]; exists {
		return fmt.Errorf("Transaction Id already exists", newTransaction.TrnsId)
	}
	wallet.WalletTransHistory[newTransaction.TrnsId] = newTransaction
	return nil
}

func getTransactionNumber() string {
	count := 0
	return fmt.Sprintf("%05d", count+1)
}

func main() {
	digitalWallet := NewDigitalWallet(make(map[string]digitalWalletSystem.User), make(map[string]digitalWalletSystem.Account), make(map[string]digitalWalletSystem.Transaction))

	// Create users
	user1 := digitalWalletSystem.NewUser("U001", "John Doe", "john@example.com", "password123")
	user2 := digitalWalletSystem.NewUser("U002", "Jane Smith", "jane@example.com", "password456")
	digitalWallet.CreateUser(user1)
	digitalWallet.CreateUser(user2)

	// Create accounts
	account1 := digitalWalletSystem.NewAccount("A001", user1, "1234567890", openAccBalance(digitalWalletSystem.USD))
	account2 := digitalWalletSystem.NewAccount("A002", user2, "9876543210", openAccBalance(digitalWalletSystem.USD))
	digitalWallet.CreateAccount(account1)
	digitalWallet.CreateAccount(account2)

	// Credit funds
	account1.Credit(big.NewFloat(1000.00))
	account2.Credit(big.NewFloat(500.00))
	account1.Credit(big.NewFloat(1000.00))

	// Transfer funds
	amount := big.NewFloat(100.00)

	digitalWallet.TransferFund(account1, account2, amount, digitalWalletSystem.USD)
}

func openAccBalance(currency digitalWalletSystem.Currency) digitalWalletSystem.AccountBalance {
	return digitalWalletSystem.NewAccountBalance(big.NewFloat(0.0), currency)
}
