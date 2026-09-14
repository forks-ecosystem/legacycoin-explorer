package explorer

// exp.go — стартовый блок интерфейса: публичная страница статуса узла.
// Без авторизации: эксплорер — публичный ресурс.

import (
	"html/template"
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

var statusTemplate = template.Must(template.New("status").Funcs(template.FuncMap{
	"formatHashRate": func(v int64) string { return formatHashRate(float64(v)) },
	"truncate":       truncate,
}).Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#141414;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;background:var(--black);}
body{display:flex;flex-direction:column;color:var(--text);font-family:'Arial',sans-serif;font-size:15px;}
a{color:var(--gold);text-decoration:none;}
.tk{flex-shrink:0;padding:12px 22px;display:flex;align-items:center;gap:12px;font-weight:700;font-size:15px;}
.tk b{color:var(--gold);font-family:var(--mono);font-size:14px;}
.bge{display:inline-block;padding:2px 7px;font-size:11px;font-weight:600;}
.bg-g{background:rgba(34,197,94,.12);color:var(--green);border:1px solid rgba(34,197,94,.25);}
.bg-d{background:rgba(239,68,68,.12);color:var(--red);border:1px solid rgba(239,68,68,.25);}
.sc{flex:1;min-height:0;overflow:auto;}
.in{max-width:1000px;margin:0 auto;padding:22px;}
.dc{background:var(--panel);border:1px solid var(--border);padding:20px;}
.dc h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;padding-bottom:9px;border-bottom:1px solid var(--border);}
.dr{display:flex;justify-content:space-between;gap:14px;padding:6px 0;border-bottom:1px solid rgba(255,255,255,.04);}
.dr:last-child{border-bottom:none;}
.dk{font-size:12px;color:var(--muted);flex-shrink:0;}
.dv{font-size:12px;text-align:right;word-break:break-all;font-family:var(--mono);}
.dv.gold{color:var(--gold);font-weight:700;}
.foot{flex-shrink:0;display:flex;justify-content:space-between;gap:12px;padding:12px 22px;font-size:13px;}
</style>
<div class="tk"><span>Node status</span>{{if .Status.NodeRunning}}<span class="bge bg-g">online</span>{{else}}<span class="bge bg-d">offline</span>{{end}}</div>
<div class="sc"><div class="in">
<div class="dc"><h3>Status</h3>
<div class="dr"><span class="dk">Running</span><span class="dv">{{.Status.NodeRunning}}</span></div>
<div class="dr"><span class="dk">Height</span><span class="dv gold">{{.Status.Height}}</span></div>
<div class="dr"><span class="dk">Best Block Hash</span><span class="dv">{{truncate .Status.BestHash 60}}</span></div>
<div class="dr"><span class="dk">Difficulty</span><span class="dv">{{.Status.Difficulty}}</span></div>
<div class="dr"><span class="dk">Connections</span><span class="dv">{{.Status.Connections}}</span></div>
<div class="dr"><span class="dk">Network</span><span class="dv">{{.Status.Network}}</span></div>
<div class="dr"><span class="dk">Coin</span><span class="dv">{{.Status.Coin}}</span></div>
<div class="dr"><span class="dk">Mempool</span><span class="dv">{{.Status.Mempool}}</span></div>
<div class="dr"><span class="dk">Hash Rate</span><span class="dv">{{formatHashRate .Status.HashRate}}</span></div>
</div>
</div></div>
<div class="foot"><a href="/" target="_top">← Home</a></div>
`))

// handleStatusPage — публичная первая страница интерфейса.
func (s *Server) handleStatusPage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		if err := statusTemplate.Execute(w, map[string]interface{}{
			"Status": s.rpc.NodeStatus(),
		}); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

// handleAPIStatus — JSON для страницы статуса.
func (s *Server) handleAPIStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonOK(w, s.rpc.NodeStatus())
	}
}
