// exp_block.go — сервисная карточка блока в фрейме fC (Enter → /exp-block/{kod}).
// Трехпанельный вид: Header / Summary / Transactions. Без чужого навбара и редиректов.
package explorer

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var expBlockTemplate = template.Must(template.New("expBlock").Funcs(template.FuncMap{
	"formatTime":  formatTime,
	"formatLBTC":  formatLBTC,
	"truncate":    truncate,
	"blockReward": blockRewardForHeight,
	"add":         func(a, b int64) int64 { return a + b },
	"sub":         func(a, b int64) int64 { return a - b },
}).Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#141414;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;background:var(--black);}
body{display:flex;flex-direction:column;color:var(--text);font-family:'Arial',sans-serif;font-size:15px;line-height:1.6;}
a{color:var(--gold);text-decoration:none;}
.tk{flex-shrink:0;padding:12px 22px;display:flex;align-items:center;gap:12px;font-weight:700;font-size:15px;}
.tk b{color:var(--gold);font-family:var(--mono);font-size:14px;}
.bge{display:inline-block;padding:2px 7px;font-size:11px;font-weight:600;}
.bg-g{background:rgba(34,197,94,.12);color:var(--green);border:1px solid rgba(34,197,94,.25);}
.sc{flex:1;min-height:0;overflow:auto;}
.in{max-width:1000px;margin:0 auto;padding:22px;}
.offline{background:rgba(239,68,68,.1);border:1px solid rgba(239,68,68,.3);color:var(--red);padding:13px 17px;font-size:14px;border-radius:8px;}
.offline a{display:inline-block;margin-top:12px;}
.dg{display:grid;grid-template-columns:1fr 1fr;gap:18px;margin-bottom:24px;}
@media(max-width:700px){.dg{grid-template-columns:1fr;}}
.dc{background:var(--panel);border:1px solid var(--border);padding:20px;}
.dc h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;padding-bottom:9px;border-bottom:1px solid var(--border);}
.dr{display:flex;justify-content:space-between;gap:14px;padding:6px 0;border-bottom:1px solid rgba(255,255,255,.04);}
.dr:last-child{border-bottom:none;}
.dk{font-size:12px;color:var(--muted);flex-shrink:0;}
.dv{font-size:12px;text-align:right;word-break:break-all;font-family:var(--mono);}
.dv.gold{color:var(--gold);font-weight:700;}
.pt{font-size:16px;font-weight:700;margin-bottom:10px;}.pt span{color:var(--gold);}
.tw{border:1px solid var(--border);}
table{width:100%;border-collapse:collapse;}
thead th{background:linear-gradient(to right,var(--panel2),#000);color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:10px 13px;text-align:left;border-bottom:1px solid var(--border);white-space:nowrap;}
tbody tr{border-bottom:1px solid var(--border);}tbody tr:last-child{border-bottom:none;}
tbody td{padding:10px 13px;font-size:13px;vertical-align:middle;}
.hash{font-family:var(--mono);font-size:12px;word-break:break-all;}
.bg-d{background:rgba(212,160,23,.12);color:var(--gold);border:1px solid rgba(212,160,23,.25);}
.foot{flex-shrink:0;display:flex;justify-content:space-between;gap:12px;padding:12px 22px;font-size:13px;}
</style>
<script>
function home(){ if(window.parent&&window.parent!==window&&typeof window.parent.lf==='function'){window.parent.lf('/exp-1');}else{location.href='/';} }
</script>
{{if .Error}}<div class="tk"><span>Block card</span></div>
<div class="sc"><div class="in"><div class="offline">⚠ {{.Error}}<a href="javascript:home()">← Home</a></div></div></div>
{{else}}<div class="tk"><span>Block</span> <b>#{{.Block.Height}}</b> <span class="bge bg-g">{{.Block.Confirmations}} confs</span></div>
<div class="sc"><div class="in">
<div class="dg">
<div class="dc"><h3>Header</h3>
<div class="dr"><span class="dk">Height</span><span class="dv gold">{{.Block.Height}}</span></div>
<div class="dr"><span class="dk">Hash</span><a class="dv" href="/block/{{.Block.Hash}}" target="_top">{{truncate .Block.Hash 50}}</a></div>
<div class="dr"><span class="dk">Previous</span><a class="dv" href="/block/{{.PrevHash}}" target="_top">{{truncate .PrevHash 50}}</a></div>
<div class="dr"><span class="dk">Merkle Root</span><span class="dv">{{truncate .Block.MerkleRoot 50}}</span></div>
<div class="dr"><span class="dk">Time</span><span class="dv">{{formatTime .Block.Time}}</span></div>
<div class="dr"><span class="dk">nBits</span><span class="dv">{{.Block.Bits}}</span></div>
<div class="dr"><span class="dk">Nonce</span><span class="dv">{{.Block.Nonce}}</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Block.Version}}</span></div>
</div>
<div class="dc"><h3>Summary</h3>
<div class="dr"><span class="dk">Transactions</span><span class="dv">{{len .Block.Tx}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Block.Size}} B</span></div>
<div class="dr"><span class="dk">Block Reward</span><span class="dv gold">{{formatLBTC (blockReward .Block.Height)}} LBTC</span></div>
<div class="dr"><span class="dk">Confirmations</span><span class="dv">{{.Block.Confirmations}}</span></div>
<div class="dr"><span class="dk">Algorithm</span><span class="dv">Yespower 1.0</span></div>
<div class="dr"><span class="dk">Difficulty Algo</span><span class="dv">DGW3</span></div>
</div>
</div>
<div class="pt">Transactions <span>({{len .Block.Tx}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>TXID</th><th>Type</th></tr></thead><tbody>
{{range $i, $tx := .Block.Tx}}
<tr>
<td style="color:var(--muted);">{{$i}}</td>
<td class="hash">{{$tx}}</td>
<td>{{if eq $i 0}}<span class="bge bg-d">Coinbase</span>{{else}}<span class="bge bg-g">Transfer</span>{{end}}</td>
</tr>
{{else}}
<tr><td colspan="3" style="text-align:center;color:var(--muted);padding:22px;">No transactions.</td></tr>
{{end}}
</tbody></table></div>
</div></div>
<div class="foot">
<div style="display:flex;gap:14px;">
<a href="/exp-block/{{sub .Block.Height 1}}">← Block {{sub .Block.Height 1}}</a>
<a href="/exp-block/{{add .Block.Height 1}}">Block {{add .Block.Height 1}} →</a>
</div>
</div>
{{end}}
</div>`))

func (s *Server) handleExpBlock(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/exp-block/"))

	var block *Block
	// Сначала тёплый кэш (узел под rate-limit), затем RPC.
	block = s.blockFromCache(id)
	if block == nil {
		var err error
		if h, perr := strconv.ParseInt(id, 10, 64); perr == nil {
			block, err = s.rpc.GetBlockAtHeight(h)
		} else {
			block, err = s.rpc.GetBlock(id)
		}
		if err != nil || block == nil {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if terr := expBlockTemplate.Execute(w, map[string]interface{}{
				"Error": "Block not found: " + id,
			}); terr != nil {
				http.Error(w, terr.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	if tip, terr := s.rpc.GetBlockCount(); terr == nil && tip >= block.Height {
		block.Confirmations = tip - block.Height + 1
	}

	prevHash := s.prevHashFor(block.Height)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := expBlockTemplate.Execute(w, map[string]interface{}{
		"Block":    block,
		"PrevHash": prevHash,
	}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// blockFromCache ищет блок в тёплом кэше последних блоков по height или hash.
// Только in-memory данные (recentBlocksWarm) — без сетевых вызовов: иначе
// простой запрос карточки старого блока тянул бы 50 блоков с узла.
func (s *Server) blockFromCache(id string) *Block {
	blocks := s.recentBlocksWarm(50)
	if len(blocks) == 0 {
		return nil
	}
	if h, perr := strconv.ParseInt(id, 10, 64); perr == nil {
		for _, b := range blocks {
			if b != nil && b.Height == h {
				return b
			}
		}
		return nil
	}
	id = strings.ToLower(id)
	for _, b := range blocks {
		if b != nil && strings.ToLower(b.Hash) == id {
			return b
		}
	}
	return nil
}

// prevHashFor возвращает hash предыдущего блока: сначала из тёплого кэша
// последних блоков (без сетевых вызовов), затем с узла (несколько попыток —
// узел под rate-limit).
func (s *Server) prevHashFor(height int64) string {
	if height <= 0 {
		return ""
	}
	if blocks := s.recentBlocksWarm(20); len(blocks) > 0 {
		for _, b := range blocks {
			if b != nil && b.Height == height-1 && b.Hash != "" {
				return b.Hash
			}
		}
	}
	for attempt := 0; attempt < 3; attempt++ {
		if h, herr := s.rpc.GetBlockHash(height - 1); herr == nil {
			return h
		}
		time.Sleep(300 * time.Millisecond)
	}
	return ""
}
