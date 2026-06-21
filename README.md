<p align="center"><img src="https://raw.githubusercontent.com/go-composites/brand/main/social/go-composites.png" alt="go-composites/buffer" width="720"></p>

# buffer

[![ci](https://github.com/go-composites/buffer/actions/workflows/ci.yml/badge.svg)](https://github.com/go-composites/buffer/actions/workflows/ci.yml)

The mutable text-buffer composite of [go-composites](https://github.com/go-composites):
a `StringBuilder`. Where the [`string`](https://github.com/go-composites/string)
composite is an *immutable* value, a `Buffer` is a *mutable* accumulator —
`Append`, `AppendRune` and `Reset` write in place and return the receiver, so
calls chain. It never goes `nil`: `Null()` yields a Null-Object `Buffer` whose
mutators are no-ops.

## Install

```sh
go get github.com/go-composites/buffer
```

## Usage

```golang
package main

import (
    "fmt"

    Buffer "github.com/go-composites/buffer/src"
)

func main() {
    b := Buffer.New().
        Append("Hello, ").
        Append("World").
        AppendRune('!')

    fmt.Println(b.ToGoString()) // Hello, World!
    fmt.Println(b.Len())        // 13
    fmt.Println(b.IsEmpty())    // false

    b.Reset()
    fmt.Println(b.IsEmpty())    // true
}
```

## API

### Constructors

- `New() Interface` — a new, empty `Buffer`.
- `From(s string) Interface` — a new `Buffer` seeded with initial text.
- `Null() Interface` — the Null-Object `Buffer`.

### Mutators (each mutates the receiver and returns it, so calls chain)

| method | meaning |
| --- | --- |
| `Append(s string)` | append text to the end of the buffer |
| `AppendRune(r rune)` | append a single rune |
| `Reset()` | clear the buffer |

### Conversion & predicates (return plain Go values)

| method | returns | notes |
| --- | --- | --- |
| `ToGoString()` | Go `string` | the accumulated text (`""` for the Null-Object) |
| `Len()` | Go `int` | length in bytes |
| `IsEmpty()` | Go `bool` | buffer holds no text |
| `IsNull()` | Go `bool` | `true` only for the Null-Object |

## Mutable, unlike `string`

`Buffer` is the mutable counterpart to the value-style
[`string`](https://github.com/go-composites/string) composite. A `string`
value never changes — operations return fresh values. A `Buffer` accumulates:
`Append` and friends modify the same underlying builder and hand back the
receiver, making it the right tool for incrementally building text without
allocating an intermediate value per concatenation.

## Null-Object

`Null()` returns the never-nil Null-Object `Buffer`: it holds no text, so
`ToGoString()` is `""`, `Len()` is `0`, `IsEmpty()` is `true`, and the mutators
`Append`, `AppendRune` and `Reset` are no-ops that return the null `Buffer`.
`IsNull()` reports `true` for it.

## License

BSD-3-Clause © the go-composites/buffer authors.
