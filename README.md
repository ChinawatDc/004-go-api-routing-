# 004-go-api-routing

โปรเจกต์นี้สาธิตการสร้าง RESTful API เบื้องต้นด้วยภาษา Go โดยใช้ Gin Framework พร้อมจัดการ Routing และ Middleware พื้นฐาน เช่น Logging และการแยกไฟล์ Route

---

## 📦 Features

- ใช้ Gin ในการจัดการ API Routing
- Middleware แบบกำหนดเอง (LoggerMiddleware)
- กลุ่ม route (`/api/users`) พร้อมพารามิเตอร์
- โครงสร้างโปรเจกต์แบบแยกไฟล์เพื่อความชัดเจน

---

## 📁 โครงสร้างโปรเจกต์

```
go-api-routing/
├── go.mod
├── main.go
├── middleware/
│ └── logger.go
└── routes/
└── user.go
```

---

## 🚀 วิธีติดตั้งและใช้งาน

### 1. Clone โปรเจกต์

```bash
mkdir go-api-routing && cd go-api-routing
```

### 2. สร้างไฟล์ตามที่กำหนด และสร้าง go.mod

```bash
go mod init go-api-routing
go get github.com/gin-gonic/gin
```

### 3. สร้างโครงสร้างโปรเจกต์พื้นฐาน

```bash
mkdir routes middleware
touch main.go
touch routes/user.go
touch middleware/logger.go
touch README.md
```

หรือหากใช้ Windows:

```bash
mkdir routes, middleware
ni main.go
ni routes\user.go
ni middleware\logger.go
ni README.md
```

### 4. รันโปรเจกต์

```bash
go run main.go
```

## 📌 API Endpoints

| Method | URL              | คำอธิบาย               |
| ------ | ---------------- | ---------------------- |
| GET    | `/`              | ข้อความต้อนรับ         |
| GET    | `/api/users`     | รายชื่อผู้ใช้          |
| GET    | `/api/users/:id` | รายละเอียดผู้ใช้ตาม id |

## 🛠 Middleware ที่มี

LoggerMiddleware
พิมพ์ log การเรียก API โดยรวมข้อมูล เช่น method, path, status code, และเวลาในการประมวลผล

## 📚 ความรู้เพิ่มเติม

- [Gin Middleware](https://gin-gonic.com/docs/examples/custom-middleware/)
- [Gin Routing](https://gin-gonic.com/docs/examples/route-groups/)
- [Gin Parameter Handling](https://gin-gonic.com/docs/examples/parameters/)

## 📄 License

MIT © 2025 chinawatDC
