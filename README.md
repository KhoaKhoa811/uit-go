# Hướng dẫn cài đặt và chạy hệ thống ở local

## Yêu cầu hệ thống

Trước khi bắt đầu, hãy đảm bảo máy tính của bạn đã cài đặt các công cụ sau:

1. **Git** - Công cụ quản lý mã nguồn
   - Tải về và cài đặt tại: https://git-scm.com/

2. **Docker Desktop** - Môi trường container hóa
   - Tải về và cài đặt tại: https://www.docker.com/

## Các bước cài đặt

### 1. Clone repository
```bash
git clone <repository-url>
```

### 2. Truy cập vào thư mục backend
```bash
cd uit-go/backend
```

### 3. Chạy project với Docker
```bash
docker compose up --build -d
```

Lệnh này sẽ:
- Build các Docker images cần thiết
- Khởi động các containers ở chế độ background (`-d`)
- Tự động cấu hình môi trường development

## Kiểm tra hệ thống

Sau khi chạy lệnh trên, hệ thống sẽ được khởi động. Bạn có thể kiểm tra trạng thái của các containers bằng lệnh:
```bash
docker compose ps
```

## Dừng hệ thống

Để dừng hệ thống, sử dụng lệnh:
```bash
docker compose down
```

## Ghi chú

- Đảm bảo Docker Desktop đang chạy trước khi thực hiện các lệnh Docker
- Nếu gặp lỗi về quyền truy cập, hãy chạy terminal với quyền administrator (Windows) hoặc sử dụng `sudo` (Linux/Mac)
