FROM golang:1.24.1-bookworm AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
RUN apt update && apt install -y tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY . .
RUN go build -o rental-server cmd/main.go

#FROM registry.redhat.io/ubi9/ubi
FROM alpine
RUN apk update --no-cache && apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /build/rental-server /app/rental-server
CMD ["/app/rental-server"]

