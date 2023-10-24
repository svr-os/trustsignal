package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetTemplate retrieves the template details.
func (client *TrustSignalClient) GetTemplate(apiKey ...string) (*TemplateResponse, error) {
	url := fmt.Sprintf("%s/accounts/templates", client.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		client.Logger.Error("Failed to create new request for Get Template", "error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	key := client.APIKey
	if len(apiKey) > 0 {
		key = apiKey[0]
	}
	req.Header.Set("api_key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		client.Logger.Error("Failed to send request for Get Template", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Get Template API", "status_code", resp.StatusCode)
	}

	var response TemplateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode TemplateResponse", "error", err)
		return nil, err
	}

	return &response, nil
}

// CreateTemplate creates a new template.
func (client *TrustSignalClient) CreateTemplate(data CreateTemplateData, apiKey ...string) (*TemplateResponse, error) {
	url := fmt.Sprintf("%s/accounts/templates", client.BaseURL)
	body, err := json.Marshal(data)
	if err != nil {
		client.Logger.Error("Failed to marshal CreateTemplateData", "error", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		client.Logger.Error("Failed to create new request for Create Template", "error", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	key := client.APIKey
	if len(apiKey) > 0 {
		key = apiKey[0]
	}
	req.Header.Set("api_key", key)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		client.Logger.Error("Failed to send request for Create Template", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Create Template API", "status_code", resp.StatusCode)
	}

	var response TemplateResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode TemplateResponse", "error", err)
		return nil, err
	}

	return &response, nil
}
