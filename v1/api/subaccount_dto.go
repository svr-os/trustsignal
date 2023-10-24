package api

// SubAccountData contains the information required to create a new sub-account.
type SubAccountData struct {
	Email    string `json:"email"`
	Phone    int64  `json:"phone"`
	Name     string `json:"name"`
	Company  string `json:"company"`
	Password string `json:"password"`
}

// SubAccountResponse contains the response fields after creating a sub-account.
type SubAccountResponse struct {
	// TODO: Verify response structure.
	Email   string `json:"email,omitempty"`
	Phone   int64  `json:"phone,omitempty"`
	Name    string `json:"name,omitempty"`
	Company string `json:"company,omitempty"`
}
