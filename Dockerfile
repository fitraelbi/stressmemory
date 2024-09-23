# Gunakan base image yang ringan seperti Alpine dengan Golang
FROM golang:alpine AS builder

# Buat direktori kerja untuk aplikasi
WORKDIR /app

# Salin semua file ke container
COPY . .

# Unduh dependency dan build aplikasi
RUN go mod download
RUN go build -o memory-stress-service

# Buat image yang lebih ringan hanya dengan binary aplikasi
FROM alpine:latest

WORKDIR /root/

# Salin binary dari stage builder
COPY --from=builder /app/memory-stress-service .

# Set port untuk aplikasi
EXPOSE 8080

# Jalankan aplikasi
CMD ["./memory-stress-service"]
