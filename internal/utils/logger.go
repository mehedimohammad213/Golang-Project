package utils

import (
	"io"
	"log"
	"os"
)

var Logger *log.Logger

func InitLogger() {
	w := io.Writer(os.Stdout)
	if p := os.Getenv("LOG_FILE"); p != "" {
		f, err := os.OpenFile(p, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Printf("LOG_FILE=%s: %v (using stdout only)", p, err)
		} else {
			w = io.MultiWriter(os.Stdout, f)
		}
	}
	Logger = log.New(w, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func GetLogger() *log.Logger {
	if Logger == nil {
		InitLogger()
	}
	return Logger
}
