# Agent Guidelines for aoc2025

This is a personal Advent of Code repository used to learn Go. The primary goal
is **learning**, not producing solutions. When assisting:

- Do NOT write code or produce solutions unless explicitly asked.
- Do NOT give puzzle hints or spoilers unless explicitly asked.
- DO review code for Go idioms and conventions when asked.
- DO point out better algorithmic approaches the user may be unaware of, by
  describing the pattern and why it is more efficient — without writing the code.
- DO give general Go hints about methods, syntax, algorithms and patterns.

---

## Repository Structure

```
.
├── go.mod          # module: aoc2025, no external deps
├── AGENTS.md
├── day01/main.go
├── dayXX/
│   ├── main.go
│   └── input.txt   # gitignored
└── ...
```

Each day is a standalone `package main` under `dayXX/`. There are no
shared packages, no external dependencies, and no test files.

---

## Build / Run Commands

Run a solution from within the day's directory:

```sh
cd dayXX
go run .
```

---

## Lint / Format

Use standard Go tooling:

```sh
# Format all files (always run before committing)
gofmt -w ./...

# Or with goimports (organizes imports too)
goimports -w ./...

# Static analysis
go vet ./...

# Full lint (if golangci-lint is installed)
golangci-lint run ./...
```

---

## Code Style Guidelines

### Package and Module Layout

- Every day is `package main` with a single `main.go`.
- The module is `aoc2025` with Go 1.26.1. Modern stdlib features are used freely
  (e.g., `slices`, `maps`, `for range n` over integers, built-in `max`/`min`).
- No external dependencies — stdlib only.

### Canonical `main()` Pattern

Every solution follows this structure exactly:

```go
func main() {
    data, err := os.ReadFile("input.txt")
    if err != nil {
        panic(err)
    }

    lines := strings.Split(strings.TrimSpace(string(data)), "\n")

    // Optional common data pre-processing

    fmt.Println("Part 1: ", part1(lines))
    fmt.Println("Part 2: ", part2(lines))
}
```

- `part1` and `part2` are pure functions: they take the pre-processed input and 
  return a value (typically `int` or `uint64`). They do not read files or mutate
  global state.

### Imports

- Single `import (...)` block; no manual blank-line grouping between stdlib
  packages — `gofmt`/`goimports` handles ordering automatically.
- Stdlib only; no third-party packages.

```go
import (
    "fmt"
    "os"
    "slices"
    "strconv"
    "strings"
)
```

### Naming Conventions

| Construct | Convention | Example |
|---|---|---|
| Functions | `camelCase`, lowercase first | `parseLine`, `findHighest` |
| Exported types | `PascalCase` | `Range`, `Problem` |
| Struct fields | lowercase (unexported) | `start`, `end`, `numbers` |
| Loop counters | single letters or short names | `i`, `j`, `x`, `y`, `n` |
| Accumulators | short and descriptive | `ans`, `answer`, `total` |
| Grid variables | `grid`, `rows`, `cols` | |

- Avoid overly verbose names. Go prefers short, clear names with context
  provided by the surrounding code.
- Receiver names should be 1–2 letters matching the type name (e.g., `r` for
  `Range`).

### Types

- Use `int` by default for counts, indices, and answers.
- Use `uint64` only when the puzzle involves large IDs or values that overflow
  `int` on 32-bit (e.g., seed ranges, large hashes).
- Use `[][]byte` for character grids — it avoids repeated string conversions and
  supports in-place mutation efficiently.
- Use `strings.Builder` for string accumulation in loops.
- Use custom structs to group related values (`Range{start, end uint64}`).

### Error Handling

All errors in puzzle solutions are handled with `panic`:

```go
data, err := os.ReadFile("input.txt")
if err != nil {
    panic(err)
}

n, err := strconv.Atoi(s)
if err != nil {
    panic(err)
}
```

This is intentional for this context — puzzle inputs are trusted, and panicking
on unexpected input is appropriate. Do not introduce `log.Fatal`, `os.Exit`, or
error return chains unless asked.

### Common Patterns

**Grid operations:**
- Represent grids as `[][]byte`, loaded by splitting `[]byte` on `\n`.
- Always `copyGrid` (deep copy rows) before mutating a grid in a second pass.

**Sorting and searching:**
- Use `sort.Slice` for custom ordering.
- Use `sort.Search` (binary search) on sorted slices when applicable.
- Use the `slices` package (`slices.Clone`, `slices.Reverse`, `slices.Sort`,
  `slices.Contains`) for cleaner slice operations (Go 1.21+).

**Integer range loops (Go 1.22+):**
```go
for i := range n {   // equivalent to: for i := 0; i < n; i++
    ...
}
```

**String parsing:**
- `strings.Fields(line)` — split on any whitespace.
- `strings.Split(line, ",")` — split on a delimiter.
- `strconv.Atoi` / `strconv.ParseUint` for string-to-number conversion.

**Multiple return values:**
Use multiple return values freely for parsers and helpers:
```go
func parseRange(s string) (int, int) { ... }
```

### What to Avoid

- Global mutable state.
- Importing packages that are not needed (`goimports` will flag unused imports).
- Overly deep nesting — extract helpers to flatten logic.
- Shadowing the `err` variable across unrelated operations in the same scope.
- Using `append` carelessly when the underlying array may be shared; use
  `slices.Clone` or copy explicitly when mutation isolation is needed.

---

## Go Learning Focus Areas

When reviewing code, consider flagging opportunities to use:

- **`maps` package** (`maps.Keys`, `maps.Clone`, `maps.Copy`) — Go 1.21+
- **`slices` package** — often cleaner than manual index manipulation
- **Binary search** (`sort.Search`) instead of linear scans on sorted data
- **BFS/DFS** patterns for grid traversal problems
- **Dynamic programming** (memoization with `map`) for overlapping subproblems
- **Two-pointer / sliding window** for range/subsequence problems
- **Bitmasking** for subset enumeration on small sets
- **Struct methods** vs. standalone functions when a type has multiple behaviors
- **Goroutines and channels** — not typically needed for AoC, but worth knowing
