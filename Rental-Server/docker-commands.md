```
docker build --build-arg SERVICE_LOCATION=cmd/userService/main.go -t user-service .
```

```
docker build --build-arg SERVICE_LOCATION=cmd/carsService/main.go -t car-service . 
```

```
docker build --build-arg SERVICE_LOCATION=cmd/bookingService/main.go -t booking-service .
```