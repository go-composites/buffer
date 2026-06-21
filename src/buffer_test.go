package Buffer_test

import (
	"testing"

	Buffer "github.com/go-composites/buffer/src"
)

func TestNewIsEmpty(t *testing.T) {
	b := Buffer.New()
	if !b.IsEmpty() {
		t.Fatalf("New() should be empty")
	}
	if b.Len() != 0 {
		t.Fatalf("New() Len = %d, want 0", b.Len())
	}
	if b.ToGoString() != "" {
		t.Fatalf("New() ToGoString = %q, want empty", b.ToGoString())
	}
	if b.IsNull() {
		t.Fatalf("New() should not be null")
	}
}

func TestFromSeedsText(t *testing.T) {
	b := Buffer.From("seed")
	if b.ToGoString() != "seed" {
		t.Fatalf("From ToGoString = %q, want %q", b.ToGoString(), "seed")
	}
	if b.IsEmpty() {
		t.Fatalf("From(\"seed\") should not be empty")
	}
	if b.Len() != 4 {
		t.Fatalf("From Len = %d, want 4", b.Len())
	}
}

func TestAppendChainAndRune(t *testing.T) {
	b := Buffer.New().
		Append("Hello, ").
		Append("World").
		AppendRune('!')
	if got := b.ToGoString(); got != "Hello, World!" {
		t.Fatalf("chain ToGoString = %q, want %q", got, "Hello, World!")
	}
	if b.Len() != len("Hello, World!") {
		t.Fatalf("chain Len = %d, want %d", b.Len(), len("Hello, World!"))
	}
	if b.IsEmpty() {
		t.Fatalf("non-empty buffer reported empty")
	}
}

func TestAppendMutatesReceiver(t *testing.T) {
	b := Buffer.New()
	b.Append("a")
	b.Append("b")
	if got := b.ToGoString(); got != "ab" {
		t.Fatalf("mutation ToGoString = %q, want %q", got, "ab")
	}
}

func TestReset(t *testing.T) {
	b := Buffer.From("text").Reset()
	if !b.IsEmpty() {
		t.Fatalf("Reset should leave buffer empty")
	}
	if b.Len() != 0 {
		t.Fatalf("Reset Len = %d, want 0", b.Len())
	}
	if b.ToGoString() != "" {
		t.Fatalf("Reset ToGoString = %q, want empty", b.ToGoString())
	}
}

func TestNullBuffer(t *testing.T) {
	n := Buffer.Null()
	if !n.IsNull() {
		t.Fatalf("Null() should be null")
	}
	if !n.IsEmpty() {
		t.Fatalf("Null() should be empty")
	}
	if n.Len() != 0 {
		t.Fatalf("Null() Len = %d, want 0", n.Len())
	}
	if n.ToGoString() != "" {
		t.Fatalf("Null() ToGoString = %q, want empty", n.ToGoString())
	}
	if got := n.Append("x"); !got.IsNull() || got.ToGoString() != "" {
		t.Fatalf("Null().Append should be a no-op null")
	}
	if got := n.AppendRune('x'); !got.IsNull() || got.ToGoString() != "" {
		t.Fatalf("Null().AppendRune should be a no-op null")
	}
	if got := n.Reset(); !got.IsNull() || got.ToGoString() != "" {
		t.Fatalf("Null().Reset should be a no-op null")
	}
}
