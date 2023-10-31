package structuralpatterns

import "github.com/AlexandreLima658/design_patterns_go/structural_patterns/gateway"

type BrasPagAdapter struct{}

func (b BrasPagAdapter) Pay(props PaymentProperties) {
	braspag := gateway.BrasPag{
		Price:        props.Value,
		Installments: props.Installments,
		Card:         props.CardNumber,
		CCV:          props.CCV,
	}
	
	braspag.Pay()
}
