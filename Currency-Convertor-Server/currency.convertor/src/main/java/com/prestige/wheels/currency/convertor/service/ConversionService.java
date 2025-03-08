package com.prestige.wheels.currency.convertor.service;

import com.prestige.wheels.currency.convertor.soap.model.ConversionRequest;
import com.prestige.wheels.currency.convertor.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.convertor.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.convertor.soap.model.Currency;
import org.springframework.stereotype.Service;

import java.math.BigDecimal;

@Service
public class ConversionService {

    ConversionResponsePayload convert (ConversionRequestPayload payload) {

        // first of all check if is not EUR convert it to EUR
        payload.getGivenCurrency();

        // afterward convert from euro to what is needed

        return null;
    }

    private BigDecimal convertToEuro(Currency currency, double amount) {

        if (currency == Currency.EUR){
            return new BigDecimal(amount);
        }

        return null;

    }




}
