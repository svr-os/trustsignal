package api

type CreateHeaderData struct {
	Code        string `json:"code"`
	DLTEntityID string `json:"dlt_entity_id"`
	Company     string `json:"company"`
	Brand       string `json:"brand"`
}

type HeaderResponse struct {
	// TODO: Verify and add response structure, verify difference in response for get and create
	// Code        string `json:"code,omitempty"`
	// DLTEntityID string `json:"dlt_entity_id,omitempty"`
	// Company     string `json:"company,omitempty"`
	// Brand       string `json:"brand,omitempty"`
}
