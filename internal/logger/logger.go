package logger

import "igaopk.com/goPower/internal/config"

func LogMessage(message string, level string) {
	isDebug := config.GetEnv("DEBUG_MODE", "false")

	if !validateLevel(level) {
		panic("Invalid log level: " + level)
	}

	if isDebug == "true" {
		println(level, message)
	}
}
