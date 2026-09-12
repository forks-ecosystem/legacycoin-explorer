// exp_a.go
package explorer

import (
	"html/template"
	"net/http"
)

var expATemplate = template.Must(template.New("expA").Funcs(template.FuncMap{
	"formatHashRate": formatHashRate,
}).Parse(`
<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;}
.sidebar{position:fixed;left:0;top:54px;bottom:0;z-index:90;width:220px;min-width:220px;background:#010101;border-right:1px solid #222;display:flex;flex-direction:column;transition:width .2s,min-width .2s;overflow:hidden;}
.sidebar.collapsed{width:44px;min-width:44px;border-right:none;}
.sidebar .logo{display:block;padding:8px 20px;text-align:left;}
.sidebar .logo img{height:40px;width:auto;}
.sidebar .logo:hover{opacity:.85;text-decoration:none;}
.sidebar.collapsed .logo,.sidebar.collapsed .sg{display:none;}
.sidebar .sg{display:flex;flex-direction:column;gap:1px;width:100%;min-width:0;background:none;border:none;margin:0;flex:1;overflow-y:auto;overflow-x:hidden;}
.sidebar .sc{padding:14px 10px;text-align:center;}
.sidebar .sc .sl{line-height:1.4;}
.sc{background:var(--panel);padding:16px;}
.sl{font-size:10px;text-transform:uppercase;letter-spacing:1.5px;margin-bottom:4px}
.sv{font-size:20px;font-weight:700;color:var(--gold);font-family:var(--mono);}
.ss{font-size:11px;color:var(--muted);margin-top:2px;}
.sbf{margin-top:auto;padding:10px 10px;border-top:1px solid var(--border);display:flex;align-items:center;justify-content:center;gap:8px;cursor:pointer;color:var(--muted);font-size:13px;white-space:nowrap;user-select:none;}
.sbf:hover{color:var(--gold);}
.sidebar.collapsed .sbtxt{display:none;}
</style>
<div class="sidebar" id="sb">
{{if .NodeOnline}}<div class="sg">
<div class="sc"><div class="sl">Status</div><div class="sv" style="font-size:14px;color:#22C55E;">● ONLINE</div></div>
{{with .Info}}<div class="sc"><div class="sl">Height</div><div class="sv">{{.Blocks}}</div></div>
<div class="sc"><div class="sl">Difficulty</div><div class="sv" style="font-size:14px;">{{printf "%.5f" .Difficulty}}</div><div class="ss">DGW3 per-block</div></div>{{end}}
{{with .Mining}}<div class="sc"><div class="sl">Network Hash Rate</div><div class="sv" style="font-size:14px;">{{formatHashRate $.NetHashrate}}</div></div>
<div class="sc"><div class="sl">Mempool</div><div class="sv">{{.PooledTx}}</div><div class="ss">pending txs</div></div>{{end}}
{{with .Info}}<div class="sc"><div class="sl">Peers</div><div class="sv">{{.Connections}}</div></div>{{end}}
</div>{{end}}
<div class=sbf onclick=toggleSb()><span>☰</span><span class="sbtxt">Collapse [Esc]</span></div>
</div>

<script>
function toggleSb(){
  document.getElementById('sb').classList.toggle('collapsed')
  parent.clHH()
}
</script>
`))

func (s *Server) handleExpA(w http.ResponseWriter, r *http.Request) {
	data := struct {
		NodeOnline  bool
		Info        *NodeInfo
		Mining      *MiningInfo
		NetHashrate float64
	}{}

	if s.rpc.Ping() {
		data.NodeOnline = true
		if info, err := s.cachedInfo(); err == nil {
			data.Info = info
			if d, derr := s.rpc.GetCurrentDifficulty(); derr == nil {
				data.Info.Difficulty = d
			}
		}
		if mining, err := s.cachedMining(); err == nil {
			data.Mining = mining
		}
		if hps, err := s.rpc.GetNetworkHashPerSec(); err == nil {
			data.NetHashrate = hps
		}
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	expATemplate.Execute(w, data)
}
