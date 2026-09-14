// exp_address.go — сервисная карточка адреса (/exp-address/{addr}) во фрейме fC.
// В стиле карточки блока (exp_block.go): Header / Balance / Transactions.
// Без чужого навбара и редиректов.
package explorer

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var expAddressTemplate = template.Must(template.New("expAddress").Funcs(template.FuncMap{
	"formatTime": formatTime,
	"truncate":   truncate,
	"lbtc":       func(v float64) string { return fmt.Sprintf("%.8f LBTC", v) },
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
.bg-d{background:rgba(212,160,23,.12);color:var(--gold);border:1px solid rgba(212,160,23,.25);}
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
.foot{flex-shrink:0;display:flex;justify-content:space-between;gap:12px;padding:12px 22px;font-size:13px;}
</style>
<script>
function home(){ if(window.parent&&window.parent!==window&&typeof window.parent.lf==='function'){window.parent.lf('/exp-1');}else{location.href='/';} }
</script>
{{if .Error}}<div class="tk"><span>Address</span></div>
<div class="sc"><div class="in"><div class="offline">⚠ {{.Error}}<a href="javascript:home()">← Home</a></div></div></div>
{{else}}<div class="tk"><span>Address</span> <b>{{truncate .Address.Address 30}}</b> {{if .Address.IsMine}}<span class="bge bg-d">mine</span>{{end}}{{if .Address.IsWatchOnly}}<span class="bge bg-d">watch-only</span>{{end}}</div>
<div class="sc"><div class="in">
<div class="dg">
<div class="dc"><h3>Header</h3>
<div class="dr"><span class="dk">Address</span><span class="dv">{{.Address.Address}}</span></div>
<div class="dr"><span class="dk">Valid</span><span class="dv">{{if .Address.IsValid}}Yes{{else}}No{{end}}</span></div>
<div class="dr"><span class="dk">Is Mine</span><span class="dv">{{if .Address.IsMine}}Yes{{else}}No{{end}}</span></div>
<div class="dr"><span class="dk">Is Script</span><span class="dv">{{if .Address.IsScript}}Yes{{else}}No{{end}}</span></div>
<div class="dr"><span class="dk">Watch Only</span><span class="dv">{{if .Address.IsWatchOnly}}Yes{{else}}No{{end}}</span></div>
<div class="dr"><span class="dk">PubKey Hash</span><span class="dv">{{if .Address.PubKeyHash}}{{.Address.PubKeyHash}}{{else}}—{{end}}</span></div>
</div>
<div class="dc"><h3>Balance</h3>
{{if .Balance}}
<div class="dr"><span class="dk">Balance</span><span class="dv gold">{{lbtc .Balance.Balance}}</span></div>
<div class="dr"><span class="dk">Received</span><span class="dv">{{lbtc .Balance.Received}}</span></div>
{{else}}
<div class="dr"><span class="dk">Balance</span><span class="dv">—</span></div>
{{end}}
<div class="dr"><span class="dk">Transactions</span><span class="dv gold">{{len .Txs}}</span></div>
</div>
</div>
<div class="pt">Transactions <span>({{len .Txs}})</span></div>
<div class="tw"><table><thead><tr><th>TXID</th><th>Block</th><th>Time</th><th>Confs</th><th>Outputs</th></tr></thead><tbody>
{{range $tx := .Txs}}<tr>
<td class="hash"><a href="/tx/{{$tx.Txid}}" target="_top">{{truncate $tx.Txid 32}}</a></td>
<td>{{if $tx.Blockhash}}<a href="/block/{{$tx.Blockhash}}" target="_top">{{$tx.Height}}</a>{{else}}mempool{{end}}</td>
<td style="color:var(--muted);font-size:12px;">{{formatTime $tx.Time}}</td>
<td><span class="bge bg-g">{{$tx.Confirmations}}</span></td>
<td style="color:var(--muted);">{{len $tx.Vout}}</td>
</tr>{{else}}
<tr><td colspan="5" style="text-align:center;color:var(--muted);padding:22px;">No transactions found.</td></tr>
{{end}}
</tbody></table></div>
</div></div>
<div class="foot">
<a href="javascript:home()">← Latest</a>
</div>
{{end}}
`))

func (s *Server) handleExpAddress(w http.ResponseWriter, r *http.Request) {
	addr := strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/exp-address/"))
	if !isAddress(addr) {
		s.renderExpAddressError(w, "Invalid address: "+addr)
		return
	}
	info, err := s.rpc.ValidateAddress(addr)
	if err != nil || info == nil || !info.IsValid {
		s.renderExpAddressError(w, "Address not found or validation failed: "+addr)
		return
	}
	bal, _ := s.rpc.GetAddressBalance(addr)
	txs, _ := s.rpc.FindAddressTxs(addr, 1000)
	s.renderExpAddress(w, map[string]interface{}{
		"Address": info,
		"Balance": bal,
		"Txs":     txs,
	})
}

func (s *Server) renderExpAddressError(w http.ResponseWriter, msg string) {
	s.renderExpAddress(w, map[string]interface{}{"Error": msg})
}

func (s *Server) renderExpAddress(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := expAddressTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}