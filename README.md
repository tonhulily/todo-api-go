# Ứng dụng Todo-list Backend (Go)

Đây là mã nguồn cho phần backend của ứng dụng Todo-list, được xây dựng bằng Go và Gin Framework. Backend này cung cấp các API để quản lý danh sách công việc.

---

## Các tính năng chính

* `POST /tasks`: Tạo một công việc mới.
* `GET /tasks`: Lấy toàn bộ danh sách công việc.
* `PUT /tasks/:id`: Cập nhật trạng thái của một công việc (ví dụ: `PUT /tasks/123` để cập nhật công việc có ID là 123).
* `DELETE /tasks/:id`: Xóa một công việc khỏi danh sách (ví dụ: `DELETE /tasks/123`).
> **Lưu ý:** Khi triển khai trên **Render**, bạn sẽ cấu hình các biến môi trường này trực tiếp trên nền tảng cloud.
---

## Cấu hình biến môi trường

Để ứng dụng hoạt động, bạn cần có một file `.env` chứa các biến môi trường sau:

* **Tạo file `.env`**: Ở thư mục gốc của dự án, hãy tạo một file tên là `.env` và thêm nội dung này vào:
    ```ini
    PORT=8080
    FRONTEND_URL=http://localhost:5173
    ```

> **Lưu ý:** Nếu bạn triển khai ứng dụng trên **Render**, bạn chỉ cần cài đặt biến môi trường **`FRONTEND_URL`** với giá trị là URL của frontend đã deploy.

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