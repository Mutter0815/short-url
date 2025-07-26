FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o /short-url ./cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /short-url .

EXPOSE 8080
CMD ["./short-url"]
