package com.prestige.wheels.currency.convertor.repository.ecb.model;


import jakarta.xml.bind.annotation.*;
import lombok.Getter;
import lombok.Setter;

import java.util.List;

@Setter
@Getter
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
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class Sender {

        @XmlElement(name = "name", namespace = "http://www.gesmes.org/xml/2002-08-01")
        private String name;

    }

    @Setter
    @Getter
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class Cube {

        @XmlElement(name = "Cube")
        private TimeCube timeCube;

    }

    @Getter
    @Setter
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class TimeCube{
        @XmlAttribute(name = "time")
        private String time;

        @XmlElement(name = "Cube")
        private List<CurrencyRate> currencyRates;
    }

    @Setter
    @Getter
    @XmlAccessorType(XmlAccessType.FIELD)
    public static class CurrencyRate {

        @XmlAttribute(name = "currency")
        private String currency;

        @XmlAttribute(name = "rate")
        private double rate;

    }
}