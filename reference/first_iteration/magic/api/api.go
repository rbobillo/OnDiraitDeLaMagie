package main

import (
  "fmt"
)

// GetWizards function requests the Magic Inventory
// to find wizards
func GetWizards() (string, error) {
  fmt.Println("Getting Wizards")
  return "", nil
}

// InitMagic starts the Magic service
func InitMagic() (string, error) {
  fmt.Println("Init Magic service")
  return "", nil
}

func main() {
  InitMagic()
}
