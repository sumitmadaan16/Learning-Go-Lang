# 🧠 Learning-GoLang

A modular, hands-on repository to master Go (Golang) — from the very basics through to concurrency patterns, RESTful APIs, and a production-style framework layer.

Each folder is a **standalone runnable Go program** focused on one topic. You can jump straight to any section that interests you.

---

## 📁 Repository Structure

```
go-learning/
├── basics/                # Variables, constants, scope, pointers
├── interfaces/            # Type assertions, type switches
├── error-handling/        # Errors, custom errors, panic & recover
├── concurrency/           # Goroutines, channels, mutexes, worker pool
├── rest-api/              # net/http from scratch → structured CRUD with PostgreSQL
│   ├── first_api/         # In-memory CRUD, pure net/http, sync.Mutex
│   └── basic_CURD/        # PostgreSQL, ServeMux, custom middleware, layered arch
└── fiber_framework/       # Fiber v2 REST API, Basic Auth, unit tests, benchmarks
    ├── config/
    ├── handlers/
    ├── middleware/
    └── models/
```

---

## 🗺️ Learning Path

Follow this order if you're starting from scratch:

| Step | Folder | Topics |
|------|--------|--------|
| 1 | [`basics/`](./basics/) | Hello world, variable types, constants, scope, closures, pointers |
| 2 | [`interfaces/`](./interfaces/) | `interface{}`, type assertions, type switches |
| 3 | [`error-handling/`](./error-handling/) | `error`, custom errors, wrapping, `panic` / `recover` |
| 4 | [`concurrency/`](./concurrency/) | Goroutines, WaitGroups, Mutex, RWMutex, channels, worker pool, ticker |
| 5 | [`rest-api/`](./rest-api/) | Raw `net/http`, in-memory store, PostgreSQL, custom middleware |
| 6 | [`fiber_framework/`](./fiber_framework/) | Fiber v2, Basic Auth, structured routing, unit + benchmark tests |

---

## 🚀 How to Run Any Section

Make sure [Go is installed](https://go.dev/dl/) (1.22+), then:

```bash
# Pick any section
cd basics
go run .

cd ../concurrency
go run .

cd ../rest-api/first_api
go run .

cd ../../fiber_framework
go run .
```

Each folder has its own `go.mod` and `README.md` explaining what it covers.

---

## 📚 Topics Covered So Far

### ✅ Basics
- Primitive types: `int`, `float`, `bool`
- Composite types: arrays, slices, maps, structs
- Constants and `iota` enumerations
- Variable scope: package, function, block, loop
- Closures
- Pointers and pointer-to-pointer

### ✅ Interfaces
- Empty interface (`interface{}`)
- Safe type assertions with `(value, ok)` pattern
- Type switches

### ✅ Error Handling
- `errors.New`, multi-return error pattern
- Error wrapping with `%w`, `errors.Unwrap`, `errors.Is`
- Custom error types via struct + `error` interface
- `panic`, `recover`, stack unwinding with `defer`

### ✅ Concurrency
- Goroutines and the Go scheduler
- `sync.WaitGroup` for goroutine synchronisation
- `sync.Mutex` — mutual exclusion for shared state
- `sync.RWMutex` — concurrent reads, exclusive writes
- Unbuffered channels (synchronous handshake)
- Buffered channels (async up to capacity)
- Directional channels (`chan<-`, `<-chan`)
- Worker pool pattern (goroutines + channels + WaitGroup)
- `time.Ticker` for periodic events

### ✅ REST API (`rest-api/`)
- `net/http` server with `http.HandleFunc` and `http.ListenAndServe`
- Manual path parsing and method-based dispatch
- In-memory store with `sync.Mutex` for concurrency safety
- `http.NewServeMux` for explicit route registration
- Custom Logger middleware wrapping `http.Handler`
- PostgreSQL integration via `database/sql` + `lib/pq`
- Parameterised queries, `RETURNING id`, `sql.ErrNoRows`, `RowsAffected`
- Layered project structure: `config`, `models`, `handlers`, `middleware`

### ✅ Fiber Framework (`fiber_framework/`)
- Fiber v2 — `fiber.New()`, `fiber.Ctx`, declarative method routing
- Built-in middleware: `logger.New()`, `basicauth.New()`
- Handlers using `c.Params()`, `c.BodyParser()`, `c.JSON()`, `c.Status()`
- PostgreSQL-backed models with error wrapping
- Unit tests with `app.Test()` and `net/http/httptest`
- Test helpers to eliminate boilerplate (`assertNoErr`, `assertStatus`, `decodeJSON`)
- Benchmark tests with `b.ResetTimer()` and proper connection cleanup

---

## 🛣️ Upcoming

- [ ] Design patterns in Go
- [ ] Database migrations
- [ ] JWT-based authentication

---

## 🤝 Contributing / Following Along

This repo is open-sourced for learning purposes. Feel free to explore, fork, and experiment. Each file is heavily commented to explain the *why*, not just the *what*.

> Made with 💙 while learning Go.