package main

import "log"

// Checkpanic error
func CheckPanic(e error) {
	if e != nil {
		log.Panic(e)
	}
}

func CheckError(e error) {
	if e != nil {
		log.Print(e)
	}
}
