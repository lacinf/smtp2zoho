package email

import (
	"errors"
	"bytes"
	"encoding/json"
	"net/http"
	"smtp2zoho/config"
	"smtp2zoho/zoho"
)

// EmailPayload defines the structure of the email payload to be sent to the API.
type EmailPayload struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
}

// SendEmail sends an email using the configured API and Zoho access token.
// Logs all steps and errors using the centralized logger.
func SendEmail(cfg *config.Config, toAddress, subject, body string) error {
	// Obtain Zoho access token
	accessToken, err := zoho.GetAccessToken(cfg)
	if err != nil {
		config.Log(cfg, config.LogError, config.ErrEmailAccessToken, err)
		return err
	}

	// Prepare the email payload
	payload := EmailPayload{
		FromAddress: cfg.FromAddress,
		ToAddress:   toAddress,
		Subject:     subject,
		Content:     body,
	}

	// Serialize the payload to JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		config.Log(cfg, config.LogError, config.ErrEmailPayloadSerialization, err)
		return err
	}

	// Create the HTTP POST request to the API
	req, err := http.NewRequest("POST", cfg.APIURL, bytes.NewBuffer(jsonData))
	if err != nil {
		config.Log(cfg, config.LogError, config.ErrEmailCreateRequest, err)
		return err
	}

	req.Header.Set("Authorization", "Zoho-oauthtoken "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	// Send the request
	config.Log(cfg, config.LogInfo, config.InfoSMTPSendingToAPI, toAddress)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		config.Log(cfg, config.LogError, config.ErrEmailSendRequest, err)
		return err
	}
	defer resp.Body.Close()

	// Check for unsuccessful response status
	if resp.StatusCode != 200 {
		config.Log(cfg, config.LogError, config.ErrEmailSendFailed, resp.Status)
		return errors.New("send failed")
	}

	// Log successful email send
	config.Log(cfg, config.LogInfo, config.InfoEmailSent, toAddress, subject)
	return nil
}
