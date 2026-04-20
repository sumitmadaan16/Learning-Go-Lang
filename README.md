# 🧠 Learning-GoLang

A modular, hands-on repository to master Go (Golang) — from the very basics through to concurrency patterns, RESTful APIs, and a capstone project.

Each folder is a **standalone runnable Go program** focused on one topic. You can jump straight to any section that interests you.

---

## 📁 Repository Structure

```
go-learning/
├── basics/            # Variables, constants, scope, pointers
├── interfaces/        # Type assertions, type switches
├── error-handling/    # Errors, custom errors, panic & recover
├── concurrency/       # Goroutines, channels, mutexes, worker pool
└── rest-api/          # 🚧 Coming soon — RESTful API with net/http
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
| 5 | [`rest-api/`](./rest-api/) | 🚧 Coming soon |

---

## 🚀 How to Run Any Section

Make sure [Go is installed](https://go.dev/dl/) (1.22+), then:

```bash
# Pick any section
cd basics
go run .

cd ../concurrency
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

---

## 🛣️ Upcoming

- [ ] RESTful API with `net/http`
- [ ] Routing, middleware, JSON handling
- [ ] Database integration

---

## 🤝 Contributing / Following Along

This repo is open-sourced for learning purposes. Feel free to explore, fork, and experiment. Each file is heavily commented to explain the *why*, not just the *what*.

> Made with 💙 while learning Go.