package logger

import (
	"fmt"
	"os"
)

func Debugf(format string, values ...interface{}) {
	fmt.Printf("Debug: "+format+"\r\n", values...)
}

func Fatal(error error) {
	fmt.Printf("Fatal: %s\r\n", error)
	os.Exit(1)
}

func Notify(message string) {
	fmt.Printf("Notify: %s\r\n", message)
}

func Infof(format string, values ...interface{}) {
	fmt.Printf("Info: "+format+"\r\n", values...)
}
