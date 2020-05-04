package main

import (
	"fmt"
	"github.com/DesignPatterns/structural/adapter/adapters"
)

type PaymentMethod interface {
	Send(fromEmail string, toEmail string, amount int) error
}

type Account struct {
	Email string
	Method PaymentMethod
}

func main() {
	account1 := &Account{
		Email:  "unishubh1@gmail.com",
		Method: &adapters.BankAdapter{},
	}

	account2 := &Account{
		Email:  "smartscribs@gmail.com",
		Method: &adapters.Paypal_Adapter{},
	}

	err := account1.Pay(account2, 10)
	if err!= nil {
		fmt.Println("An error occurred")
	}
}

func (a *Account) Pay(receiver *Account,amount int) error {
	fromEmail := a.Email
	toEmail := receiver.Email

	return a.Method.Send(fromEmail,toEmail,amount)
}