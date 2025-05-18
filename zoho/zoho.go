package zoho

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"smtp2zoho/config"
	"sync"
	"time"
)

var (
	accessToken string
	expiresAt   time.Time
	mutex       sync.Mutex
)

// tokenResponse represents the JSON structure of Zoho's access token response.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

// GetAccessToken retrieves and caches the Zoho access token using the refresh token flow.
// It locks the token retrieval to ensure thread safety across concurrent calls.
func GetAccessToken(cfg *config.Config) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()

	// Return cached token if still valid
	if accessToken != "" && time.Now().Before(expiresAt) {
		return accessToken, nil
	}

	// Prepare request data for refresh token flow
	data := url.Values{}
	data.Set("refresh_token", cfg.RefreshToken)
	data.Set("client_id", cfg.ClientID)
	data.Set("client_secret", cfg.ClientSecret)
	data.Set("grant_type", "refresh_token")

	// Request new access token from Zoho
	resp, err := http.PostForm("https://accounts.zoho.com/oauth/v2/token", data)
	if err != nil {
		config.Log(cfg, config.LogError, config.ErrZohoRequestTokenFailed, err)
		return "", err
	}
	defer resp.Body.Close()

	// Check if response is not successful
	if resp.StatusCode != 200 {
		config.Log(cfg, config.LogError, config.ErrZohoInvalidStatus, resp.Status)
		return "", fmt.Errorf("invalid status: %s", resp.Status)
	}

	// Decode the access token response
	var tr tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tr); err != nil {
		config.Log(cfg, config.LogError, config.ErrZohoDecodeResponseFailed, err)
		return "", err
	}

	// Validate the access token field
	if tr.AccessToken == "" {
		config.Log(cfg, config.LogError, config.ErrZohoEmptyAccessToken)
		return "", fmt.Errorf("empty access token")
	}

	// Cache the access token and its expiration time (minus 60s for safety margin)
	accessToken = tr.AccessToken
	expiresAt = time.Now().Add(time.Duration(tr.ExpiresIn-60) * time.Second)

	return accessToken, nil
}
