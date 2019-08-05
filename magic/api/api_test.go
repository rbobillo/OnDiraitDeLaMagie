package api

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/internal"
	"testing"
)

func TestGetWizards(t *testing.T) {
	err := GetWizards(nil, nil)

	if err != nil {
		internal.Error("No Error Expected")
	}
}

func TestInitMagic(t *testing.T) {
	err := InitMagic(nil)

	if err != nil {
		internal.Error("No Error Expected")
	}
}

func BenchmarkHamming(b *testing.B) {
}
