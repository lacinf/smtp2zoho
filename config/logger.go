package config

import (
	"log"
)

// Log writes a formatted message to the log if the configured LogLevel allows it.
func Log(cfg *Config, level LogLevel, msg string, args ...interface{}) {
	if cfg.LogLevel >= level {
		prefix := ""
		switch level {
		case LogError:
			prefix = "[ERROR]"
		case LogInfo:
			prefix = "[INFO]"
		case LogDebug:
			prefix = "[DEBUG]"
		}

		log.Printf("%s "+msg, append([]interface{}{prefix}, args...)...)
	}
}
