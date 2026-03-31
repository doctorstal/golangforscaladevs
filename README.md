# Go for Scala Developers

> A practical tour of Go for developers coming from Scala. Each section maps to a runnable example in this repo.

---

## 1. Mutability — No `val`, No Problem

Scala leans hard on immutability: `val` by default, `var` only when you must. Go takes the opposite stance — **everything is mutable by default**, and the language gives you no compile-time safety net for it.

There are no `val`/`var` distinctions. There are compile-time constants (`const`), but they only work for primitive types and are evaluated at compile time. For everything else, mutation is expected and normal.

**Key takeaway:** Stop fighting mutation. Embrace explicit pointer semantics instead.

---

## 2. Pointers, Pass-by-Value, Pass-by-Reference

**`01_pointers_basics.go`** — Go is always pass-by-value. To share state, you pass a *pointer* (`*T`). Unlike C, you rarely need pointer arithmetic — the `unsafe` package exists but is almost never the right answer.

Structs are automatically dereferenced through `.`, so `ptr.field` works without `(*ptr).field`.

**`02_pointerspassbyvaluetocopy.go`** — Passing a struct without a pointer copies it, just like a Scala case class copy. Safe, simple, predictable.

**`03_pointerspassbyvaluetocopyfail.go`** — The trap: if your struct contains pointer fields (e.g., `[]*person`), copying the struct copies the pointer, not the value it points to. Mutating through the copy affects the original. Scala's `copy()` on case classes is deep-safe because values are immutable; Go's is not.

> In Scala you get a deep copy for free with immutable data structures. In Go, you own the copy semantics — be deliberate.

---

## 3. Type System: Structs, Receivers, Interfaces

**`04_typesystem.go`** — Go has no classes. Structs hold data; *receiver functions* attach behaviour. A pointer receiver (`func (p *person) Rename(...)`) mutates the original; a value receiver (`func (p person) copy() person`) works on a copy. Prefer pointer receivers — they're consistent and avoid silent copies on large structs.

Interfaces are **implicit** — no `implements` keyword. If your type has the right methods, it satisfies the interface. This is close to Scala's structural typing, but simpler and without macros.

Unlike Scala's `Option[T]`, Go has no built-in null-safety. A pointer can be `nil`, and calling a method on it panics unless you guard explicitly (see `NameOrDefault`). This is one of the sharpest edges coming from Scala.

**`05_typecomposition.go`** — Go has no inheritance. Instead, use **embedding**: embed a struct inside another to "inherit" its methods. It resembles Scala's trait mixin, but it's resolved at compile time with no dynamic dispatch surprises. You can still override an embedded method by defining one with the same name on the outer struct.

> Scala: traits + linearisation. Go: embedding + explicit delegation. Less magic, more control.

---

## 4. Back to `if` and `for` — No Monads Here

**`06_back_to_ifs_and_loops.go`** — Scala developers reach for `map`, `flatMap`, `filter`, and `fold` on every collection. Go has none of that in the standard library (pre-generics at least). You write `for` loops. That's it.

Go's `for` is the only loop construct — no `while`, no `do-while`. It covers every case:

```go
for i := 0; i < 5; i++ { }   // classic
for i < 5 { }                 // while-style
for { break }                 // infinite
for i, v := range slice { }   // range
```

The `slices` and `maps` packages (Go 1.21+) add helpers like `slices.Contains`, `slices.SortFunc`, and `maps.Keys`, but they don't reach the expressiveness of Scala's collections. You'll write more boilerplate — and that's intentional.

**`07_maploopexample.go`** — Maps have no default-value lookup built in. The idiomatic `v, ok := m[key]` two-value form is the equivalent of Scala's `Map.get(key): Option[V]`. Missing keys return the zero value, not a panic. The file also shows how to attach receiver methods to a named map type — a neat trick to simulate `getOrElse`.

> Scala: `list.map(f).filter(p).fold(z)(op)`. Go: write the loop. Both get the job done; Go just makes the control flow visible.

---

## 5. Error Handling — Explicit, Verbose, Intentional

**`08_error_handling.go`** — Scala uses exceptions (or `Try`/`Either`). Go uses multiple return values. Functions that can fail return `(result, error)`, and you check `err != nil` after every call.

There is no equivalent of `Try { }.getOrElse { }`. If you ignore an error, the compiler doesn't stop you — but a linter will. The convention is: **handle every error, every time**.

```go
res, err := canProduceError(false)
if err != nil {
    log.Fatal(err)
}
```

This is verbose compared to `for-comprehension` chains, but error paths are always visible in the code — no hidden exception propagation.

### Panic and Recovery — Not a Replacement for `try/catch`

**`5.1_panic_and_recovery.go`** — Go does have `panic`, and it does have a recovery mechanism. But they are **not** a general-purpose error handling strategy — don't reach for them the way you'd reach for `throw`/`catch` in Scala.

`panic` is for truly unrecoverable programmer mistakes: index out of bounds, nil dereferences, broken invariants. It unwinds the call stack and crashes the program unless intercepted.

`recover()` can catch a panic, but only when called **directly inside a deferred function**:

```go
func safeDiv(input int) (result int, err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("recovered from panic: %v", r)
        }
    }()
    result = mightPanic(input)
    return
}
```

The pattern converts a panic into an `error` — back to idiomatic Go. The legitimate use case is at **package or server boundaries**: you don't want one bad request to crash the whole process, so a top-level middleware recovers panics and returns a 500. Inside your own code, prefer returning `error` values from the start.

> Scala: `throw`/`catch` is normal control flow. Go: panicking is a last resort. If you find yourself writing `recover()` everywhere, you're fighting the language.

---

## 6. `defer` — Cleanup Without `finally`

**`09_defer.go`** — Go has no `try/finally`. Instead, `defer` schedules a function call to run when the enclosing function returns — regardless of whether it returns normally or panics.

```go
f, _ := os.Open("data.txt")
defer f.Close()   // always runs
```

Deferred calls execute in LIFO order. This maps loosely to Scala's `scala.util.Using` or `try { } finally { }`, but is more lightweight and idiomatic in Go for resource cleanup.

---

## 7. Fun Facts

**`10_fun_fact.go`** — You can specify array indices explicitly in a slice literal. Combined with `iota`-based enums, this gives you a type-safe, index-keyed array of names with zero boilerplate. Scala achieves similar things with `sealed trait` + `Map`, but Go's approach is a single array.

**`11_anonymous_parameters.go`** — Interface method parameters can be unnamed. A method signature `Bar(string)` is valid — the parameter exists but has no name. `_` (blank identifier) is still used throughout Go as in Scala to discard values.

---

## Key Philosophy Differences

| Topic | Scala | Go |
|---|---|---|
| Immutability | Default (`val`) | Opt-in (manual) |
| Null safety | `Option[T]` | Pointer + nil guard |
| Collections | Rich monad-based API | `for` loops + `slices`/`maps` packages |
| Error handling | Exceptions / `Either` | Multiple return values |
| OOP | Classes, traits, inheritance | Structs, interfaces, embedding |
| Concurrency | Futures, Akka | Goroutines, channels, `sync` |
| Philosophy | Expression-oriented, functional | Explicit, procedural, simple |

Go is deliberately minimal. It trades expressiveness for readability and fast compile times. Coming from Scala, the biggest adjustment isn't the syntax — it's accepting that boilerplate is often the right choice.

---

## Running the Examples

```bash
go run .
```

Uncomment individual function calls in `main.go` to run specific examples.
