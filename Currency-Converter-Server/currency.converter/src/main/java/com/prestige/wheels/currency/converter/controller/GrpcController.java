package com.prestige.wheels.currency.converter.controller;

import com.prestige.wheels.currency.converter.grpc.ConvertorGrpc;
import com.prestige.wheels.currency.converter.grpc.CurrencyConvertor;
import com.prestige.wheels.currency.converter.service.ConversionService;
import com.prestige.wheels.currency.converter.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.converter.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.converter.soap.model.Currency;
import io.grpc.stub.StreamObserver;
import lombok.AllArgsConstructor;
import lombok.RequiredArgsConstructor;
import org.springframework.stereotype.Service;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

@Service
@RequiredArgsConstructor
public class GrpcController extends ConvertorGrpc.ConvertorImplBase {

    private final ConversionService conversionService;

    @Override
    public void convert(CurrencyConvertor.conversionRequest request, StreamObserver<CurrencyConvertor.conversionResponse> responseObserver) {

        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(request.getAmount());
        payload.setGivenCurrency( Currency.fromValue(  request.getGiven().name()));
        payload.setRequiredCurrency( Currency.fromValue( request.getTarget().name()));

        ConversionResponsePayload convert = conversionService.convert(payload);

        CurrencyConvertor.conversionResponse.Builder builder = CurrencyConvertor.conversionResponse.newBuilder().
                setAmount(convert.getAmount())
                .setConverted(CurrencyConvertor.currency.valueOf(convert.getConvertedCurrency().name()));

        responseObserver.onNext(builder.build());
        responseObserver.onCompleted();

    }

    @Override
    public void getAvailableCurrency(CurrencyConvertor.empty request, StreamObserver<CurrencyConvertor.currencyList> responseObserver) {

        CurrencyConvertor.currency[] values = CurrencyConvertor.currency.values();

        List<CurrencyConvertor.currency> list = new ArrayList<>();

        for (CurrencyConvertor.currency value : values) {
            if(value.equals(CurrencyConvertor.currency.UNRECOGNIZED)){
                continue;
            }
            list.add(value);
        }

        CurrencyConvertor.currencyList.Builder builder = CurrencyConvertor.currencyList.newBuilder()
                .addAllCurrencies(list);

        responseObserver.onNext(builder.build());
        responseObserver.onCompleted();
    }
}
