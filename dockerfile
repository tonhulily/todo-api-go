# Sử dụng image Golang chính thức làm base image
FROM golang:1.19-alpine

# Đặt thư mục làm việc bên trong container
WORKDIR /app

# Sao chép các file module và tải dependencies
COPY go.mod go.sum ./
RUN go mod download

# Sao chép source code của ứng dụng
COPY . .

# Xây dựng ứng dụng Go
RUN go build -o /go-todo-api

# Expose port mà ứng dụƯng sẽ chạy
EXPOSE 8080

# Lệnh để chạy ứng dụng
CMD ["/go-todo-api"]