package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (client *TrustSignalClient) SendSMS(data SendSMSData, apiKey ...string) (*SendSMSResponse, error) {
	url := fmt.Sprintf("%s/sms/countrycode", client.BaseURL)
	body, err := json.Marshal(data)
	if err != nil {
		client.Logger.Error("Failed to marshal SendSMSData", "error", err)
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		client.Logger.Error("Failed to create new request for Send SMS", "error", err)
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
		client.Logger.Error("Failed to send request for Send SMS", "error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		client.Logger.Warn("Received non-200 status code from Send SMS API", "status_code", resp.StatusCode)
	}

	var response SendSMSResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		client.Logger.Error("Failed to decode SendSMSResponse", "error", err)
		return nil, err
	}

	client.Logger.Info("SMS sent successfully", "receiver", data.To[0]) // Assuming sending to one recipient for now

	return &response, nil
}
