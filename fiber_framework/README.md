# 🚀 Fiber Framework — RESTful API with Go

A production-style REST API built with the [Fiber v2](https://gofiber.io/) framework, PostgreSQL, Basic Auth middleware, and full unit + benchmark testing.

This module demonstrates how to move beyond Go's standard `net/http` into a high-performance, Express-inspired framework — with proper project layering, middleware, and test coverage.

---

## 📁 Directory Structure

```
fiber_framework/
├── config/
│   └── config.go          # PostgreSQL connection setup
├── handlers/
│   ├── car_handler.go      # HTTP handler functions (Fiber ctx)
│   ├── get_car_test.go     # Unit tests for GET endpoints
│   └── car_benchmarking_test.go  # Benchmark tests
├── middleware/
│   └── logger.go           # Custom logger (commented — Fiber's built-in used instead)
├── models/
│   └── cars.go             # Car struct + all DB operations
├── go.mod
└── main.go                 # App entry point, routes, and middleware setup
```

---

## ⚙️ How It Works

### Request Lifecycle

```
Client (HTTP request)
       ↓
Fiber Server  (app.Listen)
       ↓
Router  (main.go — method + path)
       ↓
Middleware  (logger → basicauth)
       ↓
Handler  (car_handler.go — business logic)
       ↓
Model  (cars.go — DB queries)
       ↓
PostgreSQL Database
       ↓
JSON Response → Client
```

---

## 🗄️ Database Setup

Connects to PostgreSQL using `lib/pq`. Update the DSN in `config/config.go` to match your local setup:

```go
dsn := "user=postgres password=YOUR_PASSWORD dbname=YOUR_DB host=localhost port=5432 sslmode=disable"
```

Create the `cars` table before running:

```sql
CREATE TABLE cars (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100),
    brand VARCHAR(100),
    year  INT,
    price NUMERIC(10, 2)
);
```

---

## 🔐 Authentication

Basic Auth is enabled globally via Fiber's built-in middleware. All routes require the following credentials:

| Username | Password |
|----------|----------|
| `Admin`  | `12345`  |

---

## 📡 API Endpoints

| Method   | Path        | Description           |
|----------|-------------|-----------------------|
| `GET`    | `/cars`     | Fetch all cars        |
| `GET`    | `/cars/:id` | Fetch a car by ID     |
| `POST`   | `/cars`     | Create a new car      |
| `PUT`    | `/cars/:id` | Update a car by ID    |
| `DELETE` | `/cars/:id` | Delete a car by ID    |

### Sample Request Body (POST / PUT)

```json
{
  "name": "Model S",
  "brand": "Tesla",
  "year": 2023,
  "price": 79999.99
}
```

---

## 🧪 Running Tests

Tests live in `handlers/` and use Fiber's `app.Test()` alongside `net/http/httptest`. They connect to the real PostgreSQL DB, inserting and cleaning up test data automatically.

```bash
cd fiber_framework

# Run all unit tests
go test ./handlers/... -v

# Run benchmarks
go test ./handlers/... -bench=. -benchmem
```

### What's tested

**Unit tests (`get_car_test.go`)**
- `GET /cars/:id` — valid ID returns 200 + correct car
- `GET /cars/:id` — non-existent ID returns 404
- `GET /cars/:id` — non-numeric ID returns 400
- `GET /cars` — returns 200 with at least one car

**Benchmark (`car_benchmarking_test.go`)**
- `BenchmarkGetById` — measures handler throughput for `GET /cars/:id`; seeds once, resets timer before the loop, and drains the response body on every iteration to avoid connection leaks

---

## 🚀 Running the Server

```bash
cd fiber_framework
go run .
```

Server starts on `http://localhost:8080`.

---

## 📚 Concepts Covered

- **Fiber v2** — fast HTTP framework built on `fasthttp`; `fiber.New()`, `fiber.Ctx`, `c.Params()`, `c.BodyParser()`, `c.JSON()`, `c.Status()`
- **Routing** — declarative method-specific routes (`app.Get`, `app.Post`, `app.Put`, `app.Delete`)
- **Middleware** — `logger.New()` for request logging; `basicauth.New()` for HTTP Basic Auth
- **Models layer** — DB logic fully separated from handlers; uses `database/sql` with parameterised queries
- **Error handling** — errors wrapped with `%w`, surfaced as JSON responses with proper HTTP status codes
- **Testing** — `app.Test()` for integration-style handler tests; helper functions (`assertNoErr`, `assertStatus`, `decodeJSON`) to eliminate boilerplate
- **Benchmarking** — `b.ResetTimer()` pattern, seed-once-before-loop, response body draining

> Made with 💙 while learning Go.