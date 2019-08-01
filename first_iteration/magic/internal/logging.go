package internal

import "log"

type Log string

func (msg Log) Debug() {
	log.SetPrefix("\033[0mDEBUG\033[0m ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}

func (msg Log) Error() {
	log.SetPrefix("\033[91mERROR\033[0m ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}

func (msg Log) Info() {
	log.SetPrefix("\033[95mINFO\033[0m  ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}

func (msg Log) Warn() {
	log.SetPrefix("\033[93mWARN\033[0m  ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}
