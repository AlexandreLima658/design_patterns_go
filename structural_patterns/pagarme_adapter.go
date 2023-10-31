package structuralpatterns


import "github.com/AlexandreLima658/design_patterns_go/structural_patterns/gateway"

type PagarmeAdapter struct{}

func (p PagarmeAdapter) Pay(props PaymentProperties) {
	pagarme := gateway.Pagarme{
		Amount:       props.Value,
		Installments: props.Installments,
		CardNumber:   props.CardNumber,
		SecurityCode: props.CCV,
	}
	pagarme.Pay()
}
