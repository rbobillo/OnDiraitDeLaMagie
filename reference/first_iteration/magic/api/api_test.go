package api

import "testing"

func TestGetWizards(t *testing.T) {
	err := GetWizards(nil, nil)

	if err != nil {
		t.Fatal("No Error Expected")
	}
}

func TestInitMagic(t *testing.T) {
	err := InitMagic(nil)

	if err != nil {
		t.Fatal("No Error Expected")
	}
}

func BenchmarkHamming(b *testing.B) {
}
