# ⚠️ Error Handling

Covers Go's idiomatic approach to errors — from basic error values to custom types and panic/recover.

## Concepts Covered

| File | Concept |
|------|---------|
| `error_handling.go` | `errors.New`, error returns, wrapping with `%w`, `errors.Unwrap`, `errors.Is` |
| `custom_errors.go` | Implementing the `error` interface via a struct + pointer receiver |
| `panic_and_recover.go` | `panic`, `recover` inside `defer`, stack unwinding |

## How to Run

```bash
cd error-handling
go run .
```

> To see stack unwinding, uncomment `StackUnwindingDemo()` in `main.go` — note it will crash the program intentionally.