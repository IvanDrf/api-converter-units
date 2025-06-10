package logger

import (
	"fmt"
	"log/slog"
	"os"
)

type Logger struct {
	l *slog.Logger
}

func (log *Logger) InfoLevel(method, msg string) {
	log.l.Info(fmt.Sprintf("%s: %s", method, msg))
}

func (log *Logger) ErrorLevel(method, msg string) {
	log.l.Info(fmt.Sprintf("%s: %s", method, msg))
}

func (log *Logger) Fatal(msg string) {
	log.l.Error(msg)
	os.Exit(1)
}

func InitLogger() Logger {
	return Logger{slog.New(slog.NewJSONHandler(os.Stdout, nil))}
}
