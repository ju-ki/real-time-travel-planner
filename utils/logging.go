package utils

import (
	"io"
	"log"
	"os"
	"time"
)

func LoggingSettings(logFile string) {
	currentDate := time.Now().Format("2006-01-02")
	logfile, err := os.OpenFile("logs/"+currentDate+"_"+logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	multiWriter := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(multiWriter)
}
