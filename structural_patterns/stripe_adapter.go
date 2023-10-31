package structuralpatterns

import "github.com/AlexandreLima658/design_patterns_go/structural_patterns/gateway"

type StripeAdapter struct{}

func (s StripeAdapter) Pay(props PaymentProperties) {
	stripe := gateway.Stripe{}

	stripe.Pay(
		props.Value,
		props.Installments,
		props.CardNumber,
	)
}
