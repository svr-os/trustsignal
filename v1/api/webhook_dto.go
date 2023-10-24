package api

import (
	"time"
)

// WebhookData contains the information received in the webhook payload from TrustSignal.
type WebhookData struct {
	TransactionID string    `json:"transaction_id"`
	MID           string    `json:"mid"`
	To            string    `json:"to"`
	Route         string    `json:"route"`
	SubmitTime    time.Time `json:"st"`
	DeliveryTime  time.Time `json:"dlrt"`
	Status        string    `json:"status"`
	Error         string    `json:"error"`
	Pr1           string    `json:"pr1"`
	Pr2           string    `json:"pr2"`
}
