# Create http service

## Yêu Cầu
### Cơ bản
Request tới một api và in kết quả ra màn hình

### Cho phép sử dụng các options khi request (Method, headers, body...)
Mỗi một request đều có thể có một options khác nhau. Người dùng có thể sử dụng các options đó

### Cho phép upload một file mới
Cho phép người dùng upload file mới sử dụng FormData

### Hỗ trợ parse response theo các kiểu khác nhau (json, text...)
Người dùng có thể tuỳ chọn mỗi kiểu response khác nhau tuỳ theo mục đích.

### Hỗ trợ interceptor để hook vào request, response
Cho phép người dùng có thể thêm các hook vào request và response để decoupling code giữa service và logic của app.

### Clean code sử dụng ES6 modules
Viết lại code sử dụng ES6 modules thành một module hoàn chỉnh.


## Tài liệu
- [Fetch API](https://developer.mozilla.org/en-US/docs/Web/API/Fetch_API/Using_Fetch)
- [Mock API](https://jsonplaceholder.typicode.com/)
- [Interceptor example](https://github.com/axios/axios#interceptors)

