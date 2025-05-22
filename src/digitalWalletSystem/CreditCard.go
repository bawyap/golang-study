package main

type CardType string

const (
	MASTER_CARD      CardType = "MASTER_CARD"
	VISA             CardType = "VISA"
	AMERICAN_EXPRESS CardType = "AMERICAN_EXPRESS"
)

type CreditCard struct {
	CardNumber     string
	CardHolder     User
	Cvv            int
	ExpirationDate string
	CardType       CardType
}
