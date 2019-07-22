package api

import "testing"

func TestProtectHogwarts(t *testing.T) {
	err := ProtectHogwarts(nil, nil, nil)

	if err != nil {
		t.Fatal("No Error Expected")
	}
}

func TestInitMagic(t *testing.T) {
	err := InitHogwarts(nil)

	if err != nil {
		t.Fatal("No Error Expected")
	}
}

func BenchmarkHamming(b *testing.B) {
}
