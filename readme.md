# Linked list

- Test: `go test ./...`
- Benchmark:
  1. `cd list`
  2. `go test -cpuprofile cpu.prof -memprofile mem.prof -benchmem -count 2 -benchtime=50000x -bench .`
- Read profiles:
  - cpu.prof:
    1. `go tool pprof cpu.prof`
    2. `pdf` | `svg` | `png`
  - mem.prof:
    1. `go tool pprof mem.prof`
    2. `pdf` | `svg` | `png`
