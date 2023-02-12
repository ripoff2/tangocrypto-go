package tangocrypto_go

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceSubmit = "submit"
)

type TransactionContent struct {
	Hash                 string      `json:"hash"`
	BlockID              string      `json:"block_id"`
	BlockIndex           int         `json:"block_index"`
	OutSum               string      `json:"out_sum"`
	Fee                  string      `json:"fee"`
	Deposit              string      `json:"deposit"`
	Size                 int         `json:"size"`
	InvalidBefore        interface{} `json:"invalid_before"`
	InvalidHereafter     string      `json:"invalid_hereafter"`
	ValidContract        bool        `json:"valid_contract"`
	ScriptSize           int         `json:"script_size"`
	UtxoCount            string      `json:"utxo_count"`
	WithdrawalCount      string      `json:"withdrawal_count"`
	DelegationCount      string      `json:"delegation_count"`
	StakeCertCount       string      `json:"stake_cert_count"`
	PoolUpdate           bool        `json:"pool_update"`
	PoolRetire           bool        `json:"pool_retire"`
	AssetMintOrBurnCount string      `json:"asset_mint_or_burn_count"`
	Block                Block       `json:"block"`
	Assets               []Assets    `json:"assets"`
}

type Block struct {
	Hash    string `json:"hash"`
	EpochNo int    `json:"epoch_no"`
	BlockNo int    `json:"block_no"`
	SlotNo  int    `json:"slot_no"`
}

func (c *apiClient) Transaction(ctx context.Context, hash string) (content TransactionContent, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s", c.server, c.appID, resourceTransactions, hash))
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

	if err = json.NewDecoder(resp.Body).Decode(&content); err != nil {
		return
	}

	return content, nil
}

func (c *apiClient) TransactionSubmit(ctx context.Context, cbor []byte) (hash string, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s", c.server, c.appID, resourceTransactions, resourceSubmit))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestURL.String(), bytes.NewReader(cbor))
	if err != nil {
		return
	}

	resp, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&hash); err != nil {
		return
	}

	return hash, nil
}
