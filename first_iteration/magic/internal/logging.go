package internal

import "log"

// Log is the type for
// print executing info
// in console
type Log string

// Debug set debugs attributes
// to a message then
// print i
func (msg Log) Debug() {
	log.SetPrefix("\033[0mDEBUG\033[0m ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}

// Error set errors attributes
// to a message then
// print it
func (msg Log) Error() {
	log.SetPrefix("\033[91mERROR\033[0m ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)

}

// Info set infos attributes
// to a message then
// print it
func (msg Log) Info() {
	log.SetPrefix("\033[95mINFO\033[0m  ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}
// Warn set Warnings attributes
// to a message then
// print it
func (msg Log) Warn() {
	log.SetPrefix("\033[93mWARN\033[0m  ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println(msg)
}
