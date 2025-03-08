package com.prestige.wheels.currency.convertor.controller;

import com.prestige.wheels.currency.convertor.repository.CurrencyRatesRepository;
import com.prestige.wheels.currency.convertor.service.ConversionService;
import com.prestige.wheels.currency.convertor.soap.model.*;
import jakarta.xml.bind.JAXBException;
import lombok.RequiredArgsConstructor;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.ws.server.endpoint.annotation.Endpoint;
import org.springframework.ws.server.endpoint.annotation.PayloadRoot;
import org.springframework.ws.server.endpoint.annotation.RequestPayload;
import org.springframework.ws.server.endpoint.annotation.ResponsePayload;

import java.io.IOException;

@Endpoint
@RequiredArgsConstructor
public class ConversionController {

    public static final String NAMESPACE_URI = "http://prestige-wheels.at/conversion/";

    private final ConversionService conversionService;


    @PayloadRoot(localPart = "conversionRequest", namespace = NAMESPACE_URI)
    @ResponsePayload
    ConversionResponse convertCurrency(@RequestPayload ConversionRequest request) {
        ConversionResponse conversionResponse = new ConversionResponse();


        ConversionResponsePayload responsePayload = conversionService.convert(request.getReq());

        conversionResponse.setResponse(responsePayload);

        return conversionResponse;
    }

}
