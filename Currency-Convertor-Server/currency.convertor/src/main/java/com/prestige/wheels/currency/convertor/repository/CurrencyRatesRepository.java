package com.prestige.wheels.currency.convertor.repository;

import com.prestige.wheels.currency.convertor.repository.ecb.model.Envelope;
import com.prestige.wheels.currency.convertor.soap.model.Currency;
import jakarta.xml.bind.JAXBContext;
import jakarta.xml.bind.JAXBException;
import jakarta.xml.bind.Unmarshaller;
import lombok.extern.slf4j.Slf4j;
import org.springframework.core.io.Resource;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.math.BigDecimal;
import java.net.URI;
import java.text.ParseException;
import java.text.SimpleDateFormat;
import java.util.Date;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;

@Service
@Slf4j
public class CurrencyRatesRepository {

    private final RestClient restClient;

    private final Unmarshaller unmarshaller;

    private Map<Currency, Double> rates;

    SimpleDateFormat sdf = new SimpleDateFormat("yyyy-MM-dd");

    private Date dateFromLastExchangeUpdate;

    public CurrencyRatesRepository() throws JAXBException {
        this.unmarshaller = getUnmarshaller();
        this.restClient = RestClient.create();
        this.rates = new HashMap<>();
    }

    public Optional<Double> getRateByCurrency(Currency currency) {
        updateRatesIfEmptyOrOlderAsToday();

        return Optional.of(rates.get(currency));
    }

    private void updateRatesIfEmptyOrOlderAsToday() {

        if(rates == null || rates.isEmpty()){
            try {
                getRatesFromExchangeAndUpdateRatesMap();
            } catch (Exception e) {
                throw new RuntimeException(e);
            }
        }

        if(dateFromLastExchangeUpdate !=null && dateFromLastExchangeUpdate.before(new Date())) {
            log.info("Currencies last exchange update date: {}", dateFromLastExchangeUpdate);
            //TODO Do this and gather info when it is new ?? weekend ??
        }

    }

    private void getRatesFromExchangeAndUpdateRatesMap() throws IOException, JAXBException {
        ByteArrayInputStream ratesXmlAsByteStream = getEuroFxRefBytes().orElseThrow(() -> new IllegalStateException("Euro Fx Ref bytes not retrievable. Contact the currency Converter support !"));

        Envelope envelope = (Envelope) unmarshaller.unmarshal(ratesXmlAsByteStream);

        envelopToRatesMap(envelope);

        log.info("Rates retrieved from Euro Fx Ref: {}", envelope);
    }

    private Optional<ByteArrayInputStream> getEuroFxRefBytes() throws IOException {
        ResponseEntity<Resource> response = restClient.get()
                .uri(URI.create("https://www.ecb.europa.eu/stats/eurofxref/eurofxref-daily.xml"))
                .retrieve()
                .toEntity(Resource.class);

        // Handle redirect manually
        if (response.getStatusCode().is3xxRedirection() && response.getHeaders().getLocation() != null) {
            URI newLocation = response.getHeaders().getLocation();
            System.out.println("Redirected to: " + newLocation);

            response = restClient.get()
                    .uri(newLocation)
                    .retrieve()
                    .toEntity(Resource.class);
        }

        if (response.getStatusCode().is2xxSuccessful()) {
            return Optional.of(new ByteArrayInputStream(response.getBody().getContentAsByteArray()));
        }

        return Optional.empty();
    }

    private static Unmarshaller getUnmarshaller() throws JAXBException {
        return JAXBContext.newInstance(Envelope.class).createUnmarshaller();
    }

    private void envelopToRatesMap(Envelope envelope) {
        envelope.getCube().getTimeCube().getCurrencyRates().forEach(cur -> rates.put(cur.getCurrency(), cur.getRate()));
        try {
            this.dateFromLastExchangeUpdate =sdf.parse(envelope.getCube().getTimeCube().getTime());
        } catch (ParseException e) {
            throw new RuntimeException(e);
        }
    }

}
