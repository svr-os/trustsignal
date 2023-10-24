package logger

import (
	"log"
)

type DefaultLogger struct{}

func (l *DefaultLogger) Debug(args ...interface{}) {
	log.Println("[DEBUG]", args)
}

func (l *DefaultLogger) Info(args ...interface{}) {
	log.Println("[INFO]", args)
}

func (l *DefaultLogger) Warn(args ...interface{}) {
	log.Println("[WARN]", args)
}

func (l *DefaultLogger) Error(args ...interface{}) {
	log.Println("[ERROR]", args)
}
