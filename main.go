package main

import (
	"fmt"
	"os"
	"path/filepath"
	"smtp2zoho/config"
	"smtp2zoho/email"
	"smtp2zoho/smtp"
	"flag"
)

func main() {
	
	showVersion := flag.Bool("version", false, config.VersionFlagDescription)
	flag.Parse()

	if *showVersion {
		fmt.Println(config.Version)
		os.Exit(0)
	}

	// Load configuration
	cfg := config.Load()

	// Block startup if configuration is invalid
	if cfg == nil {
		select {}
	}

	// Start SMTP server in a goroutine
	go func() {
		if err := smtp.StartSMTP(cfg); err != nil {
			config.Log(cfg, config.LogError, config.ErrSMTPStartFailed, err)
			select {} // Block without exiting the process
		}
	}()

	// Send startup notification email (only if LogLevel is Info or higher)
	if cfg.LogLevel >= config.LogInfo {
		execName := filepath.Base(os.Args[0])
		err := email.SendEmail(
			cfg,
			cfg.FromAddress,
			fmt.Sprintf(config.StartupEmailSubject, execName),
			fmt.Sprintf(config.StartupEmailBody, execName),
		)
		if err != nil {
			config.Log(cfg, config.LogError, config.ErrStartupEmailFailed, err)
		}
	}

	// Block main thread to keep the application running
	select {}
}
