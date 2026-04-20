# ⚡ Concurrency

Covers Go's concurrency primitives — goroutines, channels, mutexes, and real-world patterns like the worker pool.

## Concepts Covered

| File | Concept |
|------|---------|
| `goroutines.go` | Launching goroutines, why `time.Sleep` is needed without synchronisation |
| `wait_groups.go` | `sync.WaitGroup` — the idiomatic way to wait for goroutines |
| `basic_mutex.go` | `sync.Mutex` — preventing data races on shared state |
| `read_write_mutex.go` | `sync.RWMutex` — allowing concurrent reads with exclusive writes |
| `unbuffered_channel.go` | Unbuffered channels — sender and receiver must meet (perfect handshake) |
| `buffered_channel.go` | Buffered channels — sender doesn't block until the buffer is full |
| `channel_direction.go` | Directional channels `chan<-` and `<-chan` — express intent, prevent misuse |
| `worker_pool.go` | 🏆 Capstone pattern: fixed worker pool using goroutines + channels + WaitGroup |
| `time_ticker.go` | `time.Ticker` — periodic events using channels |

## How to Run

```bash
cd concurrency
go run .
```

> Note: `GoRoutineDemo` uses `time.Sleep(12s)` intentionally to let you observe the goroutine counting in real time. This is expected behaviour, not a bug.