# Build stage
FROM golang:1.21.5-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main app.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.14.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY .env .
COPY start.sh .
COPY wait-for.sh .
COPY repository/migration ./migration

EXPOSE 8080
CMD [ "/app/start.sh" ]