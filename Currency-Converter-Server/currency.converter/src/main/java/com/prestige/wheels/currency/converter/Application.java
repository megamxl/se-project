package com.prestige.wheels.currency.converter;

import com.prestige.wheels.currency.converter.controller.GrpcController;
import io.grpc.Server;
import io.grpc.ServerBuilder;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.core.env.Environment;

import java.util.Arrays;

@SpringBootApplication
public class Application implements CommandLineRunner {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }

    @Autowired
    GrpcController  grpcController;

    @Autowired
    Environment env;

    @Override
    public void run(String... args) throws Exception {

        if (Arrays.asList(env.getActiveProfiles()).contains("test")) {
            return;
        }

        Server server = ServerBuilder
                .forPort(8082)
                .addService(grpcController).build();

        server.start();
        server.awaitTermination();
    }
}

