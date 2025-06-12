# 🧾 Simple Payroll System - PayFlow

A complete payroll management backend system built with Go and PostgreSQL.

## Features

- Employee attendance tracking
- Overtime & reimbursement submission
- Admin-managed attendance periods
- Payroll processing with breakdowns
- Payslip generation
- Admin summaries
- Fully auditable system with logs

## Architecture

- **Clean Architecture** (domain, usecase, repository, delivery)
- **PostgreSQL** as the main DB
- **JSON-based HTTP API**
- **TDD** with unit + integration tests
- **Audit logging** and request traceability

## Technologies

- Go (1.23+)
- PostgreSQL
- Chi Router
- SQL migrations
- Docker (optional)

---

## 🚀 Running Locally

### 1. Run with Docker
```bash
docker-compose up -d
```

or

### 3. Run the app
```bash
go run main.go
```

---

## 📒 Access Swagger UI
```
http://localhost:8080/swagger/index.html
```

---

## 🧪 Unit Testing

### Generate mocks with Mockery
```bash
mockery --all
```

### Run all unit tests
```bash
go test ./...
```

---

## 🔌 Integration Testing

### Setup test users

Ensure test users exist in the DB:

- **Admin**
  ```json
  {
    "username": "admin",
    "password": "admin123"
  }
  ```

- **Employee**
  ```json
  {
    "username": "employee1",
    "password": "pass1"
  }
  ```

  for another employee
  ```json
  {
    "username": "employee{number employee 1-100}",
    "password": "pass{number employee 1-100}"
  }
  ```

### Run integration tests
```bash
go test ./test/integration/... -tags=integration 
```

> ✅ The integration tests perform real login, use JWT tokens, and call all main endpoints.

---

## 📎 Headers and Traceability

All APIs support traceability headers:

```http
X-Request-ID: <uuid>
Authorization: Bearer <JWT>
```

Audit logs will record:
- Request ID
- IP Address
- User ID
- Table & Action
- Timestamp

---

## 🛠️ Troubleshooting

| Error              | Solution |
|--------------------|----------|
| `401 Unauthorized` | Use correct credentials & include JWT in headers |
| `403 Forbidden`    | Make sure the user has the correct role (admin/employee) |
| `400 Bad Request`  | Check request payload format |
| Swagger not loading | Run `swag init` again |

---

