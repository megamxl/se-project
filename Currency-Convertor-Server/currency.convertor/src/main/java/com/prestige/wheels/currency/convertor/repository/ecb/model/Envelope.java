package com.prestige.wheels.currency.convertor.repository.ecb.model;


import com.prestige.wheels.currency.convertor.soap.model.Currency;
import jakarta.xml.bind.annotation.*;
import lombok.Getter;
import lombok.Setter;
import lombok.ToString;

import java.math.BigDecimal;
import java.util.List;

@Setter
@Getter
@ToString
@XmlRootElement(name = "Envelope", namespace = "http://www.gesmes.org/xml/2002-08-01")
@XmlAccessorType(XmlAccessType.FIELD)
public class Envelope {

    @XmlElement(name = "subject", namespace = "http://www.gesmes.org/xml/2002-08-01")
    private String subject;

    @XmlElement(name = "Sender", namespace = "http://www.gesmes.org/xml/2002-08-01")
    private Sender sender;

    @XmlElement(name = "Cube")
    private Cube cube;

    @Setter
    @Getter
    @ToString
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class Sender {

        @XmlElement(name = "name", namespace = "http://www.gesmes.org/xml/2002-08-01")
        private String name;

    }

    @Setter
    @Getter
    @ToString
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class Cube {

        @XmlElement(name = "Cube")
        private TimeCube timeCube;

    }

    @Getter
    @Setter
    @ToString
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class TimeCube{
        @XmlAttribute(name = "time")
        private String time;

        @XmlElement(name = "Cube")
        private List<CurrencyRate> currencyRates;
    }

    @Setter
    @Getter
    @ToString
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class CurrencyRate {

        @XmlAttribute(name = "currency")
        private Currency currency;

        @XmlAttribute(name = "rate")
        private Double rate;

    }
}