package explorer

// exp.go — стартовый блок интерфейса: публичная страница статуса узла.
// Без авторизации: эксплорер — публичный ресурс.

import (
	"net/http"
)

// NodeStatus — живой статус узла из RPC (адаптация GetWalletInfo, без CLI).
type NodeStatus struct {
	NodeRunning bool    `json:"nodeRunning"`
	Height      int64   `json:"height"`
	BestHash    string  `json:"bestHash"`
	Difficulty  float64 `json:"difficulty"`
	Connections int     `json:"connections"`
	Network     string  `json:"network"`
	Coin        string  `json:"coin"`
	Mempool     int     `json:"mempool"`
	HashRate    int64   `json:"hashRate"`
}

func (c *RPCClient) NodeStatus() NodeStatus {
	st := NodeStatus{}
	if !c.Ping() {
		return st
	}
	st.NodeRunning = true
	if info, err := c.GetInfo(); err == nil {
		st.Height = info.Blocks
		st.BestHash = info.BestBlockHash
		st.Difficulty = info.Difficulty
		st.Connections = info.Connections
		st.Network = info.Network
		st.Coin = info.Coin
	}
	if m, err := c.GetMiningInfo(); err == nil {
		st.Mempool = m.PooledTx
		st.HashRate = m.HashesPerSec
	}
	return st
}

// handleStatusPage — публичная первая страница интерфейса.
func (s *Server) handleStatusPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.render(w, "status", nil)
	}
}

// handleAPIStatus — JSON для страницы статуса.
func (s *Server) handleAPIStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonOK(w, s.rpc.NodeStatus())
	}
}
