# 🎨 Design Patterns in Go

A hands-on implementation of classic Gang of Four design patterns in Go — using interfaces, structs, and idiomatic Go patterns like `sync.Once`.

Each pattern lives in its own package and is invoked from `main.go` so you can run everything in one shot and see the output clearly.

---

## 📁 Directory Structure

```
design-pattern/
├── singleton/
│   └── singleton.go         # sync.Once-based singleton
├── factory/
│   └── factory_pattern.go   # Interface + switch-based factory
├── observer/
│   └── observer_pattern.go  # Subject/Observer interfaces, NewsAgency example
├── decorator/
│   └── decorator_pattern.go # Coffee/Milk decorator chain
├── go.mod
└── main.go                  # Runs all four patterns
```

---

## 🚀 Running

```bash
cd design-pattern
go run .
```

---

## 🧩 Patterns Covered

### 1. Singleton — `singleton/`

Guarantees only **one instance** of a struct is ever created, even under concurrent access.

**Key idea:** `sync.Once` ensures the initialization function runs exactly once regardless of how many goroutines call `getInstance()` simultaneously.

```go
var once sync.Once

func getInstance() *singleton {
    once.Do(func() {
        instance = &singleton{val: 50}
    })
    return instance
}
```

**When to use it:** Shared resources that should have exactly one owner — DB connection pools, config loaders, loggers.

> **Singleton vs Factory — the key distinction:**
> - **Singleton's goal** — ensure only one instance of a struct exists. No polymorphism needed; just a single concrete type.
> - **Factory's goal** — provide a flexible way to create *different* types of objects. Requires a common interface so the factory can return different concrete implementations without the caller knowing the details.

---

### 2. Factory — `factory/`

Provides a **single creation point** for a family of related types, hiding which concrete type gets returned behind a common interface.

**Key idea:** The caller asks for a `Vehicle` by name; the factory decides which concrete struct (`car` or `bike`) to return. The caller only ever talks to the `Vehicle` interface.

```go
type Vehicle interface {
    Drive()
}

func GetVehicle(vType string) Vehicle {
    switch vType {
    case "car":  return car{}
    case "bike": return bike{}
    default:     return nil
    }
}
```

**When to use it:** When you want to decouple object creation from usage — adding a new type (e.g. `truck`) requires only a new struct + one new `case`, with zero changes to caller code.

---

### 3. Observer — `observer/`

Defines a **publish/subscribe** relationship: a Subject broadcasts events, and any number of Observers react independently — without the Subject knowing their implementation details.

```
Subject (publisher) maintains a list of Observers (subscribers).
When something happens, it notifies all of them.
Each Observer reacts in its own way — the Subject doesn't care how.
```

**Step-by-step implementation:**

1. **Define the interfaces** — `Observer` declares `Update(msg string)` so every observer knows how to react. `Subject` declares `Register` and `NotifyAll` so every subject knows how to manage its observers.
2. **Concrete Subject** (`newsAgency`) — holds a `[]Observer` slice. `Register` appends to it; `NotifyAll` loops through and calls `Update` on each.
3. **Concrete Observers** (`SMSClient`, `EmailClient`) — each implements `Update` differently; the subject never needs to know the difference.
4. **Client code** — create a `newsAgency`, register observers, call `NotifyAll`.

**Workflow:**

```
Subject (NewsAgency)
    │
    ├── Register(EmailClient)
    ├── Register(SMSClient)
    │
    └── NotifyAll("Breaking news!")
            ├── EmailClient.Update() → "Email received: ..."
            └── SMSClient.Update()  → "SMS received: ..."
```

**Key idea:** Two interfaces keep things decoupled — `Observer` declares `Update(msg string)`, and `Subject` declares `Register` + `NotifyAll`. Concrete types implement them independently.

```go
type Observer interface { Update(string) }
type Subject  interface {
    Register(observer Observer)
    NotifyAll(msg string)
}
```

**When to use it:** Event-driven systems — notifications, UI event listeners, logging pipelines, anything where one event should trigger multiple independent reactions.

---

### 4. Decorator — `decorator/`

Adds **new behaviour to an existing object** by wrapping it in another struct that satisfies the same interface — without modifying the original type.

**Key idea:** Both `Espresso` and `Milk` implement the `Coffee` interface. `Milk` wraps an `Espresso`, delegates the base cost, and adds its own on top. You can chain decorators indefinitely.

```go
type Coffee interface {
    Cost() int
    Description() string
}

type Milk struct{ espreso Espresso }

func (m Milk) Cost() int { return m.espreso.Cost() + 70 }
```

```
Espresso → Cost: 250
Milk(Espresso) → Cost: 250 + 70 = 320
```

**When to use it:** When you want to extend behaviour at runtime without subclassing — HTTP middleware chains, I/O stream wrappers, feature toggles layered on top of a base component.

---

## 📚 Pattern Comparison

| Pattern   | Intent | Go Mechanism |
|-----------|--------|--------------|
| Singleton | One instance only | `sync.Once` + package-level var |
| Factory   | Flexible object creation | Interface return + `switch` |
| Observer  | Event broadcasting | Two interfaces + slice of observers |
| Decorator | Add behaviour without modification | Interface wrapping + struct composition |

> Made with 💙 while learning Go.