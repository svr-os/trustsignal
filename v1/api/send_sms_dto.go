package api

type SendSMSData struct {
	SenderID string   `json:"sender_id"`
	To       []string `json:"to"`
	Route    string   `json:"route"`
	Message  string   `json:"message"`
	Pr1      string   `json:"pr1,omitempty"`
}

type SendSMSResponse struct {
	// TODO: Modify accordingly based on actual response from TrustSignal
	// example:
	// Status  string `json:"status,omitempty"`
	// Message string `json:"message,omitempty"`
}
