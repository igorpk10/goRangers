package logger

import (
	"fmt"
	"strings"

	"igaopk.com/goPower/internal/config"
)

func LogMessage(message string, level string, args ...interface{}) {
	isDebug := config.GetEnv("DEBUG_MODE")

	if !validateLevel(level) {
		panic("Invalid log level: " + level)
	}

	if isDebug == "true" {
		formatted := fmt.Sprintf(message, args...)
		fmt.Printf("[%s] %s\n", strings.ToUpper(level), formatted)
	}
}
