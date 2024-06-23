package main

import (
	"fmt"
	"testing"
)

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s\n", name)
}
func TestHello(t *testing.T) {
	got := Hello("John")
	if got != "Hello, John\n" {
		t.Errorf("Hello() = %q, want %q", got, "Hello, John")
	}
}
