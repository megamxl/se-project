```
docker build --build-arg SERVICE_LOCATION=cmd/userService/main.go -t user-service .
```

```
docker build --build-arg SERVICE_LOCATION=cmd/carsService/main.go -t car-service . 
```

```
docker build --build-arg SERVICE_LOCATION=cmd/bookingService/main.go -t booking-service .
```

protoc --go_out=Rental-Server/internal/communication/converter/grpc --go_opt=paths=source_relative --go-grpc_out=Rental-Server/internal/communication/converter/grpc --go-grpc_opt=paths=source_relative proto/currencyConvertor.proto