package adapters

import "github.com/DesignPatterns/structural/adapter/paypal"

type Paypal_Adapter struct {}

func (p *Paypal_Adapter) Send(fromEmail string, toEmail string, amount int) error {
	fromMobile := findMobileByEmail(fromEmail)
	toMobile := findMobileByEmail(toEmail)
	return paypal.Send(fromMobile,toMobile,amount)
}
func findMobileByEmail(email string) string{
	//This will query the data store to return the Account Number
	return "9877896547" // a random phone number
}
