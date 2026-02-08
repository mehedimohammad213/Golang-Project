package utils

import (
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	Logger = log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
