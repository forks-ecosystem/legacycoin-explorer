package explorer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

// RPCClient connects to a running legacycoind node via JSON-RPC.
type RPCClient struct {
	endpoint string
	user     string
	pass     string
	client   *http.Client
}

// NewRPCClient creates a new RPC client.
func NewRPCClient(host string, port int, user, pass string) *RPCClient {
	return &RPCClient{
		endpoint: fmt.Sprintf("http://%s:%d/", host, port),
		user:     user,
		pass:     pass,
		client:   &http.Client{Timeout: 30 * time.Second},
	}
}

type rpcRequest struct {
	Method string        `json:"method"`
	Params []interface{} `json:"params"`
	ID     int           `json:"id"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
	Error  *rpcError       `json:"error"`
	ID     int             `json:"id"`
}

type rpcError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *rpcError) Error() string { return fmt.Sprintf("RPC error %d: %s", e.Code, e.Message) }

func (c *RPCClient) call(method string, params ...interface{}) (json.RawMessage, error) {
	if params == nil {
		params = []interface{}{}
	}
	body, _ := json.Marshal(rpcRequest{Method: method, Params: params, ID: 1})
	req, _ := http.NewRequest("POST", c.endpoint, bytes.NewReader(body))
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("RPC connection failed: %w", err)
	}
	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)
	var rr rpcResponse
	if err := json.Unmarshal(data, &rr); err != nil {
		return nil, fmt.Errorf("RPC decode error: %w", err)
	}
	if rr.Error != nil {
		return nil, rr.Error
	}
	return rr.Result, nil
}

// ── High-level methods ────────────────────────────────────────────────────────

// NodeInfo holds data from getinfo.
type NodeInfo struct {
	Version         string  `json:"version"`
	Blocks          int64   `json:"blocks"`
	Connections     int     `json:"connections"`
	Difficulty      float64 `json:"difficulty"`
	Errors          string  `json:"errors"`
	Network         string  `json:"network"`
	Coin            string  `json:"coin"`
	Ticker          string  `json:"ticker"`
	CoreVersion     string  `json:"core_version"`
	BestBlockHash   string  `json:"bestblockhash"`
}

func (c *RPCClient) GetInfo() (*NodeInfo, error) {
	raw, err := c.call("getinfo")
	if err != nil {
		return nil, err
	}
	var info NodeInfo
	return &info, json.Unmarshal(raw, &info)
}

// MiningInfo holds data from getmininginfo.
type MiningInfo struct {
	Blocks       int64   `json:"blocks"`
	Difficulty   float64 `json:"difficulty"`
	Generate     bool    `json:"generate"`
	HashesPerSec int64   `json:"hashespersec"`
	PooledTx     int     `json:"pooledtx"`
	LiveHashrate int64   `json:"live_hashrate"`
}

func (c *RPCClient) GetMiningInfo() (*MiningInfo, error) {
	raw, err := c.call("getmininginfo")
	if err != nil {
		return nil, err
	}
	var info MiningInfo
	return &info, json.Unmarshal(raw, &info)
}

func (c *RPCClient) GetBlockCount() (int64, error) {
	raw, err := c.call("getblockcount")
	if err != nil {
		return 0, err
	}
	var n float64
	if err := json.Unmarshal(raw, &n); err != nil {
		return 0, err
	}
	return int64(n), nil
}

func (c *RPCClient) GetBestBlockHash() (string, error) {
	raw, err := c.call("getbestblockhash")
	if err != nil {
		return "", err
	}
	var h string
	return h, json.Unmarshal(raw, &h)
}

func (c *RPCClient) GetBlockHash(height int64) (string, error) {
	raw, err := c.call("getblockhash", height)
	if err != nil {
		return "", err
	}
	var h string
	return h, json.Unmarshal(raw, &h)
}

// Block holds block data from getblock.
type Block struct {
	Hash              string   `json:"hash"`
	Height            int64    `json:"height"`
	Version           uint32   `json:"version"`
	PreviousBlockHash string   `json:"previousblockhash"`
	MerkleRoot        string   `json:"merkleroot"`
	Time              uint32   `json:"time"`
	Bits              string   `json:"bits"`
	Nonce             uint32   `json:"nonce"`
	Tx                []string `json:"tx"`
	Size              int      `json:"size"`
	Confirmations     int64    `json:"confirmations"`
}

func (c *RPCClient) GetBlock(hash string) (*Block, error) {
	raw, err := c.call("getblock", hash)
	if err != nil {
		return nil, err
	}
	var b Block
	return &b, json.Unmarshal(raw, &b)
}

func (c *RPCClient) GetBlockAtHeight(height int64) (*Block, error) {
	hash, err := c.GetBlockHash(height)
	if err != nil {
		return nil, err
	}
	b, err := c.GetBlock(hash)
	if err != nil {
		return nil, err
	}
	b.Height = height
	return b, nil
}

// GetRecentBlocks returns the last n blocks, newest first.
func (c *RPCClient) GetRecentBlocks(n int) ([]*Block, error) {
	tip, err := c.GetBlockCount()
	if err != nil {
		return nil, err
	}
	blocks := make([]*Block, 0, n)
	for h := tip; h >= 0 && len(blocks) < n; h-- {
		b, err := c.GetBlockAtHeight(h)
		if err != nil {
			break
		}
		b.Confirmations = tip - h + 1
		blocks = append(blocks, b)
	}
	return blocks, nil
}

// Ping checks if the node is reachable.
func (c *RPCClient) Ping() bool {
	_, err := c.GetBlockCount()
	return err == nil
}
