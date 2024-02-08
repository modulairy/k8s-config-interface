FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

WORKDIR /app/src

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/apiserver

# FROM gcr.io/distroless/base-debian12 AS build-release-stage

# WORKDIR /

# COPY --from=builder /app/apiserver /apiserver

EXPOSE 8080

CMD [ "/app/apiserver" ]
