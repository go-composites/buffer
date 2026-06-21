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

	b.Reset()
	fmt.Println(b.IsEmpty())    // true
	fmt.Println(b.ToGoString()) // (empty)
}
