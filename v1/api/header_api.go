package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// GetHeader retrieves the header details.
func (client *TrustSignalClient) GetHeader(apiKey ...string) (*HeaderResponse, error) {
	url := fmt.Sprintf("%s/accounts/senders", client.BaseURL)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		client.Logger.Error("Failed to create new request for Get Header", "error", err)
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
		client.Logger.Error("Failed to send request for Get Header", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Get Header API", "status_code", resp.StatusCode)
	}

	var response HeaderResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode HeaderResponse", "error", err)
		return nil, err
	}

	return &response, nil
}

// CreateHeader creates a new header.
func (client *TrustSignalClient) CreateHeader(data CreateHeaderData, apiKey ...string) (*HeaderResponse, error) {
	url := fmt.Sprintf("%s/accounts/senders", client.BaseURL)
	body, err := json.Marshal(data)
	if err != nil {
		client.Logger.Error("Failed to marshal CreateHeaderData", "error", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		client.Logger.Error("Failed to create new request for Create Header", "error", err)
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
		client.Logger.Error("Failed to send request for Create Header", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Create Header API", "status_code", resp.StatusCode)
	}

	var response HeaderResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode HeaderResponse", "error", err)
		return nil, err
	}

	return &response, nil
}
