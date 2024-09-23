# Stage 1: Build the Go binary
FROM golang:alpine as builder

WORKDIR /app

# Copy go.mod and go.sum, then download dependencies
COPY . .
RUN go mod download

# Copy the source code


# Build the Go binary
RUN GOOS=linux GOARCH=amd64 go build -o /go-mem-app main.go

# Stage 2: Create the final image
FROM alpine:latest

WORKDIR /app

# Copy the Go binary from the builder stage
COPY --from=builder /go-mem-app .

# Expose port (optional, if your service exposes a port)
EXPOSE 8080

# Run the Go binary
CMD ["./go-mem-app"]
