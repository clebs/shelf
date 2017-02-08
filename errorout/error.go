// Package errorout provides error output utilities
package errorout

import (
	"fmt"
	"os"
)

// Panic checks if there is an error and if there is, it panics
func Panic(err error, message string) {
	if err != nil {
		panic(message + ": " + err.Error())
	}
}

//ErrQuit checks if there is an error, if there is it prints it out and exits the program
func ErrQuit(err error, message string) {
	if err != nil {
		fmt.Printf("Error: %s (%v)", message, err)
		os.Exit(2)
	}
}
