# Gunakan base image golang versi terbaru
FROM golang:1.23

# Set lokasi kerja di dalam container
WORKDIR /app

# Salin file module dan Go dependencies
COPY go.mod go.sum ./
ENV GOPROXY=https://proxy.golang.org
RUN go mod download

# Salin semua file dari proyek ke dalam direktori /app di container
COPY . .

# Build aplikasi ke dalam binary dengan nama 'main'
RUN go build -o main ./cmd

# Mengekspos port dimana aplikasi berada
EXPOSE 8080

# Perintah yang akan dijalankan ketika container dijalankan
CMD ["./main"]
