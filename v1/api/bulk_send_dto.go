package api

// BulkSMSData represents the payload for sending a bulk SMS.
type BulkSMSData struct {
	SenderID  string           `json:"sender_id"`
	Receivers []ReceiverDetail `json:"receivers"`
	Route     string           `json:"route"`
	Pr1       string           `json:"pr1"`
}

// ReceiverDetail contains details for individual SMS recipients.
type ReceiverDetail struct {
	To      int64  `json:"to"`
	Message string `json:"message"`
}

// BulkSendResponse represents the response from the bulk SMS send API.
type BulkSendResponse struct {
	// TODO: add proper response fields
	// For example:
	// Status    string `json:"status"`
	// MessageID string `json:"message_id"`
}
