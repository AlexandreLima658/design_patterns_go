package gateway

import "fmt"

type BrasPag struct {
	Price        float64
	Installments int64
	Card         string
	CCV          string
}

func (b *BrasPag) Pay() {
	fmt.Println("Pay with BrasPag")
	fmt.Println("Amount: ", b.Price)
	fmt.Println("Installments: ", b.Installments)
	fmt.Println("Card Number: ", b.Card)
	fmt.Println("CCV: ", b.CCV)
}
