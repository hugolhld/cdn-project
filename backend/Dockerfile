# Use official Golang image as base
FROM golang:1.23.6 AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./
RUN go mod tidy

# Copy the entire backend source code
COPY . .

COPY .env /app/.env

# Build the Go application (force Linux binary)
RUN GOOS=linux GOARCH=amd64 go build -o main .

# Use a minimal base image for running the app
FROM alpine:latest

# Set working directory
WORKDIR /app

# Copy the compiled Go binary from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Ensure the binary is executable
RUN chmod +x ./main

# # Expose the backend port
# EXPOSE 8082

# Set the default command
CMD ["./main"]
