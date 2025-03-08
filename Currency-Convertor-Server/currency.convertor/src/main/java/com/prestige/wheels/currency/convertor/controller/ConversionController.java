package com.prestige.wheels.currency.convertor.controller;

import com.prestige.wheels.currency.convertor.soap.model.*;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;

@Endpoint
public class ConversionController {

    public static final String NAMESPACE_URI = "http://prestige-wheels.at/conversion/";

    @PayloadRoot(localPart = "conversionRequest", namespace = NAMESPACE_URI)
    @ResponsePayload
    ConversionResponse convertCurrency(@RequestPayload ConversionRequest request) {
        ConversionResponse conversionResponse = new ConversionResponse();

        ConversionResponsePayload conversionResponsePayload = new ConversionResponsePayload();
        conversionResponsePayload.setAmount(23);
        conversionResponsePayload.setConvertedCurrency(Currency.CHF);

        conversionResponse.setResponse(conversionResponsePayload);

        return conversionResponse;
    }

}
