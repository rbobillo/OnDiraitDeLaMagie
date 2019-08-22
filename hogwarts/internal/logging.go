package internal

import (
"log"
)


const (
	debugLvl = "\033[0mDEBUG\033[0m "
	errorLvl = "\033[91mERROR\033[0m "
	infoLvl  = "\033[95mINFO\033[0m  "
	warnLvl  = "\033[93mWARN\033[0m  "
)

func stdOutLog(level string, message string) {
	log.SetPrefix(level)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	_ = log.Output(3, message) // calldepth = 3, so we can find the original calling function/file
}

// Debug set debugs attributes
// to a message then
// print i
func Debug(msg string) {
	stdOutLog(debugLvl, msg)
}

// Error set errors attributes
// to a message then
// print it
func Error(msg string) {
	stdOutLog(errorLvl, msg)
}

// Info set infos attributes
// to a message then
// print it
func Info(msg string) {
	stdOutLog(infoLvl, msg)
}

// Warn set Warnings attributes
// to a message then
// print it
func Warn(msg string) {
	stdOutLog(warnLvl, msg)
}

