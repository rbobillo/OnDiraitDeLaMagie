package debug

import (
	"log"
	"os"
)
import "strings"
import "fmt"
func PrintEnv() {
	log.Println("coucou")

	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
