// exp_tx.go — сервисная карточка транзакции (/exp-tx/{txid}) во фрейме fC.
// В стиле карточки блока (exp_block.go): Header / Inputs / Outputs.
// Без чужого навбара и редиректов.
package explorer

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"
)

var expTxTemplate = template.Must(template.New("expTx").Funcs(template.FuncMap{
	"formatTime": formatTime,
	"truncate":   truncate,
	"lbtc":       func(v float64) string { return fmt.Sprintf("%.8f LBTC", v) },
	"sumOut": func(vout []TxOutput) float64 {
		var total float64
		for _, v := range vout {
			total += v.Value
		}
		return total
	},
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
{{if .Error}}<div class="tk"><span>Transaction</span></div>
<div class="sc"><div class="in"><div class="offline">⚠ {{.Error}}<a href="javascript:home()">← Home</a></div></div></div>
{{else}}<div class="tk"><span>Transaction</span> <b>{{truncate .Tx.Txid 24}}</b> {{if .Tx.Confirmations}}<span class="bge bg-g">{{.Tx.Confirmations}} confs</span>{{else}}<span class="bge bg-g">mempool</span>{{end}}</div>
<div class="sc"><div class="in">
<div class="dg">
<div class="dc"><h3>Header</h3>
<div class="dr"><span class="dk">TXID</span><span class="dv">{{.Tx.Txid}}</span></div>
<div class="dr"><span class="dk">Hash</span><span class="dv">{{.Tx.Hash}}</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Tx.Version}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Tx.Size}} B</span></div>
{{if .Tx.Blockhash}}<div class="dr"><span class="dk">Block</span><a class="dv" href="/block/{{.Tx.Blockhash}}" target="_top">{{truncate .Tx.Blockhash 50}}</a></div>{{end}}
<div class="dr"><span class="dk">Confirmations</span><span class="dv">{{if .Tx.Confirmations}}{{.Tx.Confirmations}}{{else}}0 (mempool){{end}}</span></div>
<div class="dr"><span class="dk">Time</span><span class="dv">{{formatTime .Tx.Time}}</span></div>
<div class="dr"><span class="dk">Blocktime</span><span class="dv">{{formatTime .Tx.Blocktime}}</span></div>
</div>
<div class="dc"><h3>Summary</h3>
<div class="dr"><span class="dk">Outputs</span><span class="dv">{{len .Tx.Vout}}</span></div>
<div class="dr"><span class="dk">Total Output</span><span class="dv gold">{{lbtc (sumOut .Tx.Vout)}}</span></div>
<div class="dr"><span class="dk">Inputs</span><span class="dv">{{len .Tx.Vin}}</span></div>
<div class="dr"><span class="dk">Confirmed In</span><span class="dv">{{if .Tx.Blockhash}}Block {{.Tx.Height}}{{else}}— (unconfirmed){{end}}</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Tx.Version}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Tx.Size}} B</span></div>
</div>
</div>
<div class="pt">Inputs <span>({{len .Tx.Vin}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>Previous TX</th><th>Vout</th><th>Sequence</th></tr></thead><tbody>
{{range $i, $in := .Tx.Vin}}
<tr>
<td style="color:var(--muted);">{{$i}}</td>
{{if eq $in.Txid ""}}<td class="hash" colspan="2"><span class="bge bg-d">Coinbase (new coins)</span></td>{{else}}<td class="hash"><a href="/tx/{{$in.Txid}}" target="_top">{{truncate $in.Txid 50}}</a></td><td style="color:var(--muted);">{{$in.Vout}}</td>{{end}}
<td style="color:var(--muted);">{{$in.Sequence}}</td>
</tr>
{{else}}
<tr><td colspan="4" style="text-align:center;color:var(--muted);padding:22px;">No inputs.</td></tr>
{{end}}
</tbody></table></div>
<div class="pt" style="margin-top:24px;">Outputs <span>({{len .Tx.Vout}})</span></div>
<div class="tw"><table><thead><tr><th>N</th><th>Address</th><th>Value</th><th>Type</th></tr></thead><tbody>
{{range $o := .Tx.Vout}}
<tr>
<td style="color:var(--muted);">{{$o.N}}</td>
<td class="hash">{{range $i, $a := $o.ScriptPubKey.Addresses}}{{if $i}}, {{end}}<a href="/address/{{$a}}" target="_top">{{$a}}</a>{{else}}—{{end}}</td>
<td style="color:var(--gold);font-family:var(--mono);">{{lbtc $o.Value}}</td>
<td style="color:var(--muted);">{{$o.ScriptPubKey.Type}}</td>
</tr>
{{else}}
<tr><td colspan="4" style="text-align:center;color:var(--muted);padding:22px;">No outputs.</td></tr>
{{end}}
</tbody></table></div>
</div></div>
<div class="foot">
<a href="javascript:home()">← Latest</a>
{{if .Block}}<a href="/block/{{.Block.Height}}" target="_top">Block #{{.Block.Height}}</a>{{end}}
</div>
{{end}}
`))

func (s *Server) handleExpTx(w http.ResponseWriter, r *http.Request) {
	txid := strings.ToLower(strings.TrimSpace(strings.TrimPrefix(r.URL.Path, "/exp-tx/")))
	if !isTxid(txid) {
		s.renderExpTxError(w, "Invalid transaction ID: "+txid)
		return
	}
	tx, block, err := s.rpc.FindTransaction(txid)
	if err != nil || tx == nil {
		s.renderExpTxError(w, "Transaction not found: "+txid)
		return
	}
	s.renderExpTx(w, map[string]interface{}{
		"Tx":    tx,
		"Block": block,
	})
}

func (s *Server) renderExpTxError(w http.ResponseWriter, msg string) {
	s.renderExpTx(w, map[string]interface{}{"Error": msg})
}

func (s *Server) renderExpTx(w http.ResponseWriter, data map[string]interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := expTxTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}