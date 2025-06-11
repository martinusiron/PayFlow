package middleware

import (
	"log"
	"os"
)

var (
	infoLog  = log.New(os.Stdout, "INFO: ", log.LstdFlags)
	errorLog = log.New(os.Stderr, "ERROR: ", log.LstdFlags)
)

func Info(msg string) {
	infoLog.Println(msg)
}

func Error(msg string) {
	errorLog.Println(msg)
}
