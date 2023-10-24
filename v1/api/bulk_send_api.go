package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *TrustSignalClient) BulkSendSMS(data BulkSMSData, apiKey ...string) (*BulkSendResponse, error) {
	url := fmt.Sprintf("%s/sms/country-code-bulk", client.BaseURL)
	body, err := json.Marshal(data)
	if err != nil {
		client.Logger.Error("Failed to marshal BulkSMSData", "error", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		client.Logger.Error("Failed to create new request for Bulk SMS", "error", err)
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
		client.Logger.Error("Failed to send request for Bulk SMS", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Bulk SMS API", "status_code", resp.StatusCode)
	}

	var response BulkSendResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode BulkSendResponse", "error", err)
		return nil, err
	}

	client.Logger.Info("Bulk SMS sent successfully", "receiver_count", len(data.Receivers))
	return &response, nil
}
