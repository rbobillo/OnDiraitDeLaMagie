package internal

type logFunction func(string)

func HandleError(err error, msg string, log logFunction) {
	if err != nil {
		log(msg)
	}
}
