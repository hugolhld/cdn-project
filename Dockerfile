# Utilisation de l'image de base Go
FROM golang:1.21 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o app .

# Image finale
FROM debian:stable-slim
WORKDIR /root/

COPY --from=builder /app/app .

CMD ["./app"]
