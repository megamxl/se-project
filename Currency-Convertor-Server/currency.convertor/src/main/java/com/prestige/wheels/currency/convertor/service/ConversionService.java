package com.prestige.wheels.currency.convertor.service;

import com.prestige.wheels.currency.convertor.repository.CurrencyRatesRepository;
import com.prestige.wheels.currency.convertor.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.convertor.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.convertor.soap.model.Currency;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;
import java.math.RoundingMode;

@Service
@RequiredArgsConstructor
public class ConversionService {

    private final CurrencyRatesRepository currencyRatesRepository;

    public ConversionResponsePayload convert (ConversionRequestPayload payload) {

        assert payload != null;
        assert payload.getGivenCurrency() != null;
        assert payload.getRequiredCurrency() != null;

        //TODO assert payload fields
        if (payload.getGivenCurrency().equals(payload.getRequiredCurrency())) {
            return createResponsePayload(payload,payload.getAmount());
        }

        // first of all check if is not EUR convert it to EUR
        double amountInEuro = convertToEuro(payload.getGivenCurrency(),  payload.getAmount());

        if ((payload.getRequiredCurrency().equals(Currency.EUR))) {
            return createResponsePayload(payload,amountInEuro);
        }

        // afterward convert from euro to what is needed
        double converted = convertFromEuroToTargetCurrency(payload.getRequiredCurrency(), amountInEuro);

        return createResponsePayload(payload, converted);
    }

    private static ConversionResponsePayload createResponsePayload(ConversionRequestPayload payload, double converted) {
        //TODO Apply rounding here
        ConversionResponsePayload conversionResponsePayload = new ConversionResponsePayload();
        conversionResponsePayload.setConvertedCurrency(payload.getRequiredCurrency());
        conversionResponsePayload.setAmount(converted);
        return conversionResponsePayload;
    }

    private double convertToEuro(Currency currency, double amount) {

        if (currency == Currency.EUR){
            return amount;
        }

        return amount /getRateForCurrency(currency);
    }


    private double convertFromEuroToTargetCurrency(Currency currency, double amount) {

        return amount * getRateForCurrency(currency);
    }


    private double getRateForCurrency(Currency currency) {
        return currencyRatesRepository.getRateByCurrency(currency).orElseThrow(() -> new RuntimeException("Currency not found in the local repository"));
    }
}
