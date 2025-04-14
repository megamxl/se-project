package grpc

import (
	"context"
	"errors"
	"github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	grpc "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/grpc/proto"
)

type Converter struct {
	conn grpc.ConvertorClient
}

func (c Converter) Convert(request converter.Request) (converter.Response, error) {

	give, ok := grpc.Currency_value[request.GivenCurrency]
	target, ok1 := grpc.Currency_value[request.TargetCurrency]

	if !ok || !ok1 {
		return converter.Response{}, errors.New("invalid currency")
	}

	conversionRequest := grpc.ConversionRequest{
		Given:  grpc.Currency(give),
		Amount: request.Amount,
		Target: grpc.Currency(target),
	}

	convert, err := c.conn.Convert(context.Background(), &conversionRequest)
	if err != nil {
		return converter.Response{}, err
	}

	if convert != nil {
		return converter.Response{
			Currency: convert.Converted.String(),
			Amount:   convert.Amount,
		}, nil
	}

	return converter.Response{}, errors.New("invalid conversion")
}

func (c Converter) GetAvailableCurrency() ([]string, error) {

	currency, err := c.conn.GetAvailableCurrency(context.Background(), &grpc.Empty{})
	if err != nil {
		return nil, err
	}

	var currencies []string

	for _, curr := range currency.Currencies {
		currencies = append(currencies, curr.String())
	}

	return currencies, nil
}

func NewConverter(conn grpc.ConvertorClient) *Converter {
	return &Converter{conn}
}

var _ converter.Converter = (*Converter)(nil)
