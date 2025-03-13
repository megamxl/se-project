package com.prestige.wheels.currency.converter.controller;

import com.prestige.wheels.currency.converter.service.ConversionService;
import com.prestige.wheels.currency.converter.soap.model.*;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;

import java.util.Arrays;

@Endpoint

public class ConversionController {

    public static final String NAMESPACE_URI = "http://prestige-wheels.at/conversion/";

    private final ConversionService conversionService;

    GetAvailableCurrencyResponse resp;

    public ConversionController(ConversionService conversionService) {
        this.conversionService = conversionService;
        this.resp = new GetAvailableCurrencyResponse();
        resp.getCurrencies().addAll(Arrays.stream(Currency.values()).toList());
    }


    @PayloadRoot(localPart = "conversionRequest", namespace = NAMESPACE_URI)
    @ResponsePayload
    ConversionResponse convertCurrency(@RequestPayload ConversionRequest request) {
        ConversionResponse conversionResponse = new ConversionResponse();


        ConversionResponsePayload responsePayload = conversionService.convert(request.getReq());

        conversionResponse.setResponse(responsePayload);

        return conversionResponse;
    }

    @PayloadRoot(localPart = "getAvailableCurrencyRequest", namespace = NAMESPACE_URI)
    @ResponsePayload
    GetAvailableCurrencyResponse getAvailableCurrency() {
        return resp ;
    }

}
