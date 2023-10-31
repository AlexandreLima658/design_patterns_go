package gateway

import "fmt"

type Stripe struct{}

func (s Stripe) Pay(value float64, installments int64, cardNumber string) {
	fmt.Println("Pay with Stripe")
	fmt.Println("Amount: ", value)
	fmt.Println("Installments: ", installments)
	fmt.Println("CardNumber: ", cardNumber)

}
