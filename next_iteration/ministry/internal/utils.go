package internal

import "log"

// FailOnError logs fatal on error
func FailOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
