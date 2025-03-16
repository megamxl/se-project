package soap

import (
	"errors"
	"github.com/hooklift/gowsdl/soap"
	int "github.com/megamxl/se-project/Rental-Server/internal/communication/converter"
	soapGen "github.com/megamxl/se-project/Rental-Server/internal/communication/converter/soap/myservice"
	"log/slog"
)

type Service struct {
	client soapGen.ConversionPort
}

func NewSoapService(url string) *Service {

	client := soap.NewClient(url)

	return &Service{
		client: soapGen.NewConversionPort(client),
	}
}

var _ int.Converter = (*Service)(nil)

func (s Service) GetAvailableCurrency() ([]string, error) {
	//TODO implement me
	/*	_ :=
		&soapGen.GetAvailableCurrencyRequest{
			SenselessRequestPayload: &soapGen.SenselessRequestPayload{
				DontFill: "",
			},
		}*/
	return nil, errors.New("Not implemented")

}

func (s Service) Convert(request int.Request) (int.Response, error) {

	required := soapGen.Currency(request.TargetCurrency)
	given := soapGen.Currency(request.GivenCurrency)

	myVar := &soapGen.ConversionRequest{
		ConversionRequestPayload: &soapGen.ConversionRequestPayload{
			GivenCurrency:    &given, // Directly use &CurrencyEUR
			Amount:           request.Amount,
			RequiredCurrency: &required,
		},
	}

	resp, err := s.client.Conversion(myVar)
	if err != nil {
		slog.Error("Error calling SOAP service Conversion: ", err)
		return int.Response{}, err
	}

	return int.Response{
		Currency: string(*resp.ConversionResponsePayload.ConvertedCurrency),
		Amount:   resp.ConversionResponsePayload.Amount,
	}, nil
}
