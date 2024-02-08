FROM golang:1.22.0-alpine3.19 AS builder

COPY . /app/.

WORKDIR /app/src

RUN go build -o /app/apiserver

FROM alpine:3.19.1 AS build-release-stage

WORKDIR /

COPY --from=builder /app/apiserver /app/apiserver

EXPOSE 8080

CMD [ "/app/apiserver" ]
