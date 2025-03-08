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

        //TODO assert payload fields
        if (payload.getGivenCurrency().equals(payload.getRequiredCurrency())) {
            return createResponsePayload(payload, BigDecimal.valueOf(payload.getAmount()));
        }

        // first of all check if is not EUR convert it to EUR
        BigDecimal amountInEuro = convertToEuro(payload.getGivenCurrency(),  BigDecimal.valueOf(payload.getAmount()));

        // afterward convert from euro to what is needed
        BigDecimal converted = convertFromEuroToTargetCurrency(payload.getRequiredCurrency(), amountInEuro);

        return createResponsePayload(payload, converted);
    }

    private static ConversionResponsePayload createResponsePayload(ConversionRequestPayload payload, BigDecimal converted) {
        ConversionResponsePayload conversionResponsePayload = new ConversionResponsePayload();
        conversionResponsePayload.setConvertedCurrency(payload.getRequiredCurrency());
        conversionResponsePayload.setAmount(converted.doubleValue());
        return conversionResponsePayload;
    }

    private BigDecimal convertToEuro(Currency currency, BigDecimal amount) {

        if (currency == Currency.EUR){
            return amount;
        }

        BigDecimal exchangeValue = getRateForCurrency(currency);

        return amount.divide(exchangeValue, RoundingMode.HALF_EVEN);
    }


    private BigDecimal convertFromEuroToTargetCurrency(Currency currency, BigDecimal amount) {
        BigDecimal exchangeValue = getRateForCurrency(currency);

        return amount.multiply(exchangeValue);
    }


    private BigDecimal getRateForCurrency(Currency currency) {
        return currencyRatesRepository.getRateByCurrency(currency).orElseThrow(() -> new RuntimeException("Currency not found in the local repository"));
    }
}
