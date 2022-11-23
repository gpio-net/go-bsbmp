package logger

import (
	"fmt"
	"os"
)

func Debugf(format string, values ...interface{}) {
	fmt.Printf(format, values...)
}

func Fatal(error error) {
	fmt.Printf("Fatal: %s\n", error)
	os.Exit(1)
}

func Notify(message string) {
	fmt.Printf("Notify: %s\n", message)
}

func Infof(format string, values ...interface{}) {
	fmt.Printf(format, values...)
}
