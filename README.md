# Ứng dụng To-do Backend (Go)

Đây là mã nguồn cho phần backend của ứng dụng Todo-list, được xây dựng bằng Go và Gin Framework. Backend này cung cấp các API để quản lý danh sách công việc.

---

## Các tính năng chính

* `POST /tasks`: Tạo một công việc mới.
* `GET /tasks`: Lấy toàn bộ danh sách công việc.
* `PUT /tasks/:id`: Cập nhật trạng thái của một công việc.
* `DELETE /tasks/:id`: Xóa một công việc.

---

## Cấu hình Biến Môi Trường

Để chạy ứng dụng, bạn cần sử dụng các biến môi trường.

1.  **Cài đặt thư viện `godotenv`:**
    ```bash
    go get [github.com/joho/godotenv](https://github.com/joho/godotenv)
    ```

2.  **Tạo file `.env`:**
    Trong thư mục gốc của dự án, hãy tạo một file tên là `.env` với nội dung sau:
    ```ini
    PORT=8080
    FRONTEND_URL=http://localhost:5173
    ```

> **Lưu ý:** Khi triển khai trên **Render**, bạn sẽ cấu hình các biến môi trường này trực tiếp trên nền tảng cloud.

---

## Hướng dẫn Chạy trên Local

1.  Đảm bảo bạn đã cài đặt Go.
2.  Clone repository này về máy.
3.  Mở terminal, di chuyển vào thư mục dự án và chạy lệnh:
    ```bash
    go run .
    ```

---

## Thông tin Triển khai

* **Link đã deploy:** `https://todo-api-go-4ki3.onrender.com`

Khi triển khai trên Render, bạn cần cấu hình biến môi trường `FRONTEND_URL` với giá trị là URL đã deploy của frontend:
* `FRONTEND_URL=https://todo-list-frontend-1.netlify.app`

---

## Kết nối với Frontend

* Dự án frontend tương ứng được đặt tại repository: `https://github.com/tonhulily/todo-list_frontend`