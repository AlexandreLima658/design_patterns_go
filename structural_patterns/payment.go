package structuralpatterns

// Target interface
type PaymentProperties struct {
	Value        float64
	Installments int64
	CardNumber   string
	CCV          string
}

type Payment interface {
	Pay(props PaymentProperties)
}
