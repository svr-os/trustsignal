package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// CreateSubAccount creates a new sub-account.
func (client *TrustSignalClient) CreateSubAccount(data SubAccountData, apiKey ...string) (*SubAccountResponse, error) {
	url := fmt.Sprintf("%s/auth/register/subaccount", client.BaseURL)
	body, err := json.Marshal(data)
	if err != nil {
		client.Logger.Error("Failed to marshal SubAccountData", "error", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		client.Logger.Error("Failed to create new request for Create SubAccount", "error", err)
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
		client.Logger.Error("Failed to send request for Create SubAccount", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Create SubAccount API", "status_code", resp.StatusCode)
	}

	var response SubAccountResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode SubAccountResponse", "error", err)
		return nil, err
	}

	return &response, nil
}
