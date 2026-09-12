package explorer

// sharedCSS is included in every page template.
const sharedCSS = `:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;display:flex;flex-direction:column;height:100vh;overflow:hidden;padding-bottom:56px;}a{color:var(--gold);text-decoration:none;}a:hover{text-decoration:underline;}
nav{background:var(--dark);padding:0 28px;display:flex;align-items:center;gap:28px;height:54px;position:sticky;top:0;z-index:100;}
.brand{font-size:17px;font-weight:700;color:var(--gold);letter-spacing:1px;white-space:nowrap;}.brand small{font-size:11px;font-weight:400;color:var(--muted);margin-left:5px;}
.navl{display:flex;gap:20px;flex:1;}.navl a{color:var(--muted);font-size:13px;}.navl a:hover{color:var(--gold);text-decoration:none;}
.sf{display:flex;margin-left:auto;}.sf input{background:var(--panel);border:1px solid var(--border);border-right:none;color:var(--text);padding:6px 12px;font-size:13px;width:280px;font-family:var(--mono);outline:none;}.sf input:focus{border-color:var(--gold);}.sf button{background:var(--gold);color:var(--black);border:none;padding:6px 14px;cursor:pointer;font-size:13px;font-weight:700;}
.c{max-width:1280px;margin:0 auto;padding:28px 22px;flex:1;}
.pt{font-size:21px;font-weight:700;margin-bottom:0;}.pt span{color:var(--gold);}
.sg{display:flex;flex-direction:column;gap:1px;background:var(--border);border:1px solid var(--border);width:260px;flex-shrink:0;margin-bottom:24px;}
.wrap{display:flex;width:100%;flex:1;min-height:0;}
.sidebar{position:fixed;left:0;top:54px;bottom:0;z-index:90;width:220px;min-width:220px;background:#010101;border-right:1px solid #222;display:flex;flex-direction:column;transition:width .2s,min-width .2s;overflow:hidden;}
.sidebar.collapsed{width:44px;min-width:44px;border-right:none;}
.sidebar.collapsed~.main{margin-left:44px;}
.sidebar .logo{display:block;padding:8px 20px;text-align:left;}
.sidebar .logo img{height:40px;width:auto;}
.sidebar .logo:hover{opacity:.85;text-decoration:none;}
.sidebar.collapsed .logo,.sidebar.collapsed .sg{display:none;}
.sidebar .sg{display:flex;flex-direction:column;gap:1px;width:100%;min-width:0;background:none;border:none;margin:0;flex:1;overflow-y:auto;overflow-x:hidden;}
.sidebar .sc{_background:var(--panel);padding:14px 10px;text-align:center;}
.sidebar .sc .sl{line-height:1.4;}
.sbf{margin-top:auto;padding:17px 10px;border-top:1px solid var(--border);display:flex;align-items:center;justify-content:center;gap:8px;cursor:pointer;color:var(--muted);font-size:13px;white-space:nowrap;user-select:none;}
.sbf:hover{color:var(--gold);}
.sidebar.collapsed .sbtxt{display:none;}
.main{flex:1;min-width:0;min-height:0;margin-left:220px;transition:margin-left .2s;display:flex;flex-direction:column;}
.main>.c{max-width:none;margin:0;padding:9px 22px;flex:1;min-height:0;display:flex;flex-direction:column;}
@media(max-width:800px){body{height:auto;overflow:visible;}.wrap{flex-direction:column;}.sidebar{position:static;width:100%;min-width:0;border-right:none;border-bottom:1px solid #222;flex-direction:row;flex-wrap:wrap;}.sidebar .logo{flex:0 0 auto;border-bottom:none;border-right:1px solid var(--border);}.sidebar .sg{width:100%;padding-top:0;}.main{margin-left:0;display:block;}.main>.c{display:block;flex:none;}.tw{flex:none;overflow:visible;}}
.sc{background:var(--panel);padding:16px;}.sl{font-size:10px;text-transform:uppercase;letter-spacing:1.5px;margin-bottom:4px;}.sv{font-size:20px;font-weight:700;color:var(--gold);font-family:var(--mono);}.ss{font-size:11px;color:var(--muted);margin-top:2px;}
.tw{flex:1;min-height:0;overflow:auto;}table{width:auto;max-width:100%;border-collapse:separate;border-spacing:0;}
thead th{position:sticky;top:0;z-index:5;background:var(--panel2);color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:10px 13px;text-align:left;border-bottom:1px solid var(--border);white-space:nowrap;}
tbody tr{border-bottom:1px solid var(--border);}tbody tr:last-child{border-bottom:none;}tbody tr:hover{background:var(--panel);}tbody tr.sel{background:var(--panel);box-shadow:inset 3px 0 0 #3B82F6;}tbody td{padding:10px 13px;font-size:13px;vertical-align:middle;}tbody tr.spacer td{border:none;padding:0;height:16px;}tbody tr.spacer:hover{background:none;}
.hash{font-family:var(--mono);font-size:12px;}.bge{display:inline-block;padding:2px 7px;font-size:11px;font-weight:600;}.bg-g{background:rgba(34,197,94,.12);color:var(--green);border:1px solid rgba(34,197,94,.25);}.bg-d{background:rgba(212,160,23,.12);color:var(--gold);border:1px solid rgba(212,160,23,.25);}
.dg{display:grid;grid-template-columns:1fr 1fr;gap:18px;margin-bottom:24px;}@media(max-width:700px){.dg{grid-template-columns:1fr;}}
.dc{background:var(--panel);border:1px solid var(--border);padding:20px;}.dc h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;padding-bottom:9px;border-bottom:1px solid var(--border);}
.dr{display:flex;justify-content:space-between;gap:14px;padding:6px 0;border-bottom:1px solid rgba(255,255,255,.04);}.dr:last-child{border-bottom:none;}.dk{font-size:12px;color:var(--muted);flex-shrink:0;}.dv{font-size:12px;text-align:right;word-break:break-all;font-family:var(--mono);}.dv.gold{color:var(--gold);font-weight:700;}
.offline{background:rgba(239,68,68,.1);border:1px solid rgba(239,68,68,.3);color:var(--red);padding:13px 17px;margin-bottom:18px;font-size:14px;}
.ebox{background:var(--panel);border:1px solid var(--border);padding:36px;text-align:center;max-width:460px;margin:70px auto;}.ebox h2{font-size:19px;color:var(--gold);margin-bottom:9px;}.ebox p{color:var(--muted);font-size:14px;}.ebox a{display:inline-block;margin-top:16px;padding:9px 22px;background:var(--gold);color:var(--black);font-weight:700;}
.pg{display:flex;gap:8px;margin-top:18px;}.pg a,.pg span{padding:6px 13px;background:var(--panel);border:1px solid var(--border);font-size:13px;color:var(--muted);}.pg a:hover{border-color:var(--gold);color:var(--gold);text-decoration:none;}.pg .cur{border-color:var(--gold);color:var(--gold);}
.bl{display:flex;gap:18px;flex:1;min-height:0;}.bl-m{flex:1;min-width:0;min-height:0;display:flex;flex-direction:column;}.bl-s{width:320px;flex-shrink:0;}@media(max-width:800px){.bl{flex-direction:column;}.bl-s{width:100%;}}
footer{position:fixed;bottom:0;left:0;right:0;z-index:80;background:var(--dark);border-top:1px solid var(--border);padding:18px 28px;text-align:center;font-size:12px;color:var(--muted);}footer span{color:var(--gold);}
.lcCard{position:fixed;top:66px;right:22px;bottom:60px;width:340px;min-width:160px;background:var(--dark);border:1px solid var(--gold);border-radius:8px;z-index:120;display:none;flex-direction:column;box-shadow:0 10px 40px rgba(0,0,0,.65);}
.lcHead{display:flex;align-items:center;justify-content:space-between;gap:10px;padding:11px 14px;border-bottom:1px solid var(--border);color:var(--gold);font-weight:700;font-size:12px;text-transform:uppercase;letter-spacing:1px;flex-shrink:0;}
.lcX{cursor:pointer;color:var(--muted);font-size:15px;padding:0 4px;}.lcX:hover{color:var(--red);}
.lcBody{padding:10px 14px;overflow-y:auto;flex:1;min-height:0;font-size:12px;}
.lcCard .ld{display:flex;justify-content:space-between;gap:12px;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.05);}.lcCard .ld span{color:var(--muted);flex-shrink:0;}.lcCard .ld b{font-family:var(--mono);text-align:right;word-break:break-all;}
.lcRes{position:absolute;left:-5px;top:0;bottom:0;width:12px;cursor:ew-resize;background:rgba(255,255,255,.07);border-left:1px dashed rgba(212,160,23,.45);}
.tw.hdcur thead tr{box-shadow:inset 3px 0 0 #3B82F6;}`

const navSnip = `<nav><a href="/" class="brand">⛓ LegacyCoin <small>EXPLORER</small></a><div class="navl"><a href="/">1. Home</a><a href="/blocks">2. All Blocks</a></div><form class="sf" action="/search" method="GET"><input type="text" name="q" placeholder="Height, hash, txid or address…"><button type="submit">→</button></form></nav>`
const footSnip = `<script>
var lcCard=null;
function lcInit(){
  if(lcCard)return;
  var d=document.createElement('div');
  d.className='lcCard';
  d.innerHTML='<div class="lcHead"><span>Block card</span><span class="lcX" onclick="lcClose()">✕</span></div><div class="lcBody"></div><div class="lcRes"></div>';
  document.body.appendChild(d);
  lcCard=d;
  var res=d.querySelector('.lcRes'),sx=0,sw=0;
  res.addEventListener('mousedown',function(ev){ev.preventDefault();sx=ev.clientX;sw=d.offsetWidth;function mv(m){var maxW=Math.max(240,window.innerWidth-60);var w=sw-(m.clientX-sx);if(w<160)w=160;if(w>maxW)w=maxW;d.style.width=w+'px';}function up(){document.removeEventListener('mousemove',mv);document.removeEventListener('mouseup',up);}document.addEventListener('mousemove',mv);document.addEventListener('mouseup',up);});
}
function lcIsOpen(){return lcCard&&lcCard.style.display!=='none';}
function lcClose(){if(lcCard)lcCard.style.display='none';}
var lcSeq=0,lcPrevT=null;
function lcLoad(h,retry){
  lcInit();
  var seq=++lcSeq;
  var d=lcCard,body=d.querySelector('.lcBody');
  fetch('/api/block/'+h).then(function(r){return r.json();}).then(function(b){
    if(seq!==lcSeq)return;
    if(!b||b.error){
      if(retry<2){setTimeout(function(){lcLoad(h,retry+1);},350);return;}
      body.innerHTML='<div style="color:#EF4444;padding:8px;">node busy — press Enter to retry</div>';return;
    }
    body.innerHTML=''
      +'<div class="ld"><span>Height</span><b>'+b.height+'</b></div>'
      +'<div class="ld"><span>Hash</span><b>'+b.hash+'</b></div>'
      +'<div class="ld"><span>Time</span><b>'+new Date((b.time||0)*1000).toUTCString()+'</b></div>'
      +'<div class="ld"><span>Transactions</span><b>'+(b.tx?b.tx.length:'-')+'</b></div>'
      +'<div class="ld"><span>Size</span><b>'+(b.size||'-')+' B</b></div>'
      +'<div class="ld"><span>Difficulty</span><b>'+(b.difficulty?b.difficulty.toFixed(6):'-')+'</b></div>'
      +'<div class="ld"><span>nBits</span><b>'+b.bits+'</b></div>'
      +'<div class="ld"><span>Nonce</span><b>'+b.nonce+'</b></div>'
      +'<div class="ld"><span>Confirmations</span><b>'+b.confirmations+'</b></div>'
      +'<div style="margin-top:12px"><a href="/block/'+b.height+'">Open block page →</a></div>';
  }).catch(function(){if(seq===lcSeq&&retry<2)setTimeout(function(){lcLoad(h,retry+1);},350);});
}
function lcOpen(h){
  lcInit();
  lcCard.querySelector('.lcHead span').textContent='Block #'+h;
  lcCard.querySelector('.lcBody').innerHTML='<div style="color:var(--muted);padding:8px;">Loading…</div>';
  lcCard.style.display='flex';
  lcLoad(h,0);
}
function lcRows(){return document.querySelectorAll('tbody tr:not(.spacer)');}
function lcSelRow(){return document.querySelector('tbody tr.sel');}
function lcScroll(r){
  var tw=r.closest('.tw');if(!tw)return;
  var th=tw.querySelector('thead');var oh=th?th.offsetHeight+2:0;
  var rowH=r.offsetHeight||36;
  var pad=rowH*1.6;
  var y=0,el=r;
  while(el&&el!==tw){y+=el.offsetTop;el=el.offsetParent;if(!el)break;}
  var twH=tw.clientHeight;
  if(y<tw.scrollTop+oh+pad)tw.scrollTop=y-oh-pad;
  else if(y+rowH>tw.scrollTop+twH-pad)tw.scrollTop=y+rowH-twH+pad;
}
function lcSelect(r){
  var old=lcSelRow();if(old&&old!==r)old.classList.remove('sel');
  r.classList.add('sel');lcScroll(r);r.focus&&r.focus();
  var t=r.closest('.tw');if(t)t.classList.remove('hdcur');
  var hh=r.getAttribute('data-height');
  if(hh&&lcIsOpen()){clearTimeout(lcPrevT);lcPrevT=setTimeout((function(h){return function(){lcLoad(h,0);};})(hh),400);}
}
function lcNav(dir,big){
  var rows=lcRows();if(!rows.length)return;
  var i=[].indexOf.call(rows,lcSelRow());
  var step=1;
  if(big){var rowH=(rows[0]&&rows[0].offsetHeight)||36;var tw=rows[0].closest('.tw');var vis=Math.max(1,Math.floor((tw?tw.clientHeight:window.innerHeight)/rowH)-1);step=vis;}
  i=dir>0?i+step:i-step;
  if(i<0)i=0;if(i>=rows.length)i=rows.length-1;
  lcSelect(rows[i]);
}
function lcEnd(last){var rows=lcRows();if(!rows.length)return;lcSelect(rows[last?rows.length-1:0]);}
(function(){var t=document.querySelector('.tw');if(t)t.classList.add('hdcur');})();
document.addEventListener('click',function(e){
  var t=document.querySelector('.tw');if(t)t.classList.remove('hdcur');
  var tr=e.target&&e.target.closest?e.target.closest('tbody tr:not(.spacer)'):null;
  if(tr)lcSelect(tr);
});
document.addEventListener('keydown',function(e){
  if(e.target.tagName==='INPUT'||e.target.tagName==='TEXTAREA'){
    if(e.key==='Escape'||e.keyCode===27){
      e.preventDefault();
      if(e.target.value){e.target.value='';}
      else e.target.blur();
    }
    return;
  }
  var k=e.key,c=e.keyCode;
  if(k==='1'||c===49||c===97){location.href='/';return;}
  if(k==='2'||c===50||c===98){location.href='/blocks';return;}
  if(k==='Backspace'||c===8){var fi=document.querySelector('.sf input');if(fi){e.preventDefault();fi.focus();}return;}
  if(k==='Enter'||c===13){var sl=document.querySelector('tbody tr.sel[data-height]');if(sl){e.preventDefault();lcOpen(sl.getAttribute('data-height'));}return;}
  if(k==='Escape'||c===27){if(lcIsOpen()){lcClose();return;}var sb=document.querySelector('.sidebar');if(sb)sb.classList.toggle('collapsed');return;}
  if(k==='ArrowDown'||c===40){e.preventDefault();lcNav(1,false);return;}
  if(k==='ArrowUp'||c===38){e.preventDefault();lcNav(-1,false);return;}
  if(k==='PageDown'||c===34){e.preventDefault();lcNav(1,true);return;}
  if(k==='PageUp'||c===33){e.preventDefault();lcNav(-1,true);return;}
  if(k==='Home'||c===36){e.preventDefault();lcEnd(false);return;}
  if(k==='End'||c===35){e.preventDefault();lcEnd(true);return;}
  if(k==='ArrowLeft'||c===37){var l=document.querySelector('[data-dir="prev"]');if(l&&l.href){e.preventDefault();location.href=l.href;}return;}
  if(k==='ArrowRight'||c===39){var l=document.querySelector('[data-dir="next"]');if(l&&l.href){e.preventDefault();location.href=l.href;}return;}
});
</script><footer><span>LegacyCoin (LBTC)</span> Block Explorer · CPU money for everyone</footer>`

const bookmarkCSS = `.bm{background:var(--panel);border:1px solid var(--border);margin-top:22px;padding:18px;}.bm h3{font-size:11px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);margin-bottom:12px;}.bm-list{display:flex;flex-direction:column;gap:6px;}.bm-item{display:flex;align-items:center;gap:10px;padding:8px 12px;background:var(--panel2);border:1px solid var(--border);font-size:13px;}.bm-item:hover{border-color:var(--gold);}.bm-item .bm-type{font-size:10px;text-transform:uppercase;letter-spacing:1px;color:var(--muted);width:50px;}.bm-item .bm-val{font-family:var(--mono);font-size:12px;flex:1;overflow:hidden;text-overflow:ellipsis;white-space:nowrap;}.bm-item .bm-del{color:var(--muted);cursor:pointer;font-size:11px;padding:2px 6px;}.bm-item .bm-del:hover{color:var(--red);}.bm-btn{display:inline-flex;align-items:center;gap:5px;padding:6px 14px;background:var(--gold);color:var(--black);font-size:12px;font-weight:700;border:none;cursor:pointer;}.bm-btn:hover{opacity:.85;}.bm-btn.saved{background:var(--panel2);color:var(--gold);border:1px solid var(--gold);}`

const allTemplates = `
{{define "home"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="icon" href="data:,"><title>LegacyCoin Explorer</title><style>` + sharedCSS + bookmarkCSS + `</style><script>
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
    var rec={id:id,type:type,label:label||''};rec[type==='address'?'address':'txid']=value;
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(rec)});
    btn.classList.add('saved');
    btn.innerHTML = '★ Saved';
  }
}
function toggleSidebar(){var sb=document.querySelector('.sidebar');if(sb)sb.classList.toggle('collapsed');}
</script></head><body>` + navSnip + `<div class="wrap">
<div class="sidebar">
<a href="/" class="logo" title="Home"><img src="/lbtc.png" alt="LBTC"></a>
{{if .NodeOnline}}<div class="sg">
<div class="sc"><div class="sl">Status</div><div class="sv" style="font-size:14px;color:#22C55E;">● ONLINE</div></div>
{{with .Info}}<div class="sc"><div class="sl">Height</div><div class="sv">{{.Blocks}}</div></div>
<div class="sc"><div class="sl">Difficulty</div><div class="sv" style="font-size:14px;">{{printf "%.5f" .Difficulty}}</div><div class="ss">DGW3 per-block</div></div>{{end}}
{{with .Mining}}<div class="sc"><div class="sl">Network Hash Rate</div><div class="sv" style="font-size:14px;">{{formatHashRate $.NetHashrate}}</div></div>
<div class="sc"><div class="sl">Mempool</div><div class="sv">{{.PooledTx}}</div><div class="ss">pending txs</div></div>{{end}}
{{with .Info}}<div class="sc"><div class="sl">Peers</div><div class="sv">{{.Connections}}</div></div>{{end}}
</div>{{end}}
<div class="sbf" onclick="toggleSidebar()"><span>☰</span><span class="sbtxt">Collapse</span></div>
</div>
<div class="main"><div class="c">
{{if .Error}}<div class="offline">⚠ {{.Error}}</div>{{end}}
{{if .Bookmarks}}<div class="bm"><h3>★ Bookmarks</h3><div class="bm-list">
{{range .Bookmarks}}<div class="bm-item" id="bm-{{.ID}}"><span class="bm-type">{{.Type}}</span><span class="bm-val"><a href="{{if eq .Type "address"}}/address/{{.Address}}{{else}}/tx/{{.Txid}}{{end}}">{{if .Address}}{{.Address}}{{else}}{{.Txid}}{{end}}</a></span><span class="bm-del" onclick="deleteBookmark('{{.ID}}')">✕</span></div>{{end}}
</div></div>{{end}}
<div style="display:flex;justify-content:space-between;align-items:center;margin-bottom:12px;"><div class="pt">Latest <span>Blocks</span></div><a href="/blocks" style="font-size:13px;color:var(--muted);">View all →</a></div>
<div class="tw"><table><thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead><tbody>
{{range .RecentBlocks}}<tr data-height="{{.Height}}">
<td><a href="/block/{{.Height}}">{{.Height}}</a></td>
<td class="hash"><a href="/block/{{.Hash}}">{{truncate .Hash 32}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
<td>{{len .Tx}}</td>
<td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
<td style="color:var(--muted);">{{.Size}} B</td>
<td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>{{else}}<tr><td colspan="7" style="text-align:center;color:var(--muted);padding:26px;">No blocks yet.</td></tr>{{end}}
<tr class="spacer"><td colspan="7"></td></tr><tr class="spacer"><td colspan="7"></td></tr>
</tbody></table></div></div></div></div>` + footSnip + `</body></html>{{end}}

{{define "blocks"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="icon" href="data:,"><title>All Blocks — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c">
<div class="pt">All <span>Blocks</span> <small style="font-size:13px;color:var(--muted);font-weight:400;">Tip: {{.Tip}}</small></div>
<div class="tw"><table><thead><tr><th>Height</th><th>Hash</th><th>Time (UTC)</th><th>Txs</th><th>Reward</th><th>Size</th><th>Confs</th></tr></thead><tbody>
{{range .Blocks}}<tr data-height="{{.Height}}">
<td><a href="/block/{{.Height}}">{{.Height}}</a></td>
<td class="hash"><a href="/block/{{.Hash}}">{{truncate .Hash 40}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
<td>{{len .Tx}}</td>
<td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
<td style="color:var(--muted);">{{.Size}} B</td>
<td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>{{else}}<tr><td colspan="7" style="text-align:center;color:var(--muted);padding:26px;">No blocks.</td></tr>{{end}}
<tr class="spacer"><td colspan="7"></td></tr><tr class="spacer"><td colspan="7"></td></tr>
</tbody></table></div>
<div class="pg">
{{if .HasPrev}}<a data-dir="prev" href="/blocks?page={{.PrevPage}}">← Newer</a>{{else}}<span data-dir="prev">← Newer</span>{{end}}
<span class="cur">Page {{.Page}}</span>
{{if .HasNext}}<a data-dir="next" href="/blocks?page={{.NextPage}}">Older →</a>{{else}}<span data-dir="next">Older →</span>{{end}}
</div></div>` + footSnip + `</body></html>{{end}}

{{define "block"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="icon" href="data:,"><title>Block {{.Block.Height}} — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c">
<div class="pt">Block <span>#{{.Block.Height}}</span> <span class="bge bg-g" style="font-size:13px;margin-left:10px;">{{.Block.Confirmations}} confs</span></div>
<div class="bl">
<div class="bl-m">
<div class="pt" style="font-size:16px;margin-bottom:10px;">Transactions <span>({{len .Block.Tx}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>TXID</th><th>Type</th></tr></thead><tbody>
{{range $i, $tx := .Block.Tx}}<tr>
<td style="color:var(--muted);">{{$i}}</td>
<td class="hash">{{$tx}}</td>
<td>{{if eq $i 0}}<span class="bge bg-d">Coinbase</span>{{else}}<span class="bge bg-g">Transfer</span>{{end}}</td>
</tr>{{end}}
<tr class="spacer"><td colspan="3"></td></tr><tr class="spacer"><td colspan="3"></td></tr>
</tbody></table></div>
<div style="margin-top:18px;display:flex;gap:14px;">
{{if gt .Block.Height 0}}<a data-dir="prev" href="/block/{{sub .Block.Height 1}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">← Block {{sub .Block.Height 1}}</a>{{end}}
<a data-dir="next" href="/block/{{add .Block.Height 1}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">Block {{add .Block.Height 1}} →</a>
</div>
</div>
<div class="bl-s">
<div class="dc"><h3>Block Info</h3>
<div class="dr"><span class="dk">Height</span><span class="dv gold">{{.Block.Height}}</span></div>
<div class="dr"><span class="dk">Hash</span><span class="dv">{{truncate .Block.Hash 44}}</span></div>
<div class="dr"><span class="dk">Previous</span><span class="dv"><a href="/block/{{.Block.PreviousBlockHash}}">{{truncate .Block.PreviousBlockHash 36}}</a></span></div>
<div class="dr"><span class="dk">Merkle Root</span><span class="dv">{{truncate .Block.MerkleRoot 36}}</span></div>
<div class="dr"><span class="dk">Time</span><span class="dv">{{formatTime .Block.Time}}</span></div>
<div class="dr"><span class="dk">nBits</span><span class="dv">{{.Block.Bits}}</span></div>
<div class="dr"><span class="dk">Nonce</span><span class="dv">{{.Block.Nonce}}</span></div>
<div class="dr"><span class="dk">Version</span><span class="dv">{{.Block.Version}}</span></div>
<div class="dr"><span class="dk">Transactions</span><span class="dv">{{len .Block.Tx}}</span></div>
<div class="dr"><span class="dk">Size</span><span class="dv">{{.Block.Size}} bytes</span></div>
<div class="dr"><span class="dk">Block Reward</span><span class="dv gold">{{formatLBTC .Reward}}</span></div>
<div class="dr"><span class="dk">Confirmations</span><span class="dv">{{.Block.Confirmations}}</span></div>
<div class="dr"><span class="dk">Algorithm</span><span class="dv">Yespower 1.0</span></div>
<div class="dr"><span class="dk">Difficulty (DGW3)</span><span class="dv">{{printf "%.5f" .Block.Difficulty}}</span></div>
</div>
</div>
</div></div>` + footSnip + `</body></html>{{end}}

{{define "error"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><link rel="icon" href="data:,"><title>Not Found — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c"><div class="ebox"><h2>Not Found</h2><p>{{.Message}}</p><a href="/">← Home</a></div></div>` + footSnip + `</body></html>{{end}}

{{define "tx"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title>Transaction {{.Tx.Txid}} — LegacyCoin Explorer</title><style>` + sharedCSS + bookmarkCSS + `</style><script>
async function toggleBookmark(type, value, label) {
  const id = 'bm_' + type + '_' + btoa(value).replace(/[^a-zA-Z0-9]/g,'').substring(0,20);
  const btn = document.getElementById('bm-btn');
  if (btn.classList.contains('saved')) {
    await fetch('/api/bookmarks/delete/' + id, {method:'DELETE'});
    btn.classList.remove('saved');
    btn.innerHTML = '☆ Bookmark';
  } else {
    var rec={id:id,type:type,label:label||''};rec[type==='address'?'address':'txid']=value;
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(rec)});
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
{{if .Block}}<div class="dr"><span class="dk">Block</span><span class="dv gold"><a href="/block/{{.Block.Height}}">{{.Block.Height}}</a></span></div>{{else}}<div class="dr"><span class="dk">Block</span><span class="dv gold">Unconfirmed</span></div>{{end}}
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
<tr class="spacer"><td colspan="4"></td></tr><tr class="spacer"><td colspan="4"></td></tr>
</tbody></table></div>

<div class="pt" style="font-size:16px;margin:18px 0 10px;">Outputs <span>({{len .Tx.Vout}})</span></div>
<div class="tw"><table><thead><tr><th>#</th><th>Value (LBTC)</th><th>Address</th><th>Type</th></tr></thead><tbody>
{{range $out := .Tx.Vout}}<tr>
<td style="color:var(--muted);">{{$out.N}}</td>
<td style="color:var(--gold);font-family:var(--mono);">{{printf "%.8f" $out.Value}}</td>
<td class="hash">{{range $addr := $out.ScriptPubKey.Addresses}}<a href="/address/{{$addr}}">{{$addr}}</a>{{end}}</td>
<td><span class="bge bg-g">{{$out.ScriptPubKey.Type}}</span></td>
</tr>{{end}}
<tr class="spacer"><td colspan="4"></td></tr><tr class="spacer"><td colspan="4"></td></tr>
</tbody></table></div>

<div style="margin-top:18px;display:flex;gap:14px;">
{{if .Block}}<a href="/block/{{.Block.Height}}" style="padding:7px 16px;background:var(--panel);border:1px solid var(--border);font-size:13px;">← Block {{.Block.Height}}</a>{{end}}
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
    var rec={id:id,type:type,label:label||''};rec[type==='address'?'address':'txid']=value;
    await fetch('/api/bookmarks', {method:'POST',headers:{'Content-Type':'application/json'},body:JSON.stringify(rec)});
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
{{range $tx := .Txs}}<tr data-height="{{$tx.Height}}">
<td class="hash"><a href="/tx/{{$tx.Txid}}">{{truncate $tx.Txid 32}}</a></td>
<td><a href="/block/{{$tx.Height}}}}">{{$tx.Height}}</a></td>
<td style="color:var(--muted);font-size:12px;">{{formatTime $tx.Time}}</td>
<td><span class="bge bg-g">{{$tx.Confirmations}}</span></td>
<td>{{len $tx.Vout}}</td>
</tr>{{end}}
<tr class="spacer"><td colspan="5"></td></tr><tr class="spacer"><td colspan="5"></td></tr>
</tbody></table></div></div>` + footSnip + `</body></html>{{end}}

{{define "status"}}<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><link rel="icon" href="data:,"><title>Node Status — LegacyCoin Explorer</title><style>` + sharedCSS + `</style></head><body>` + navSnip + `<div class="c" style="max-width:720px;margin:0 auto;padding-top:60px;">
<div class="pt">Node <span>Status</span></div>
<div id="stat"><div style="color:var(--muted);font-size:13px;">Loading…</div></div>
<div style="text-align:right;color:var(--muted);font-size:12px;margin-top:14px;">Auto-refresh 15s · <a href="/">Open block explorer →</a></div>
</div>
<script>
function fmt(x){if(x===undefined||x===null)return '-';if(x&&typeof x==='object')return JSON.stringify(x);return String(x);}
function loadStatus(){
  var el=document.getElementById('stat');
  fetch('/api/status').then(function(r){return r.json();}).then(function(d){
    if(!d)return;
    var cards=[
      ['Status',d.nodeRunning?'● ONLINE':'● OFFLINE',d.nodeRunning?'#22C55E':'#EF4444'],
      ['Height',fmt(d.height)],
      ['Best Hash',d.bestHash?d.bestHash.substring(0,24)+'…':'-'],
      ['Difficulty',d.difficulty===undefined?'-':(d.difficulty).toFixed(6)],
      ['Peers',fmt(d.connections)],
      ['Mempool',fmt(d.mempool)],
      ['Hash Rate',fmt(d.hashRate)],
      ['Network / Coin',d.network+' / '+d.coin]
    ].map(function(c){
      return '<div class="sc"><div class="sl">'+c[0]+'</div><div class="sv" style="font-size:14px;'+(c[2]?'color:'+c[2]+';':'')+'">'+c[1]+'</div></div>';
    }).join('');
    el.innerHTML='<div class="sg">'+cards+'</div>';
  }).catch(function(e){el.innerHTML='<div style="color:#EF4444;text-align:center;">'+fmt(e.message||e)+'</div>';});
}
loadStatus();
setInterval(loadStatus,15000);
</script></body></html>{{end}}

`
