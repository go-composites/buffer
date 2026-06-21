package Buffer

import (
	"strings"
)

// Interface is the mutable text-buffer composite of go-composites: a
// StringBuilder. Unlike the immutable `string` value composite, a Buffer
// accumulates text in place — Append, AppendRune and Reset mutate the
// receiver and return it, so calls chain. It never goes nil: Null() yields
// a Null-Object Buffer that honours the full Interface.
type Interface interface {
	Append(s string) Interface
	AppendRune(r rune) Interface
	ToGoString() string
	Len() int
	IsEmpty() bool
	Reset() Interface
	IsNull() bool
}

type data struct {
	builder strings.Builder
}

/*
New returns a new, empty Buffer.Interface.
*/
func New() Interface {
	return &data{}
}

/*
From returns a new Buffer.Interface seeded with the given initial text.
*/
func From(s string) Interface {
	d := &data{}
	d.builder.WriteString(s)
	return d
}

/*
Append writes s to the end of the buffer, mutating it, and returns the
receiver so calls chain.
*/
func (d *data) Append(s string) Interface {
	d.builder.WriteString(s)
	return d
}

/*
AppendRune writes a single rune to the end of the buffer, mutating it, and
returns the receiver so calls chain.
*/
func (d *data) AppendRune(r rune) Interface {
	d.builder.WriteRune(r)
	return d
}

/*
ToGoString returns the accumulated text as a Go string.
*/
func (d *data) ToGoString() string {
	return d.builder.String()
}

/*
Len returns the length of the accumulated text in bytes.
*/
func (d *data) Len() int {
	return d.builder.Len()
}

/*
IsEmpty reports whether the buffer holds no text.
*/
func (d *data) IsEmpty() bool {
	return d.builder.Len() == 0
}

/*
Reset clears the buffer, mutating it, and returns the receiver so calls chain.
*/
func (d *data) Reset() Interface {
	d.builder.Reset()
	return d
}

/*
IsNull reports that this is a real (non-null) Buffer.
*/
func (d *data) IsNull() bool {
	return false
}

// null is the Null-Object variant of a Buffer: a placeholder that honours the
// full Interface without ever being nil. It holds no text; its mutators are
// no-ops that return the null buffer.
type null struct{}

/*
Null returns the Null-Object Buffer.
*/
func Null() Interface {
	return &null{}
}

// Append is a no-op returning the null Buffer.
func (n *null) Append(s string) Interface { return n }

// AppendRune is a no-op returning the null Buffer.
func (n *null) AppendRune(r rune) Interface { return n }

// ToGoString returns the empty string for the null Buffer.
func (n *null) ToGoString() string { return "" }

// Len is 0 for the null Buffer.
func (n *null) Len() int { return 0 }

// IsEmpty is true for the null Buffer.
func (n *null) IsEmpty() bool { return true }

// Reset is a no-op returning the null Buffer.
func (n *null) Reset() Interface { return n }

// IsNull reports that this is the null Buffer.
func (n *null) IsNull() bool { return true }
