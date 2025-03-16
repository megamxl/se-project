package converter

type Request struct {
	givenCurrency  string
	amount         float64
	targetCurrency string
}

type Response struct {
	currency string
	amount   float64
}

type Converter interface {
	convert(request Request) Response
}
