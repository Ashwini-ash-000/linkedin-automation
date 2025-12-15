package logger

import (
	"log"
)

type Logger struct{}

func New(_ interface{}) *Logger {
	return &Logger{}
}

func (l *Logger) Info(msg string) {
	log.Println("[INFO]", msg)
}

func (l *Logger) Warn(msg string) {
	log.Println("[WARN]", msg)
}

func (l *Logger) Error(msg string) {
	log.Println("[ERROR]", msg)
}
