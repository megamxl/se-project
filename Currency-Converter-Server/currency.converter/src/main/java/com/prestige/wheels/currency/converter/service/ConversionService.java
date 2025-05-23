package com.prestige.wheels.currency.converter.service;

import com.prestige.wheels.currency.converter.repository.CurrencyRatesRepository;
import com.prestige.wheels.currency.converter.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.converter.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.converter.soap.model.Currency;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

@Service
@RequiredArgsConstructor
public class ConversionService {

    public static final int roundingMultiplier = 100;
    private final CurrencyRatesRepository currencyRatesRepository;

    public ConversionResponsePayload convert (ConversionRequestPayload payload) {

        assert payload != null;
        assert payload.getGivenCurrency() != null;
        assert payload.getRequiredCurrency() != null;
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

    private static double roundResult(double converted) {
        return (Math.round(converted * roundingMultiplier)) / (double) roundingMultiplier;
    }

    private static ConversionResponsePayload createResponsePayload(ConversionRequestPayload payload, double converted) {
        ConversionResponsePayload conversionResponsePayload = new ConversionResponsePayload();
        conversionResponsePayload.setConvertedCurrency(payload.getRequiredCurrency());
        conversionResponsePayload.setAmount(roundResult(converted));
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
