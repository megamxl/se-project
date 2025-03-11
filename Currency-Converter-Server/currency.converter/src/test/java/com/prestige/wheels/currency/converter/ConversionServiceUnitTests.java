package com.prestige.wheels.currency.converter;

import com.prestige.wheels.currency.converter.repository.CurrencyRatesRepository;
import com.prestige.wheels.currency.converter.service.ConversionService;
import com.prestige.wheels.currency.converter.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.converter.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.converter.soap.model.Currency;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

import java.util.Optional;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertThrows;
import static org.mockito.Mockito.when;

public class ConversionServiceUnitTests {

    @Mock
    CurrencyRatesRepository currencyRatesRepository;

    ConversionService conversionService = new ConversionService(currencyRatesRepository);

    @BeforeEach
    void setUp() {
        MockitoAnnotations.openMocks(this);
        conversionService = new ConversionService(currencyRatesRepository);
    }

    @Test
    void testConvertToEuroWithEuro() {

        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(23);
        payload.setGivenCurrency(Currency.EUR);
        payload.setRequiredCurrency(Currency.EUR);

        ConversionResponsePayload responsePayload = conversionService.convert(payload);

        assertThat(responsePayload.getConvertedCurrency()).isEqualTo(Currency.EUR);
        assertThat(responsePayload.getAmount()).isEqualTo(23);
    }

    @Test
    void testConvertWithNullPayload() {

        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setRequiredCurrency(null);
        payload.setGivenCurrency(null);

        assertThrows(AssertionError.class, () -> conversionService.convert(null));

        assertThrows(AssertionError.class, () -> conversionService.convert(payload));

        payload.setGivenCurrency(Currency.EUR);

        assertThrows(AssertionError.class, () -> conversionService.convert(null));

    }

    @Test
    void testConvertWithDifferentCurrencies() {
        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(100);
        payload.setGivenCurrency(Currency.EUR);
        payload.setRequiredCurrency(Currency.USD);

        // Mock repository to return a conversion rate for USD
        when(currencyRatesRepository.getRateByCurrency(Currency.USD)).thenReturn(Optional.of(1.2));

        ConversionResponsePayload responsePayload = conversionService.convert(payload);

        assertThat(responsePayload.getConvertedCurrency()).isEqualTo(Currency.USD);
        assertThat(responsePayload.getAmount()).isEqualTo(120);
    }

    @Test
    void testConvertFromNonEuroToEuro() {
        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(50);
        payload.setGivenCurrency(Currency.USD);
        payload.setRequiredCurrency(Currency.EUR);

        // Mock repository to return a conversion rate for USD
        when(currencyRatesRepository.getRateByCurrency(Currency.USD)).thenReturn(Optional.of(1.2));

        ConversionResponsePayload responsePayload = conversionService.convert(payload);

        assertThat(responsePayload.getConvertedCurrency()).isEqualTo(Currency.EUR);
        assertThat(responsePayload.getAmount()).isEqualTo(41.66666666666667); // 50 / 1.2
    }


    @Test
    void testConvertFromEuroToNonEuro() {
        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(200);
        payload.setGivenCurrency(Currency.EUR);
        payload.setRequiredCurrency(Currency.USD);

        // Mock repository to return a conversion rate for USD
        when(currencyRatesRepository.getRateByCurrency(Currency.USD)).thenReturn(Optional.of(1.2));

        ConversionResponsePayload responsePayload = conversionService.convert(payload);

        assertThat(responsePayload.getConvertedCurrency()).isEqualTo(Currency.USD);
        assertThat(responsePayload.getAmount()).isEqualTo(240);
    }

    @Test
    void testConvertWithCurrencyNotFound() {
        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(100);
        payload.setGivenCurrency(Currency.USD);
        payload.setRequiredCurrency(Currency.GBP);

        // Mock repository to throw an exception when GBP is requested
        when(currencyRatesRepository.getRateByCurrency(Currency.GBP)).thenReturn(Optional.empty());

        assertThrows(RuntimeException.class, () -> conversionService.convert(payload));
    }

    @Test
    void testConvertWithIdenticalCurrencies() {
        ConversionRequestPayload payload = new ConversionRequestPayload();
        payload.setAmount(150);
        payload.setGivenCurrency(Currency.USD);
        payload.setRequiredCurrency(Currency.USD);

        ConversionResponsePayload responsePayload = conversionService.convert(payload);

        assertThat(responsePayload.getConvertedCurrency()).isEqualTo(Currency.USD);
        assertThat(responsePayload.getAmount()).isEqualTo(150);
    }
}
