package com.prestige.wheels.currency.converter;

import com.prestige.wheels.currency.converter.service.ConversionService;
import com.prestige.wheels.currency.converter.soap.model.ConversionRequestPayload;
import com.prestige.wheels.currency.converter.soap.model.ConversionResponsePayload;
import com.prestige.wheels.currency.converter.soap.model.Currency;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
class ApplicationTests {

	@Autowired
	ConversionService conversionService;

	@Test
	void integrationTest() {

		ConversionRequestPayload payload = new ConversionRequestPayload();
		payload.setAmount(23);
		payload.setGivenCurrency(Currency.EUR);
		payload.setRequiredCurrency(Currency.USD);

		ConversionResponsePayload responsePayload = conversionService.convert(payload);

		assertThat(responsePayload.getAmount())
				.isNotNull()
				.isFinite()
				.isNotNegative();

		assertThat(responsePayload.getConvertedCurrency())
				.isNotNull()
				.isEqualTo(Currency.USD);


	}

}
