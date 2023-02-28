package tangocrypto_go

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceAddress      = "address"
	resourceAddresses    = "addresses"
	resourceTotal        = "total"
	resourceTransactions = "transactions"
	resourceUTXOs        = "utxos"
)

type AddressSummary struct {
	Network           string `json:"network"`
	Address           string `json:"address"`
	StakeAddress      string `json:"stake_address"`
	Balance           int    `json:"balance"`
	TransactionsCount int    `json:"transactions_count"`
}

type AddrUTXOs struct {
	Data   []Data      `json:"data"`
	Cursor interface{} `json:"cursor"`
}

type Assets struct {
	PolicyID    string `json:"policy_id"`
	AssetName   string `json:"asset_name"`
	Fingerprint string `json:"fingerprint"`
	Quantity    int    `json:"quantity"`
}

type List struct {
	Bytes string `json:"bytes"`
}

type Fields struct {
	List  []List `json:"list,omitempty"`
	Bytes string `json:"bytes,omitempty"`
}

type Value struct {
	Fields      []Fields `json:"fields"`
	Constructor int      `json:"constructor"`
}

type InlineDatum struct {
	Hash     string `json:"hash"`
	Value    Value  `json:"value"`
	ValueRaw string `json:"value_raw"`
}

type Datum struct {
	Hash string `json:"hash"`
}

type Script struct {
	Type           string `json:"type"`
	Hash           string `json:"hash"`
	Code           string `json:"code"`
	SerialisedSize int    `json:"serialised_size"`
	Datum          Datum  `json:"datum"`
}

type Data struct {
	Address     string      `json:"address"`
	Hash        string      `json:"hash"`
	Index       int         `json:"index"`
	Value       int         `json:"value"`
	HasScript   bool        `json:"has_script"`
	Assets      []Assets    `json:"assets"`
	InlineDatum InlineDatum `json:"inline_datum"`
	Script      Script      `json:"script"`
}

func (c *apiClient) AddressSummary(ctx context.Context, address string) (addressSum AddressSummary, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s", c.server, c.appID, resourceAddress, address))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return
	}

	resp, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&addressSum); err != nil {
		return
	}

	return addressSum, nil
}

// TODO: add pagination support
// add query params to the request
func (c *apiClient) AddressUTXOs(ctx context.Context, address string) (utxos AddrUTXOs, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s/%s", c.server, c.appID, resourceAddresses, address, resourceUTXOs))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestURL.String(), nil)
	if err != nil {
		return
	}

	resp, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&utxos); err != nil {
		return
	}

	return utxos, nil
}
