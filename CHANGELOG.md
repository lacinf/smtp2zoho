# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]
### Fixed
- Adjusted SMTP server listening log from LogError back to LogInfo for correct severity alignment.

### Changed
- Aligned and cleaned up all messages in messages.go for better readability.
- Moved InfoSMTPSendingToAPI log from smtp.go to email.go, no longer logs prematurely during preparation.
- Corrected log levels in smtp.go:
  - Full message received and clean body received now correctly logged as Debug.
  - SMTP server listening now forced to LogError to always appear.
- Updated Dockerfile to use correct binary name smtp2zoho throughout.
- Updated go.mod module path to smtp2zoho.
- Updated all internal imports to smtp2zoho.
- Adjusted example.env default SMTP_AUTH_REQUIRED to false.

### Changed
- Cleaned up and reorganized project files.
- Moved example_zoho_test.sh to scripts/ (deleted from root).
- Moved smtp_test.sh and zoho_test.sh to test/ (ignored in .gitignore).
- Updated .gitignore to ignore .env, test/ folder and adjusted comments to English.
- Kept scripts/ folder tracked for example usage scripts.

### Changed
- Centralized all SMTP messages, email sending flow, and Zoho token flow logs into config.Log() with LogLevel control.
- Centralized all SMTP, Email, and Zoho-related user-facing messages into messages.go.
- Replaced all hardcoded Portuguese messages in smtp.go, email.go, and zoho.go with centralized English messages.
- Replaced all direct log calls and fmt.Errorf in smtp.go, email.go, and zoho.go.
- Configured SMTP_USER and SMTP_PASSWORD as optional, with fallback to "user"/"password" if not defined.
- Introduced SMTP_AUTH_REQUIRED as a boolean environment variable to control authentication requirement.
- Introduced SMTP_PORT as an environment variable with default to 25 if not defined.
- Ensured all configuration comments and code comments are now in clear, technical English.

### Changed
- Integrated email sending flow (`email.go`) with centralized config.Log() and LogLevel.
- Centralized all email-related error and info messages into messages.go.
- Replaced all hardcoded Portuguese messages in email.go with centralized English messages.
- Removed direct usage of fmt.Printf and fmt.Errorf for user-facing messages.
- Ensured returned errors are technical only, with all human-facing messages logged via config.Log().

### Changed
- Integrated Zoho access token retrieval (`zoho.go`) with centralized config.Log() using LogError level.
- Centralized all Zoho-related error messages into messages.go.
- Removed all hardcoded Portuguese messages from zoho.go.
- Ensured returned errors in GetAccessToken are technical only, with human-friendly messages handled exclusively via logs.

### Changed
- Startup email subject and body now include the executable name dynamically.
- Centralized startup email subject and body into messages.go.
- All log messages in main.go now use the centralized config.Log().
- Startup email sending now respects the configured LogLevel.
- Removed direct usage of the log package in main.go.

### Added
- Support for application log levels (Error, Info, Debug).
- Automatic `.env` file loading based on the executable name.
- Centralized all user-facing messages into `messages.go`.

### Changed
- Replaced `log.Fatal` with non-fatal logging and controlled application blocking during startup if critical environment variables are missing.
- Improved environment variable validation flow.
- Cleaned up unused imports and optimized code structure.

### Dependencies
- Added `github.com/joho/godotenv` for `.env` loading support.
