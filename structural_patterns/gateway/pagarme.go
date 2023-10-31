package gateway

import "fmt"

type Pagarme struct {
	Amount       float64
	Installments int64
	CardNumber   string
	SecurityCode string
}

func (p *Pagarme) Pay() {
	fmt.Println("Pay with Pagarme")
	fmt.Println("Amount: ", p.Amount)
	fmt.Println("Installments: ", p.Installments)
	fmt.Println("CardNumber: ", p.CardNumber)
	fmt.Println("SecurityCode: ", p.SecurityCode)
}
