package appconfig

import (
	"os"

	"github.com/rs/zerolog"
)

var logLevel = getOrDefault("LOG_LEVEL", "INFO")
var DumpsPort = getOrDefault("HTTPDUMPS_PORT", "8801")

func GetLogLevel() zerolog.Level {
	switch logLevel {
	case "DEBUG":
		return zerolog.DebugLevel
	case "INFO":
		return zerolog.InfoLevel
	case "WARN":
		return zerolog.WarnLevel
	case "ERROR":
		return zerolog.ErrorLevel
	case "PANIC":
		return zerolog.PanicLevel
	default:
		return zerolog.InfoLevel
	}
}

func getOrDefault(varName string, defaultReturn string) string {
	if result := os.Getenv(varName); result == "" {
		return defaultReturn
	} else {
		return result
	}
}
