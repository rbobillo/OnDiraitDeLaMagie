package main

import "testing"

func TestGetWizards(t *testing.T) {
  _, err := GetWizards()

  if err != nil {
    t.Fatal("No Error Expected")
  }
}

func TestInitMagic(t *testing.T) {
  _, err := InitMagic()

  if err != nil {
    t.Fatal("No Error Expected")
  }
}

func BenchmarkHamming(b *testing.B) {
}
