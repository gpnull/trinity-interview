- **Requirement**: https://docs.google.com/document/d/1NDrFxIS_4WnyWq0L6NP-Uo7gFMrLRCwu2UMEPtgPIYo/edit?tab=t.0

# Technical Decisions

## Frameworks and Libraries
- **Gin Gonic**: Được sử dụng làm framework chính cho việc xây dựng API RESTful. Gin cung cấp hiệu suất cao và dễ dàng sử dụng với các tính năng như middleware, routing, và JSON validation.
- **GORM**: Được chọn làm ORM (Object-Relational Mapping) để tương tác với cơ sở dữ liệu MySQL. GORM hỗ trợ các tính năng như auto-migration, associations, và transactions.
- **Viper**: Được sử dụng để quản lý cấu hình ứng dụng. Viper cho phép tải cấu hình từ nhiều nguồn khác nhau và hỗ trợ hot-reload.
- **Bcrypt**: Được sử dụng để mã hóa mật khẩu, đảm bảo an toàn cho thông tin người dùng.

## Authentication
- **JWT (JSON Web Tokens)**: Được sử dụng để xác thực người dùng. JWT cho phép tạo và xác thực token một cách an toàn, giúp bảo vệ các endpoint của API.

## Middleware
- **CORS (Cross-Origin Resource Sharing)**: Được cấu hình để cho phép các yêu cầu từ các nguồn khác nhau, hỗ trợ phát triển ứng dụng web client-server.
- **Auth Middleware**: Được triển khai để kiểm tra tính hợp lệ của JWT token trong các yêu cầu đến, đảm bảo chỉ người dùng đã xác thực mới có thể truy cập các tài nguyên bảo mật.

## Database
- **MySQL**: Được chọn làm hệ quản trị cơ sở dữ liệu. MySQL là một lựa chọn phổ biến với hiệu suất cao và dễ dàng tích hợp với GORM.

## Error Handling
- **Custom Error Types**: Được định nghĩa để quản lý các lỗi phổ biến trong ứng dụng, giúp dễ dàng xử lý và debug.

## Validation
- **Go Playground Validator**: Được sử dụng để xác thực dữ liệu đầu vào, đảm bảo tính toàn vẹn và chính xác của dữ liệu trước khi xử lý.

## Project Structure
- **Modular Design**: Dự án được tổ chức theo cấu trúc module, giúp dễ dàng mở rộng và bảo trì. Các thành phần chính bao gồm `handler`, `service`, `repository`, `model`, và `utils`.

## Configuration
- **Environment Configuration**: Sử dụng Viper để tải cấu hình từ tệp YAML, cho phép dễ dàng thay đổi cấu hình mà không cần thay đổi mã nguồn.

## Logging
- **Standard Logging**: Sử dụng gói `log` của Go để ghi lại các thông tin quan trọng và lỗi, giúp theo dõi hoạt động của ứng dụng.

# Assumptions made

## Giả định quy trình:
- **Khởi tạo Voucher**: Admin tạo voucher.
- **Khởi tạo Campaign**: Admin tạo campaign có liên kết với voucher đã được tạo.
- **Người dùng đăng ký tài khoản**: Người dùng thực hiện đăng ký với liên kết campaign được tạo.
- **Người dùng nhận voucher sau đăng ký**: Sau khi đăng ký người dùng sẽ nhận được voucher discount.

