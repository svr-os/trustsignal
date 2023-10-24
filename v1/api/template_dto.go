package api

// CreateTemplateData contains the information required to create a new template.
type CreateTemplateData struct {
	Name        string   `json:"name"`
	Content     string   `json:"content"`
	Headers     []string `json:"headers"`
	DLTEntityID string   `json:"dlt_entity_id"`
}

// TemplateResponse contains the response fields for creating or getting a template.
type TemplateResponse struct {
	// TODO: verify and add for template creation
	// Name        string   `json:"name,omitempty"`
	// Content     string   `json:"content,omitempty"`
	// Headers     []string `json:"headers,omitempty"`
	// DLTEntityID string   `json:"dlt_entity_id,omitempty"`
}
