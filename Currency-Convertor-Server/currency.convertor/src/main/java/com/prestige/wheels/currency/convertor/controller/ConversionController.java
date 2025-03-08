package com.prestige.wheels.currency.convertor.controller;

import com.prestige.wheels.currency.convertor.repository.CurrencyRatesRepository;
import com.prestige.wheels.currency.convertor.soap.model.*;
import jakarta.xml.bind.JAXBException;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;

import java.io.IOException;

@Endpoint
public class ConversionController {

    public static final String NAMESPACE_URI = "http://prestige-wheels.at/conversion/";

    @Autowired
    CurrencyRatesRepository currencyRatesRepository;

    @PayloadRoot(localPart = "conversionRequest", namespace = NAMESPACE_URI)
    @ResponsePayload
    ConversionResponse convertCurrency(@RequestPayload ConversionRequest request) {
        ConversionResponse conversionResponse = new ConversionResponse();

        try {
            currencyRatesRepository.getRatesFromExchange();
        } catch (IOException e) {
            throw new RuntimeException(e);
        } catch (JAXBException e) {
            throw new RuntimeException(e);
        }

        ConversionResponsePayload conversionResponsePayload = new ConversionResponsePayload();
        conversionResponsePayload.setAmount(23);
        conversionResponsePayload.setConvertedCurrency(Currency.CHF);

        conversionResponse.setResponse(conversionResponsePayload);

        return conversionResponse;
    }

}
