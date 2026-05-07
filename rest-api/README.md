# 🌐 REST API — Building HTTP APIs in Go

This module covers building RESTful APIs in Go from the ground up — starting with raw `net/http` and an in-memory store, then graduating to a structured, database-backed implementation with custom middleware.

Each sub-directory is a **standalone runnable program** representing a step up in complexity.

---

## 📁 Directory Structure

```
rest-api/
├── first_api/              # Step 1 — in-memory CRUD, pure net/http, no DB
│   ├── basic_first_api.go  # Handler logic, Car struct, in-memory map store
│   ├── main.go             # Entry point, route registration
│   └── go.mod
│
└── basic_CURD/             # Step 2 — PostgreSQL-backed CRUD, ServeMux, middleware
    ├── config/
    │   └── config.go       # PostgreSQL connection setup
    ├── handlers/
    │   └── car_handler.go  # CarHandler — routes dispatched by method + path
    ├── middleware/
    │   └── logger.go       # Custom Logger middleware
    ├── models/
    │   └── cars.go         # Car struct + all DB operations
    ├── main.go             # Entry point, ServeMux, middleware wiring
    └── go.mod
```

---

## 🗺️ Learning Path

| Step | Folder        | What You Learn |
|------|---------------|----------------|
| 1    | `first_api/`  | `net/http` basics, manual routing, in-memory store, `sync.Mutex` for concurrency safety |
| 2    | `basic_CURD/` | `http.ServeMux`, custom middleware, PostgreSQL with `database/sql`, layered architecture |

---

## 📦 first_api — In-Memory CRUD

A minimal REST API with **no database**. State lives in a `map[int64]Car` guarded by a `sync.Mutex`. Great for understanding the raw mechanics of `net/http` before adding any infrastructure.

### How routing works

Go's `net/http` doesn't support method-based routing out of the box. A single handler (`carHandler`) receives all requests to `/cars` and `/cars/`, then dispatches by inspecting `req.Method` and manually parsing the path:

```go
entity := strings.TrimPrefix(path, "/cars")
entity = strings.Trim(entity, "/")
// entity == ""     → /cars
// entity == "123"  → /cars/123
```

### Endpoints

| Method   | Path        | Description                        |
|----------|-------------|------------------------------------|
| `GET`    | `/cars`     | Return all cars (from memory)      |
| `GET`    | `/cars/:id` | Return a single car by ID          |
| `POST`   | `/cars`     | Add a car (random ID via `rand`)   |
| `DELETE` | `/cars/:id` | Remove a car by ID                 |

### Running

```bash
cd rest-api/first_api
go run .
```

Server starts on `http://localhost:8080`.

---

## 🗄️ basic_CURD — PostgreSQL-Backed CRUD

A structured CRUD API using `http.ServeMux`, a real PostgreSQL database, a custom logger middleware, and a clean separation of concerns across `config/`, `handlers/`, `models/`, and `middleware/`.

### Architecture

```
Client
  ↓
http.ListenAndServe  (main.go)
  ↓
middleware.Logger  (wraps the mux)
  ↓
http.ServeMux  (routes /cars and /cars/)
  ↓
handlers.CarHandler  (dispatches by method + parsed path)
  ↓
models  (DB queries via database/sql)
  ↓
PostgreSQL
  ↓
JSON Response → Client
```

### Database Setup

Update the DSN in `config/config.go` to match your local PostgreSQL setup:

```go
dsn := "user=postgres password=YOUR_PASSWORD dbname=YOUR_DB host=localhost port=5432 sslmode=disable"
```

Create the `cars` table:

```sql
CREATE TABLE cars (
    id    SERIAL PRIMARY KEY,
    name  VARCHAR(100),
    brand VARCHAR(100),
    year  INT,
    price NUMERIC(10, 2)
);
```

### Endpoints

| Method   | Path        | Description           |
|----------|-------------|-----------------------|
| `GET`    | `/cars`     | Fetch all cars from DB |
| `GET`    | `/cars/:id` | Fetch a car by ID     |
| `POST`   | `/cars`     | Insert a new car      |
| `PUT`    | `/cars/:id` | Update a car by ID    |
| `DELETE` | `/cars/:id` | Delete a car by ID    |

### Sample Request Body (POST / PUT)

```json
{
  "name": "Civic",
  "brand": "Honda",
  "year": 2022,
  "price": 24000.00
}
```

### Logger Middleware

The custom `middleware.Logger` wraps `http.Handler`, logs method + path on entry, and logs method + path + duration on exit:

```
Request started  - Method: GET,  Path: /cars
Request completed - Method: GET, Path: /cars, Duration: 3.2ms
```

### Running

```bash
cd rest-api/basic_CURD
go run .
```

Server starts on `http://localhost:8080`.

---

## 📚 Concepts Covered

### first_api
- `net/http` server setup with `http.HandleFunc` and `http.ListenAndServe`
- Manual path parsing with `strings.TrimPrefix` and `strings.Trim`
- Method-based dispatch inside a single handler using a `switch` on `req.Method`
- In-memory data store using `map[int64]Car`
- Concurrency safety with `sync.Mutex` (`Lock` / `Unlock` / `defer`)
- JSON encoding with `json.NewDecoder` and `json.NewEncoder`

### basic_CURD
- `http.NewServeMux` for explicit route registration
- Middleware pattern — wrapping `http.Handler` to add cross-cutting behaviour
- Custom `Logger` middleware measuring request duration with `time.Since`
- PostgreSQL connection via `database/sql` + `lib/pq` driver
- Parameterised queries (`$1`, `$2`, …) to prevent SQL injection
- `RETURNING id` pattern for getting the auto-generated primary key after insert
- `sql.ErrNoRows` handling for not-found cases
- `result.RowsAffected()` to confirm a delete actually hit a row
- Error wrapping with `fmt.Errorf("…: %w", err)` for traceability
- Layered project structure — `config`, `models`, `handlers`, `middleware`

> Made with 💙 while learning Go.