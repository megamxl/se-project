package converter

type Request struct {
	GivenCurrency  string
	Amount         float64
	TargetCurrency string
}

type Response struct {
	Currency string
	Amount   float64
}

type Converter interface {
	Convert(request Request) (Response, error)
	GetAvailableCurrency() ([]string, error)
}
