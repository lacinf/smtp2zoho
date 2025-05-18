package config

import (
	"log"
	"os"
	"path/filepath"
	"github.com/joho/godotenv"
)

// Version holds the application version, injected via build flags.
var Version = "dev"

// LogLevel defines the application log level.
type LogLevel int

const (
	LogError LogLevel = iota	// Logs only startup message and critical errors during runtime.
	LogInfo						// Logs everything from LogError plus startup notification sent to ZOHO_FROM_ADDRESS and basic message delivery info.
	LogDebug					// Logs everything from LogInfo plus detailed steps of the application flow for debugging.
)

// Config holds all environment-based configuration parameters.
type Config struct {
	APIURL       string
	FromAddress  string
	ClientID     string
	ClientSecret string
	RefreshToken string
	SMTPUser     string
	SMTPPassword string
	LogLevel     LogLevel
	SMTPPort string
	SMTPAuthRequired bool
}

// getEnv retrieves the environment variable by key.
// If required is true and the variable is missing, it logs an error and marks the configuration as invalid.
func getEnv(key string, required bool, missing *bool) string {
	value := os.Getenv(key)
	if value == "" && required {
		log.Printf(ErrMissingEnvVar, key)
		*missing = true
	}
	return value
}

// getEnvBool retrieves a boolean environment variable by key.
// Returns defaultVal if the variable is not defined or empty.
func getEnvBool(key string, defaultVal bool) bool {
	val := os.Getenv(key)
	if val == "" {
		return defaultVal
	}
	return val == "true" || val == "1"
}

// getLogLevel parses the LOG_LEVEL environment variable and returns the corresponding LogLevel.
// Defaults to LogInfo if not set or invalid.
func getLogLevel() LogLevel {
	switch os.Getenv("LOG_LEVEL") {
	case "debug":
		return LogDebug
	case "error":
		return LogError
	default:
		return LogInfo
	}
}

// Load initializes and returns a Config instance populated from environment variables.
func Load() *Config {
	// Discover executable name and attempt to load {exec_name}.env
	execPath, err := os.Executable()
	if err == nil {
		execName := filepath.Base(execPath)
		envFile := execName + ".env"
		if err := godotenv.Load(envFile); err == nil {
			log.Printf(InfoLoadedEnvFile, envFile)
		}
	}

	missing := false

	// Get SMTP port, fallback to 25 if not defined
	port := getEnv("SMTP_PORT", false, &missing)
	if port == "" {
		port = "25"
	}

	cfg := &Config{
		APIURL:       getEnv("ZOHO_API_URL", true, &missing),
		FromAddress:  getEnv("ZOHO_FROM_ADDRESS", true, &missing),
		ClientID:     getEnv("ZOHO_CLIENT_ID", true, &missing),
		ClientSecret: getEnv("ZOHO_CLIENT_SECRET", true, &missing),
		RefreshToken: getEnv("ZOHO_REFRESH_TOKEN", true, &missing),
		SMTPUser:     getEnv("SMTP_USER", false, &missing),
		SMTPPassword: getEnv("SMTP_PASSWORD", false, &missing),
		SMTPAuthRequired: getEnvBool("SMTP_AUTH_REQUIRED", false),
		SMTPPort: port,
		LogLevel:     getLogLevel(),
	}

	// Apply default credentials if not set
	if cfg.SMTPUser == "" {
		cfg.SMTPUser = "user"
	}
	if cfg.SMTPPassword == "" {
		cfg.SMTPPassword = "password"
	}

	// Block startup if any required variables are missing
	if missing {
		log.Println(ErrStartupBlocked)
		return nil
	}

	return cfg
}
