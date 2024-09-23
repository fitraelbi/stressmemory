# Stage 1: Build the Go binary
FROM golang:alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . .

# Unduh dependency dan build aplikasi
RUN go mod download

# Copy the source code into the container

# Build the Go application
RUN go build -o /memory-consumer

# Stage 2: Run the Go binary
FROM alpine:latest

# Set timezone (optional)
RUN apk add --no-cache tzdata

# Copy the built binary from the builder
COPY --from=builder /memory-consumer /memory-consumer

# Expose port 8080 for the web server
EXPOSE 8080

# Command to run the binary
CMD ["/memory-consumer"]
