package explorer

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const maxRetries = 5
const retryDelay = 500 * time.Millisecond

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
		client:   &http.Client{Timeout: 60 * time.Second},
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

	var lastErr error
	for attempt := 0; attempt < maxRetries; attempt++ {
		req, _ := http.NewRequest("POST", c.endpoint, bytes.NewReader(body))
		req.SetBasicAuth(c.user, c.pass)
		req.Header.Set("Content-Type", "application/json")

		resp, err := c.client.Do(req)
		if err != nil {
			lastErr = fmt.Errorf("RPC connection failed: %w", err)
			time.Sleep(retryDelay)
			continue
		}
		data, _ := io.ReadAll(resp.Body)
		resp.Body.Close()

		if resp.StatusCode == 429 {
			lastErr = fmt.Errorf("rate limited")
			time.Sleep(retryDelay * time.Duration(attempt+1))
			continue
		}
		if resp.StatusCode != 200 {
			lastErr = fmt.Errorf("RPC HTTP %d: %s", resp.StatusCode, string(data))
			time.Sleep(retryDelay)
			continue
		}

		var rr rpcResponse
		if err := json.Unmarshal(data, &rr); err != nil {
			lastErr = fmt.Errorf("RPC decode error: %w", err)
			continue
		}
		if rr.Error != nil {
			return nil, rr.Error
		}
		return rr.Result, nil
	}
	return nil, lastErr
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

// ── Transaction methods ──────────────────────────────────────────────────────

// TxInput holds a transaction input.
type TxInput struct {
	Txid         string `json:"txid"`
	Vout         int    `json:"vout"`
	ScriptSig    string `json:"scriptSig"`
	Sequence     uint32 `json:"sequence"`
}

// TxOutput holds a transaction output.
type TxOutput struct {
	Value        float64 `json:"value"`
	N            int     `json:"n"`
	ScriptPubKey struct {
		Asm       string   `json:"asm"`
		Hex       string   `json:"hex"`
		Type      string   `json:"type"`
		Addresses []string `json:"addresses"`
	} `json:"scriptPubKey"`
}

// RawTransaction holds raw transaction data.
type RawTransaction struct {
	Txid     string      `json:"txid"`
	Hash     string      `json:"hash"`
	Version  int         `json:"version"`
	Size     int         `json:"size"`
	Vin      []TxInput   `json:"vin"`
	Vout     []TxOutput  `json:"vout"`
	Blockhash string     `json:"blockhash"`
	Height    int64      `json:"blockheight"`
	Confirmations int64  `json:"confirmations"`
	Time     uint32      `json:"time"`
	Blocktime uint32     `json:"blocktime"`
}

func (c *RPCClient) GetRawTransaction(txid string) (*RawTransaction, error) {
	raw, err := c.call("getrawtransaction", txid, true)
	if err != nil {
		return nil, err
	}
	var tx RawTransaction
	return &tx, json.Unmarshal(raw, &tx)
}

// ── Address methods ─────────────────────────────────────────────────────────

// AddressInfo holds address validation data.
type AddressInfo struct {
	Address    string `json:"address"`
	IsValid    bool   `json:"isvalid"`
	IsMine     bool   `json:"ismine"`
	IsScript   bool   `json:"isscript"`
	IsWatchOnly bool  `json:"iswatchonly"`
	PubKeyHash string `json:"pubkey_hash_hex"`
}

func (c *RPCClient) ValidateAddress(address string) (*AddressInfo, error) {
	raw, err := c.call("validateaddress", address)
	if err != nil {
		return nil, err
	}
	var info AddressInfo
	return &info, json.Unmarshal(raw, &info)
}

// ── Block hex parsing ────────────────────────────────────────────────────────

// GetBlockHex returns the raw block hex string.
func (c *RPCClient) GetBlockHex(hash string) (string, error) {
	raw, err := c.call("getblock", hash)
	if err != nil {
		return "", err
	}
	var b struct {
		Hex string `json:"hex"`
	}
	if err := json.Unmarshal(raw, &b); err != nil {
		return "", err
	}
	return b.Hex, nil
}

// ExtractTxHexes parses a raw block hex and returns the raw hex of each transaction.
func ExtractTxHexes(blockHex string) ([]string, error) {
	data, err := hex.DecodeString(blockHex)
	if err != nil {
		return nil, fmt.Errorf("invalid hex: %w", err)
	}

	r := bytes.NewReader(data)

	// Block header: 4+32+32+4+4+4 = 80 bytes
	header := make([]byte, 80)
	if _, err := io.ReadFull(r, header); err != nil {
		return nil, fmt.Errorf("block header too short")
	}

	// Transaction count (varint)
	txCount, err := readVarInt(r)
	if err != nil {
		return nil, err
	}

	var txHexes []string
	for i := int64(0); i < txCount; i++ {
		txData, err := extractOneTx(r)
		if err != nil {
			return nil, fmt.Errorf("tx %d: %w", i, err)
		}
		txHexes = append(txHexes, hex.EncodeToString(txData))
	}

	return txHexes, nil
}

// extractOneTx reads one serialized transaction from r and returns its bytes.
func extractOneTx(r *bytes.Reader) ([]byte, error) {
	var buf bytes.Buffer
	start := make([]byte, 4)
	if _, err := io.ReadFull(r, start); err != nil {
		return nil, err
	}
	buf.Write(start)

	// Read inputs
	inCount, err := readVarInt(r)
	if err != nil {
		return nil, err
	}
	writeVarInt(&buf, inCount)

	for i := int64(0); i < inCount; i++ {
		// prev tx hash (32) + prev tx index (4)
		if _, err := io.CopyN(&buf, r, 36); err != nil {
			return nil, err
		}
		// scriptSig length
		scriptLen, err := readVarInt(r)
		if err != nil {
			return nil, err
		}
		writeVarInt(&buf, scriptLen)
		// scriptSig
		if _, err := io.CopyN(&buf, r, scriptLen); err != nil {
			return nil, err
		}
		// sequence (4)
		if _, err := io.CopyN(&buf, r, 4); err != nil {
			return nil, err
		}
	}

	// Read outputs
	outCount, err := readVarInt(r)
	if err != nil {
		return nil, err
	}
	writeVarInt(&buf, outCount)

	for i := int64(0); i < outCount; i++ {
		// value (8)
		if _, err := io.CopyN(&buf, r, 8); err != nil {
			return nil, err
		}
		// scriptPubKey length
		scriptLen, err := readVarInt(r)
		if err != nil {
			return nil, err
		}
		writeVarInt(&buf, scriptLen)
		// scriptPubKey
		if _, err := io.CopyN(&buf, r, scriptLen); err != nil {
			return nil, err
		}
	}

	// locktime (4)
	if _, err := io.CopyN(&buf, r, 4); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func readVarInt(r *bytes.Reader) (int64, error) {
	b, err := r.ReadByte()
	if err != nil {
		return 0, err
	}
	switch b {
	case 0xfd:
		var v uint16
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return int64(v), nil
	case 0xfe:
		var v uint32
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return int64(v), nil
	case 0xff:
		var v uint64
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return 0, err
		}
		return int64(v), nil
	default:
		return int64(b), nil
	}
}

func writeVarInt(w *bytes.Buffer, v int64) {
	if v < 0xfd {
		w.WriteByte(byte(v))
	} else if v <= 0xffff {
		w.WriteByte(0xfd)
		binary.Write(w, binary.LittleEndian, uint16(v))
	} else if v <= 0xffffffff {
		w.WriteByte(0xfe)
		binary.Write(w, binary.LittleEndian, uint32(v))
	} else {
		w.WriteByte(0xff)
		binary.Write(w, binary.LittleEndian, uint64(v))
	}
}

// ── Block search for address/tx ─────────────────────────────────────────────

// SearchTxInBlock searches for a txid in a specific block.
func (c *RPCClient) SearchTxInBlock(block *Block, txid string) bool {
	for _, tx := range block.Tx {
		if tx == txid {
			return true
		}
	}
	return false
}

// FindTransaction finds a transaction by txid.
// Uses getrawtransaction directly, falling back to block scan.
func (c *RPCClient) FindTransaction(txid string) (*RawTransaction, *Block, error) {
	tx, err := c.GetRawTransaction(txid)
	if err == nil {
		var block *Block
		if tx.Blockhash != "" {
			block, _ = c.GetBlock(tx.Blockhash)
			if block != nil {
				block.Height = tx.Height
				block.Confirmations = tx.Confirmations
			}
		}
		return tx, block, nil
	}
	tip, err := c.GetBlockCount()
	if err != nil {
		return nil, nil, err
	}
	maxScan := tip
	if maxScan > 20 {
		maxScan = 20
	}
	for h := tip; h > tip-maxScan; h-- {
		hash, err := c.GetBlockHash(h)
		if err != nil {
			continue
		}
		block, err := c.GetBlock(hash)
		if err != nil {
			continue
		}
		block.Height = h
		block.Confirmations = tip - h + 1
		if c.SearchTxInBlock(block, txid) {
			tx, blockErr := c.decodeTxFromBlock(hash, txid)
			if blockErr != nil {
				return nil, block, blockErr
			}
			tx.Blockhash = hash
			tx.Height = h
			tx.Confirmations = tip - h + 1
			return tx, block, nil
		}
	}
	return nil, nil, fmt.Errorf("transaction not found")
}

func (c *RPCClient) decodeTxFromBlock(blockHash, targetTxid string) (*RawTransaction, error) {
	blockHex, err := c.GetBlockHex(blockHash)
	if err != nil {
		return nil, err
	}
	txHexes, err := ExtractTxHexes(blockHex)
	if err != nil {
		return nil, err
	}
	for _, txHex := range txHexes {
		raw, err := c.call("decoderawtransaction", txHex)
		if err != nil {
			continue
		}
		var tx RawTransaction
		if err := json.Unmarshal(raw, &tx); err != nil {
			continue
		}
		if tx.Txid == targetTxid {
			return &tx, nil
		}
	}
	return nil, fmt.Errorf("tx not found in block hex")
}

// FindAddressTxs scans recent blocks to find transactions involving an address.
// Uses hex-based approach: one getblock + N decoderawtransaction calls per block.
func (c *RPCClient) FindAddressTxs(address string, maxBlocks int) ([]*RawTransaction, error) {
	tip, err := c.GetBlockCount()
	if err != nil {
		return nil, err
	}

	var txs []*RawTransaction
	maxScan := tip
	if maxScan > 10 {
		maxScan = 10
	}

	for h := tip; h > tip-maxScan && len(txs) < 50; h-- {
		hash, err := c.GetBlockHash(h)
		if err != nil {
			continue
		}
		blockHex, err := c.GetBlockHex(hash)
		if err != nil {
			continue
		}
		txHexes, err := ExtractTxHexes(blockHex)
		if err != nil {
			continue
		}

		for _, txHex := range txHexes {
			raw, err := c.call("decoderawtransaction", txHex)
			if err != nil {
				continue
			}
			var tx RawTransaction
			if err := json.Unmarshal(raw, &tx); err != nil {
				continue
			}

			for _, vout := range tx.Vout {
				for _, addr := range vout.ScriptPubKey.Addresses {
					if addr == address {
						tx.Blockhash = hash
						tx.Height = h
						tx.Confirmations = tip - h + 1
						txs = append(txs, &tx)
						goto nextBlock
					}
				}
			}
		}
	nextBlock:
	}
	return txs, nil
}
