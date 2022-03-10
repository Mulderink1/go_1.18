package fuzz_test

import (
	"github.com/Mulderink1/go_1.18/fuzz"
	"testing"
)

func TestEqualFuzz(t *testing.T) {
	if fuzz.EqualFuzz([]byte{1, 2, 3}, []byte{1, 2, 3}) != true {
		t.Errorf("not equal")
	}
}

func FuzzEqualFuzz(f *testing.F) {
	f.Fuzz(func(t *testing.T, a []byte, b []byte) {
		fuzz.EqualFuzz(a, b)
	})
}
