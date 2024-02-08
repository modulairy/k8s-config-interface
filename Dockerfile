FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN go build -o moperator

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/src/moperator /app/moperator
EXPOSE 8080

CMD ["./moperator"]
