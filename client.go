package tangocrypto_go

import (
	"context"
	"net/http"
)

type apiClient struct {
	server string
	appID  string
	apiKey string
	client *http.Client
}

// HttpRequestDoer defines methods for a http client.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// APIClientOptions contains optios used to initialize an API Client using
// NewAPIClient
type APIClientOptions struct {
	AppID  string
	ApiKey string

	// Server url to use
	Server string
}

// NewAPICLient creates a client from APIClientOptions. If no options are provided,
//  client with default configurations is returned.
func NewAPIClient(options APIClientOptions) APIClient {
	if options.Server == "" {
		options.Server = CardanoMainNet
	}

	c := &http.Client{}

	client := &apiClient{
		server: options.Server,
		client: c,
		appID:  options.AppID,
		apiKey: options.ApiKey,
	}

	return client
}

// APIClient defines methods implemented by the api client.
type APIClient interface {
	AddressSummary(ctx context.Context, address string) (AddressSummary, error)
	AddressUTXOs(ctx context.Context, address string) (AddrUTXOs, error)
	TransactionSubmit(ctx context.Context, cbor []byte) (string, error)
	ProtocolParameters(ctx context.Context, epochNumber string) (EpochParameters, error)
	CurrentEpoch(ctx context.Context) (CurrentEpoch, error)
	LatestBlock(ctx context.Context) (LatestBlock, error)
}
