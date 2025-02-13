# Utiliser une version stable de Go
FROM golang:1.23.6 AS builder

# Définir le répertoire de travail
WORKDIR /app

# Copier tous les fichiers nécessaires
COPY go.mod ./
RUN go mod tidy

COPY . .

# Compiler le programme
RUN go build -o load_balancer main.go

# Exécuter le Load Balancer
CMD ["./load_balancer"]
