package api

import "github.com/svr-os/trustsignal/logger"

// TrustSignalService represents the interface with methods available to interact with trust signal apis
type TrustSignalService interface {
	BulkSendSMS(data BulkSMSData, apiKey ...string) (*BulkSendResponse, error)
	SendSMS(data SendSMSData, apiKey ...string) (*SendSMSResponse, error)
	GetHeader(apiKey ...string) (*HeaderResponse, error)
	CreateHeader(data CreateHeaderData, apiKey ...string) (*HeaderResponse, error)
	GetTemplate(apiKey ...string) (*TemplateResponse, error)
	CreateTemplate(data CreateTemplateData, apiKey ...string) (*TemplateResponse, error)
	CreateSubAccount(data SubAccountData, apiKey ...string) (*SubAccountResponse, error)
}

type TrustSignalClient struct {
	BaseURL string
	APIKey  string
	Logger  logger.Logger
}

type ClientOptions func(*TrustSignalClient)

func NewTrustSignalClient(apiKey string, options ...ClientOptions) TrustSignalService {
	client := &TrustSignalClient{
		BaseURL: "https://api.trustsignal.io/v1",
		APIKey:  apiKey,
		Logger:  &logger.DefaultLogger{},
	}

	for _, opt := range options {
		opt(client)
	}

	return client
}

func WithLogger(l logger.Logger) ClientOptions {
	return func(c *TrustSignalClient) {
		c.Logger = l
	}
}
