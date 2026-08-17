package explorer

// sharedCSS is included in every page template.
const sharedCSS = `:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#141414;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;}a{color:var(--gold);text-decoration:none;}a:hover{text-decoration:underline;}
nav{background:var(--dark);border-bottom:1px solid var(--border);padding:0 28px;display:flex;align-items:center;gap:28px;height:54px;position:sticky;top:0;z-index:100;}
.brand{font-size:17px;font-weight:700;color:var(--gold);letter-spacing:1px;white-space:nowrap;}.brand small{font-size:11px;font-weight:400;color:var(--muted);margin-left:5px;}
.navl{display:flex;gap:20px;flex:1;}.navl a{color:var(--muted);font-size:13px;}.navl a:hover{color:var(--gold);text-decoration:none;}
.sf{display:flex;margin-left:auto;}.sf input{background:var(--panel);border:1px solid var(--border);border-right:none;color:var(--text);padding:6px 12px;font-size:13px;width:280px;font-family:var(--mono);outline:none;}.sf input:focus{border-color:var(--gold);}.sf button{background:var(--gold);color:var(--black);border:none;padding:6px 14px;cursor:pointer;font-size:13px;font-weight:700;}
.c{max-width:1280px;margin:0 auto;padding:28px 22px;}
.pt{font-size:21px;font-weight:700;margin-bottom:18px;}.pt span{color:var(--gold);}
.sg{display:grid;grid-template-columns:repeat(auto-fit,minmax(150px,1fr));gap:1px;background:var(--border);border:1px solid var(--border);margin-bottom:24px;}
.sc{background:var(--panel);padding:16px;}.sl{font-size:10px;text-transform:uppercase;letter-spacing:1.5px;color:var(--muted);margin-bottom:4px;}.sv{font-size:20px;font-weight:700;color:var(--gold);font-family:var(--mono);}.ss{font-size:11px;color:var(--muted);margin-top:2px;}
.tw{border:1px solid var(--border);overflow-x:auto;}table{width:100%;border-collapse:collapse;}
thead th{background:var(--panel2);color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:10px 13px;text-align:left;border-bottom:1px solid var(--border);white-space:nowrap;}
tbody tr{border-bottom:1px solid var(--border);}tbody tr:last-child{border-bottom:none;}tbody tr:hover{background:var(--panel);}tbody td{padding:10px 13px;font-size:13px;vertical-align:middle;}
.hash{font-family:var(--mono);font-size:12px;}.bge{display:inline-block;padding:2px 7px;font-size:11px;font-weight:600;}.bg-g{background:rgba(34,197,94,.12);color:var(--green);border:1px solid rgba(34,197,94,.25);}.bg-d{background:rgba(212,160,23,.12);color:var(--gold);border:1px solid rgba(212,160,23,.25);}
.dg{display:grid;grid-template-columns:1fr 1fr;gap:18px;margin-bottom:24px;}@media(max-width:700px){.dg{grid-template-columns:1fr;}}
.dc{background:var(--panel);border:1px solid var(--border);padding:20px;}.dc h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;padding-bottom:9px;border-bottom:1px solid var(--border);}
.dr{display:flex;justify-content:space-between;gap:14px;padding:6px 0;border-bottom:1px solid rgba(255,255,255,.04);}.dr:last-child{border-bottom:none;}.dk{font-size:12px;color:var(--muted);flex-shrink:0;}.dv{font-size:12px;text-align:right;word-break:break-all;font-family:var(--mono);}.dv.gold{color:var(--gold);font-weight:700;}
.offline{background:rgba(239,68,68,.1);border:1px solid rgba(239,68,68,.3);color:var(--red);padding:13px 17px;margin-bottom:18px;font-size:14px;}
.ebox{background:var(--panel);border:1px solid var(--border);padding:36px;text-align:center;max-width:460px;margin:70px auto;}.ebox h2{font-size:19px;color:var(--gold);margin-bottom:9px;}.ebox p{color:var(--muted);font-size:14px;}.ebox a{display:inline-block;margin-top:16px;padding:9px 22px;background:var(--gold);color:var(--black);font-weight:700;}
.pg{display:flex;gap:8px;margin-top:18px;}.pg a,.pg span{padding:6px 13px;background:var(--panel);border:1px solid var(--border);font-size:13px;color:var(--muted);}.pg a:hover{border-color:var(--gold);color:var(--gold);text-decoration:none;}.pg .cur{border-color:var(--gold);color:var(--gold);}
footer{border-top:1px solid var(--border);padding:18px 28px;text-align:center;font-size:12px;color:var(--muted);margin-top:56px;}footer span{color:var(--gold);}`

const navSnip = `<nav><a href="/" class="brand">⛓ LegacyCoin <small>EXPLORER</small></a><div class="navl"><a href="/">Home</a><a href="/blocks">All Blocks</a></div><form class="sf" action="/search" method="GET"><input type="text" name="q" placeholder="Height, hash, txid or address…"><button type="submit">→</button></form></nav>`
const footSnip = `<footer><span>LegacyCoin (LBTC)</span> Block Explorer · CPU money for everyone</footer>`

const bookmarkCSS = `.bm{background:var(--panel);border:1px solid var(--border);margin-top:22px;padding:18px;}.bm h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;}.bm-list{display:flex;flex-direction:column;gap:6px;}.bm-item{display:flex;align-items:center;gap:10px;padding:8px 12px;background:var(--panel2);border:1px solid var(--border);font-size:13px;}.bm-item:hover{border-color:var(--gold);}.bm-item .bm-type{font-size:10px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);width:50px;}.bm-item .bm-val{font-family:var(--mono);font-size:12px;flex:1;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;}.bm-item .bm-del{color:var(--muted);cursor:pointer;font-size:11px;padding:2px 6px;}.bm-item .bm-del:hover{color:var(--red);}.bm-btn{display:inline-flex;align-items:center;gap:5px;padding:6px 14px;background:var(--gold);color:var(--black);font-size:12px;font-weight:700;border:none;cursor:pointer;}.bm-btn:hover{opacity:.85;}.bm-btn.saved{background:var(--panel2);color:var(--gold);border:1px solid var(--gold);}`

const allTemplates = `
{{define "home"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>LegacyCoin Explorer</title><style>` + sharedCSS + bookmarkCSS + `</style><script>
async function deleteBookmark(id) {
  await fetch('/api/bookmarks/delete/' + id, {method:'DELETE'});
  const el = document.getElementById('bm-' + id);
  if (el) el.remove();
}
async function toggleBookmark(type, value, label) {
  const id = 'bm_' + type + '_' + btoa(value).replace(/[^a-zA-Z0-9]/g,'').substring(0,20);
  const btn = document.getElementById('bm-btn');
  if (btn.classList.contains('saved')) {
    await fetch('/api/bookmarks/delete/' + id, {method:'DELETE'});
    btn.classList.remove('saved');
    btn.innerHTML = '☆ Bookmark';
  } else {
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id:id,type:type,(type==='address'?'address':'txid'):value,label:label||''})});
    btn.classList.add('saved');
    btn.innerHTML = '★ Saved';
  }
}
</script></head><body>` + navSnip + `<div class="c">
{{if .Error}}<div class="offline">⚠ {{.Error}}</div>{{end}}
{{if .NodeOnline}}<div class="sg">
<div class="sc"><div class="sl">Status</div><div class="sv" style="font-size:14px;color:#22C55E;">● ONLINE</div></div>
{{with .Info}}<div class="sc"><div class="sl">Height</div><div class="sv">{{.Blocks}}</div></div>
<div class="sc"><div class="sl">Difficulty</div><div class="sv" style="font-size:14px;">{{printf "%.5f" .Difficulty}}</div><div class="ss">DGW3 per-block</div></div>
<div class="sc"><div class="sl">Peers</div><div class="sv">{{.Connections}}</div></div>{{end}}
{{with .Mining}}<div class="sc"><div class="sl">Hash Rate</div><div class="sv" style="font-size:14px;">{{.HashesPerSec}} H/s</div></div>
<div class="sc"><div class="sl">Mempool</div><div class="sv">{{.PooledTx}}</div><div class="ss">pending txs</div></div>{{end}}
</div>{{end}}
{{if .Bookmarks}}<div class="bm"><h3>★ Bookmarks</h3><div class="bm-list">
{{range .Bookmarks}}<div class="bm-item" id="bm-{{.ID}}"><span class="bm-type">{{.Type}}</span><span class="bm-val"><a href="{{if eq .Type "address"}}/address/{{.Address}}{{else}}/tx/{{.Txid}}{{end}}">{{if .Address}}{{.Address}}{{else}}{{.Txid}}{{end}}</a></span><span class="bm-del" onclick="deleteBookmark('{{.ID}}')">✕</span></div>{{end}}
</div></div>{{end}}
<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:12px;"><div class="pt">Latest <span>Blocks</span></div><a href="/blocks" style="font-size:13px;color:var(--muted);">View all →</a></div>
<div class="tw"><table><thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead><tbody>
{{range .RecentBlocks}}<tr>
<td><a href="/block/{{.Height}}">{{.Height}}</a></td>
<td class="hash"><a href="/block/{{.Hash}}">{{truncate .Hash 32}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
<td>{{len .Tx}}</td>
<td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
<td style="color:var(--muted);">{{.Size}} B</td>
<td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>{{else}}<tr><td colspan="7" style="text-align:center;color:var(--muted);padding:26px;">No blocks yet.</td></tr>{{end}}
</tbody></table></div></div>` + footSnip + `</body></html>{{end}}

{{define "blocks"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>All Blocks — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c">
<div class="pt">All <span>Blocks</span> <small style="font-size:13px;color:var(--muted);font-weight:400;">Tip: {{.Tip}}</small></div>
<div class="tw"><table><thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead><tbody>
{{range .Blocks}}<tr>
<td><a href="/block/{{.Height}}">{{.Height}}</a></td>
<td class="hash"><a href="/block/{{.Hash}}">{{truncate .Hash 40}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
<td>{{len .Tx}}</td>
<td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
<td style="color:var(--muted);">{{.Size}} B</td>
<td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>{{else}}<tr><td colspan="7" style="text-align:center;color:var(--muted);padding:26px;">No blocks.</td></tr>{{end}}
</tbody></table></div>
<div class="pg">
{{if .HasPrev}}<a href="/blocks?page={{.PrevPage}}">← Newer</a>{{else}}<span>← Newer</span>{{end}}
<span class="cur">Page {{.Page}}</span>
{{if .HasNext}}<a href="/blocks?page={{.NextPage}}">Older →</a>{{else}}<span>Older →</span>{{end}}
</div></div>` + footSnip + `</body></html>{{end}}

{{define "block"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Block {{.Block.Height}} — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c">
<div class="pt">Block <span>#{{.Block.Height}}</span> <span class="bge bg-g" style="font-size:13px;margin-left:10px;">{{.Block.Confirmations}} confs</span></div>
<div class="dg">
<div class="dc"><h3>Header</h3>
<div class="dr"><span class="dk">Height</span><span class="dv gold">{{.Block.Height}}</span></div>
<div class="dr"><span class="dk">Hash</span><span class="dv">{{truncate .Block.Hash 44}}</span></div>
<div class="dr"><span class="dk">Previous</span><span class="dv"><a href="/block/{{.Block.PreviousBlockHash}}">{{truncate .Block.PreviousBlockHash 36}}</a></span></div>
<div class="dr"><span class="dk">Merkle Root</span><span class="dv">{{truncate .Block.MerkleRoot 36}}</span></div>
<div class="dr"><span class="dk">Time</span><span class="dv">{{formatTime .Block.Time}}</span></div>
<div class="dr"><span class="dk">nBits</span><span class="dv">{{.Block.Bits}}</span></div>
<div class="dr"><span class="dk">Nonce</span><span class="dv">{{.Block.Nonce}}</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Block.Version}}</span></div>
</div>
<div class="dc"><h3>Summary</h3>
<div class="dr"><span class="dk">Transactions</span><span class="dv">{{len .Block.Tx}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Block.Size}} bytes</span></div>
<div class="dr"><span class="dk">Block Reward</span><span class="dv gold">{{formatLBTC .Reward}}</span></div>
<div class="dr"><span class="dk">Confirmations</span><span class="dv">{{.Block.Confirmations}}</span></div>
<div class="dr"><span class="dk">Algorithm</span><span class="dv">Yespower 1.0</span></div>
<div class="dr"><span class="dk">Difficulty Algo</span><span class="dv">DGW3</span></div>
</div>
</div>
<div class="pt" style="font-size:16px;margin-bottom:10px;">Transactions <span>({{len .Block.Tx}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>TXID</th><th>Type</th></tr></thead><tbody>
{{range $i, $tx := .Block.Tx}}<tr>
<td style="color:var(--muted);">{{$i}}</td>
<td class="hash">{{$tx}}</td>
<td>{{if eq $i 0}}<span class="bge bg-d">Coinbase</span>{{else}}<span class="bge bg-g">Transfer</span>{{end}}</td>
</tr>{{end}}
</tbody></table></div>
<div style="margin-top:18px;display:flex;gap:14px;">
{{if gt .Block.Height 0}}<a href="/block/{{sub .Block.Height 1}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">← Block {{sub .Block.Height 1}}</a>{{end}}
<a href="/block/{{add .Block.Height 1}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">Block {{add .Block.Height 1}} →</a>
</div></div>` + footSnip + `</body></html>{{end}}

{{define "error"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Not Found — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c"><div class="ebox"><h2>Not Found</h2><p>{{.Message}}</p><a href="/">← Home</a></div></div>` + footSnip + `</body></html>{{end}}

{{define "tx"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Transaction {{.Tx.Txid}} — LegacyCoin Explorer</title><style>` + sharedCSS + bookmarkCSS + `</style><script>
async function toggleBookmark(type, value, label) {
  const id = 'bm_' + type + '_' + btoa(value).replace(/[^a-zA-Z0-9]/g,'').substring(0,20);
  const btn = document.getElementById('bm-btn');
  if (btn.classList.contains('saved')) {
    await fetch('/api/bookmarks/delete/' + id, {method:'DELETE'});
    btn.classList.remove('saved');
    btn.innerHTML = '☆ Bookmark';
  } else {
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id:id,type:type,(type==='address'?'address':'txid'):value,label:label||''})});
    btn.classList.add('saved');
    btn.innerHTML = '★ Saved';
  }
}
</script></head><body>` + navSnip + `<div class="c">
<div style="display:flex;justify-content:space-between;align-items:flex-start;">
<div class="pt">Transaction <span style="font-size:14px;">{{truncate .Tx.Txid 44}}</span></div>
<button id="bm-btn" class="bm-btn" onclick="toggleBookmark('tx','{{.Tx.Txid}}','')">☆ Bookmark</button>
</div>
<div class="dg">
<div class="dc"><h3>Overview</h3>
<div class="dr"><span class="dk">TXID</span><span class="dv">{{truncate .Tx.Txid 44}}</span></div>
<div class="dr"><span class="dk">Block</span><span class="dv gold"><a href="/block/{{.Block.Height}}">{{.Block.Height}}</a></span></div>
<div class="dr"><span class="dk">Confirmations</span><span class="dv">{{.Tx.Confirmations}}</span></div>
<div class="dr"><span class="dk">Time</span><span class="dv">{{formatTime .Tx.Time}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Tx.Size}} bytes</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Tx.Version}}</span></div>
</div>
<div class="dc"><h3>Summary</h3>
<div class="dr"><span class="dk">Inputs</span><span class="dv">{{len .Tx.Vin}}</span></div>
<div class="dr"><span class="dk">Outputs</span><span class="dv">{{len .Tx.Vout}}</span></div>
</div>
</div>

<div class="pt" style="font-size:16px;margin-bottom:10px;">Inputs <span>({{len .Tx.Vin}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>Previous TX</th><th>Vout</th><th>Sequence</th></tr></thead><tbody>
{{range $i, $in := .Tx.Vin}}<tr>
<td style="color:var(--muted);">{{$i}}</td>
<td class="hash"><a href="/tx/{{$in.Txid}}">{{truncate $in.Txid 32}}</a></td>
<td>{{$in.Vout}}</td>
<td style="color:var(--muted);">{{$in.Sequence}}</td>
</tr>{{end}}
</tbody></table></div>

<div class="pt" style="font-size:16px;margin:18px 0 10px;">Outputs <span>({{len .Tx.Vout}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>Value (LBTC)</th><th>Address</th><th>Type</th></tr></thead><tbody>
{{range $out := .Tx.Vout}}<tr>
<td style="color:var(--muted);">{{$out.N}}</td>
<td style="color:var(--gold);font-family:var(--mono);">{{printf "%.8f" $out.Value}}</td>
<td class="hash">{{range $addr := $out.ScriptPubKey.Addresses}}<a href="/address/{{$addr}}">{{$addr}}</a>{{end}}</td>
<td><span class="bge bg-g">{{$out.ScriptPubKey.Type}}</span></td>
</tr>{{end}}
</tbody></table></div>

<div style="margin-top:18px;display:flex;gap:14px;">
<a href="/block/{{.Block.Height}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">← Block {{.Block.Height}}</a>
</div></div>` + footSnip + `</body></html>{{end}}

{{define "address"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Address {{.Address.Address}} — LegacyCoin Explorer</title><style>` + sharedCSS + bookmarkCSS + `</style><script>
async function toggleBookmark(type, value, label) {
  const id = 'bm_' + type + '_' + btoa(value).replace(/[^a-zA-Z0-9]/g,'').substring(0,20);
  const btn = document.getElementById('bm-btn');
  if (btn.classList.contains('saved')) {
    await fetch('/api/bookmarks/delete/' + id, {method:'DELETE'});
    btn.classList.remove('saved');
    btn.innerHTML = '☆ Bookmark';
  } else {
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify({id:id,type:type,(type==='address'?'address':'txid'):value,label:label||''})});
    btn.classList.add('saved');
    btn.innerHTML = '★ Saved';
  }
}
</script></head><body>` + navSnip + `<div class="c">
<div style="display:flex;justify-content:space-between;align-items:flex-start;">
<div class="pt">Address <span style="font-size:14px;">{{.Address.Address}}</span></div>
<button id="bm-btn" class="bm-btn" onclick="toggleBookmark('address','{{.Address.Address}}','')">☆ Bookmark</button>
</div>
<div class="dg">
<div class="dc"><h3>Info</h3>
<div class="dr"><span class="dk">Address</span><span class="dv">{{.Address.Address}}</span></div>
<div class="dr"><span class="dk">Valid</span><span class="dv">{{if .Address.IsValid}}<span style="color:#22C55E;">Yes</span>{{else}}<span style="color:#EF4444;">No</span>{{end}}</span></div>
<div class="dr"><span class="dk">Is Mine</span><span class="dv">{{if .Address.IsMine}}Yes{{else}}No{{end}}</span></div>
<div class="dr"><span class="dk">Script</span><span class="dv">{{if .Address.IsScript}}Yes{{else}}No{{end}}</span></div>
</div>
<div class="dc"><h3>Summary</h3>
<div class="dr"><span class="dk">Transactions</span><span class="dv gold">{{len .Txs}}</span></div>
</div>
</div>

<div class="pt" style="font-size:16px;margin-bottom:10px;">Transactions <span>({{len .Txs}})</span></div>
<div class="tw"><table><thead><tr><th>TXID</th><th>Block</th><th>Time</th><th>Confs</th><th>Outputs</th></tr></thead><tbody>
{{range $tx := .Txs}}<tr>
<td class="hash"><a href="/tx/{{$tx.Txid}}">{{truncate $tx.Txid 32}}</a></td>
<td><a href="/block/{{$tx.Height}}}}">{{$tx.Height}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime $tx.Time}}</td>
<td><span class="bge bg-g">{{$tx.Confirmations}}</span></td>
<td>{{len $tx.Vout}}</td>
</tr>{{end}}
</tbody></table></div></div>` + footSnip + `</body></html>{{end}}
`
