FROM alpine as certs
RUN apk update && apk add ca-certificates

FROM golang:1.18-alpine3.15 AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 go build -o faas -mod=vendor -ldflags='-s -w'  -installsuffix cgo cmd/main.go

FROM scratch
COPY --from=certs /etc/ssl/certs /etc/ssl/certs

WORKDIR /faas
COPY --from=builder ./build/faas ./cmd/

EXPOSE 80
ENTRYPOINT ["./cmd/faas","-config=/configs/config.yml"]
