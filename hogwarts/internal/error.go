package internal

import "fmt"

var (
	/*ErrStudentsNotFounds is use when trying to find a Student in database
	with is id but the id doesn't exist or is not found */
	ErrStudentsNotFounds = fmt.Errorf("student not found")
)


type logFunction func(string)

func HandleError(err error, msg string, log logFunction) {
	if err != nil {
		log(msg)
	}
}

