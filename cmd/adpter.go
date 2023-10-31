package main

import structuralpatterns "github.com/AlexandreLima658/design_patterns_go/structural_patterns"

func main() {
	props := structuralpatterns.PaymentProperties{
		Value:        1000,
		Installments: 10,
		CardNumber:   "123 456 789 123",
		CCV:          "123",
	}

	// payWith(props, structuralpatterns.BrasPagAdapter{})
	payWith(props, structuralpatterns.PagarmeAdapter{})
	// payWith(props, structuralpatterns.StripeAdapter{})

}

func payWith(props structuralpatterns.PaymentProperties, gateway structuralpatterns.Payment) {
	gateway.Pay(props)
}
