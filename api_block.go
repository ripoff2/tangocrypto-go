package tangocrypto_go

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	blocksResource = "blocks"
	latestResource = "latest"
)

type LatestBlock struct {
	Hash          string `json:"hash"`
	EpochNo       int    `json:"epoch_no"`
	SlotNo        int    `json:"slot_no"`
	EpochSlotNo   int    `json:"epoch_slot_no"`
	BlockNo       int    `json:"block_no"`
	PreviousBlock int    `json:"previous_block"`
	NextBlock     int    `json:"next_block"`
	SlotLeader    string `json:"slot_leader"`
	OutSum        int    `json:"out_sum"`
	Fees          int    `json:"fees"`
	Confirmations int    `json:"confirmations"`
	Size          int    `json:"size"`
	Time          string `json:"time"`
	TxCount       int    `json:"tx_count"`
	VrfKey        string `json:"vrf_key"`
	OpCert        string `json:"op_cert"`
}

func (c *apiClient) LatestBlock(ctx context.Context) (b LatestBlock, err error) {
	requestURL, err := url.Parse(fmt.Sprintf("%s/%s/v1/%s/%s", c.server, c.appID, blocksResource, latestResource))
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

	if err = json.NewDecoder(resp.Body).Decode(&b); err != nil {
		return
	}

	return b, nil
}