package logger

import "log"

func Error(format string, args ...any) {
	log.Fatalf(format, args...)
}
