package com.prestige.wheels.currency.convertor.repository;

import com.prestige.wheels.currency.convertor.repository.ecb.model.*;
import jakarta.xml.bind.JAXBContext;
import jakarta.xml.bind.JAXBException;
import jakarta.xml.bind.Marshaller;
import jakarta.xml.bind.Unmarshaller;
import org.springframework.core.io.Resource;
import org.springframework.http.ResponseEntity;
import org.springframework.stereotype.Service;
import org.springframework.web.client.RestClient;

import java.io.ByteArrayInputStream;
import java.io.IOException;
import java.net.URI;
import java.nio.charset.StandardCharsets;

@Service
public class CurrencyRatesRepository {

    RestClient restClient = RestClient.create();

 // TODO Make private
 public void getRatesFromExchange() throws IOException, JAXBException {
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
     JAXBContext context = JAXBContext.newInstance(Envelope.class);
     Unmarshaller unmarshaller = context.createUnmarshaller();
     Marshaller marshaller = context.createMarshaller();

     Envelope envelope = (Envelope) unmarshaller.unmarshal(new ByteArrayInputStream(response.getBody().getContentAsByteArray()));

     System.out.println(new String(response.getBody().getContentAsByteArray(), StandardCharsets.UTF_8));
    }

}
