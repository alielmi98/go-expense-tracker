# Use the official Go image as the base image
FROM golang:1.22.2-alpine as builder

# Set the working directory to /app
WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -v -o expense-api ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/expense-api /app/expense-api
COPY --from=builder /app/config/config-docker.yml /app/config/config-docker.yml 
COPY --from=builder /app/docs /app/docs

ENV APP_ENV=docker

CMD ["./expense-api"]