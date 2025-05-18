package smtp

import (
	"errors"
	"bytes"
	"log"
	"net"
	"smtp2zoho/config"
	"smtp2zoho/email"

	"github.com/mhale/smtpd"
	"github.com/jhillyerd/enmime"
)

// StartSMTP initializes and starts the SMTP server using the provided configuration.
func StartSMTP(cfg *config.Config) error {
	handler := func(origin net.Addr, from string, to []string, data []byte) error {
		// Log SMTP connection and message data
		config.Log(cfg, config.LogInfo, config.InfoSMTPConnection, origin.String(), from, to)
		config.Log(cfg, config.LogDebug, config.InfoSMTPMessageReceived, string(data))

		// Parse the MIME message
		env, err := enmime.ReadEnvelope(bytes.NewReader(data))
		if err != nil {
			config.Log(cfg, config.LogError, config.ErrSMTPMimeParseFailed, err)
			return err
		}

		// Log all headers
		for k, v := range env.Root.Header {
			config.Log(cfg, config.LogDebug, config.InfoSMTPHeader, k, v)
		}

		// Extract subject and log warnings if missing
		subject := env.GetHeader("Subject")
		if subject == "" {
			config.Log(cfg, config.LogError, config.WarnSMTPSubjectEmpty)
		} else {
			config.Log(cfg, config.LogInfo, config.InfoSMTPSubjectReceived, subject)
		}

		// Extract body preferring HTML, fallback to Text
		body := env.HTML
		if body == "" {
			config.Log(cfg, config.LogInfo, config.WarnSMTPBodyHTMLFallback)
			body = env.Text
		}

		config.Log(cfg, config.LogDebug, config.InfoSMTPBodyReceived, body)

		// For each recipient, send the parsed email to the API
		for _, recipient := range to {
			config.Log(cfg, config.LogDebug, config.InfoSMTPPreparingAPI, recipient, subject)
			err := email.SendEmail(cfg, recipient, subject, string(body))
			if err != nil {
				config.Log(cfg, config.LogError, config.ErrSMTPEmailSendFailed, err)
			}
		}
		return nil
	}

	authHandler := func(remoteAddr net.Addr, mechanism string, username []byte, password []byte, _ []byte) (bool, error) {
		// Log authentication attempt
		config.Log(cfg, config.LogInfo, config.InfoSMTPAuthAttempt, remoteAddr.String(), mechanism)

		if mechanism != "PLAIN" {
			config.Log(cfg, config.LogError, config.ErrSMTPAuthUnsupported, mechanism)
			return false, errors.New("unsupported mechanism")
		}

		user := string(username)
		pass := string(password)

		// Log received credentials (for debugging purposes, use caution in production)
		log.Printf("[SMTP] AUTH PLAIN recebido - User: %s, Pass: %s", user, pass)

		if user == cfg.SMTPUser && pass == cfg.SMTPPassword {
			config.Log(cfg, config.LogInfo, config.InfoSMTPAuthPlainSuccess, user)
			return true, nil
		}

		config.Log(cfg, config.LogError, config.ErrSMTPAuthFailed, user)
		return false, errors.New("invalid credentials")
	}

	// Initialize and start the SMTP server
	server := &smtpd.Server{
		Addr:  "0.0.0.0:" + cfg.SMTPPort,
		Handler:      handler,
		Hostname:     "localhost",
		AuthHandler:  authHandler,
		AuthMechs:    map[string]bool{"PLAIN": true},
		AuthRequired: false,
	}

	// Log that the server is starting and listening on the configured port
	config.Log(cfg, config.LogInfo, config.InfoSMTPListening, cfg.SMTPPort)
	return server.ListenAndServe()
}
