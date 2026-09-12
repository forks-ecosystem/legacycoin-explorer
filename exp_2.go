// exp_2.go
package explorer

import (
	"html/template"
	"net/http"
	"strconv"
)

var exp2Template = template.Must(template.New("exp2").Funcs(template.FuncMap{
	"formatTime":  formatTime,
	"formatLBTC":  formatLBTC,
	"truncate":    truncate,
	"blockReward": blockRewardForHeight,
}).Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;}
body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;display:flex;flex-direction:column;}
.pt{font-size:21px;font-weight:700;margin-bottom:14px;flex-shrink:0;}
.pt span{color:var(--gold);}
.pt small{font-size:13px;color:var(--muted);font-weight:400;}
.tw{flex:1;min-height:0;overflow:auto;border-top:1px solid var(--border);}
table{width:100%;border-collapse:separate;border-spacing:0;}
thead th{position:sticky;top:0;z-index:5;background:var(--panel2);color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:10px 13px;text-align:left;border-bottom:1px solid var(--border);white-space:nowrap;}
tbody tr{border-bottom:1px solid var(--border);}
tbody tr:hover{background:var(--panel);}
tbody td{padding:10px 13px;font-size:13px;vertical-align:middle;}
tr.spacer{height:16px;}
.hash{font-family:var(--mono);font-size:12px;}
a{color:var(--gold);text-decoration:none;}
a:hover{text-decoration:underline;}
.bge{display:inline-block;padding:2px 7px;font-size:11px;font-weight:600;}
.bg-g{background:rgba(34,197,94,.12);color:var(--green);border:1px solid rgba(34,197,94,.25);}
.offline{background:rgba(239,68,68,.1);border:1px solid rgba(239,68,68,.3);color:var(--red);padding:13px 17px;margin-bottom:18px;font-size:14px;}
.pg{display:flex;gap:8px;margin-top:10px;flex-shrink:0;}
.pg a,.pg span{padding:6px 13px;background:var(--panel);border:1px solid var(--border);font-size:13px;color:var(--muted);}
.pg a:hover{border-color:var(--gold);color:var(--gold);text-decoration:none;}
.pg .cur{border-color:var(--gold);color:var(--gold);}
.sel{background:var(--panel2);box-shadow:inset 3px 0 0 #3B82F6;}
.tw.hdcur thead tr{box-shadow:inset 3px 0 0 #3B82F6;}
.lcCard{position:fixed;top:8px;right:8px;bottom:8px;width:340px;min-width:160px;max-width:94%;background:var(--dark);border:1px solid var(--gold);border-radius:8px;z-index:120;display:none;flex-direction:column;box-shadow:0 10px 40px rgba(0,0,0,.65);}
.lcHead{display:flex;align-items:center;justify-content:space-between;gap:10px;padding:11px 14px;border-bottom:1px solid var(--border);color:var(--gold);font-weight:700;font-size:12px;text-transform:uppercase;letter-spacing:1px;flex-shrink:0;}
.lcX{cursor:pointer;color:var(--muted);font-size:15px;padding:0 4px;}.lcX:hover{color:var(--red);}
.lcBody{padding:10px 14px;overflow-y:auto;flex:1;min-height:0;font-size:12px;}
.lcCard .ld{display:flex;justify-content:space-between;gap:12px;padding:7px 0;border-bottom:1px solid rgba(255,255,255,.05);}.lcCard .ld span{color:var(--muted);flex-shrink:0;}.lcCard .ld b{font-family:var(--mono);text-align:right;word-break:break-all;}
</style>

<div class="pt">All <span>Blocks</span> <small>Tip: {{.Tip}}</small></div>
{{if .Error}}<div class="offline">⚠ {{.Error}}</div>{{end}}

<div class="tw">
<table>
<thead>
<tr>
  <th>Height</th>
  <th>Hash</th>
  <th>Time (UTC)</th>
  <th>Txs</th>
  <th>Reward</th>
  <th>Size</th>
  <th>Confs</th>
</tr>
</thead>
<tbody>
{{range .Blocks}}
<tr data-height="{{.Height}}">
  <td><a href="/block/{{.Height}}" target="_top">{{.Height}}</a></td>
  <td class="hash"><a href="/block/{{.Hash}}" target="_top">{{truncate .Hash 40}}</a></td>
  <td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
  <td>{{len .Tx}}</td>
  <td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
  <td style="color:var(--muted);">{{.Size}} B</td>
  <td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>
{{else}}
<tr><td colspan="7" style="text-align:center;color:var(--muted);padding:26px;">No blocks.</td></tr>
{{end}}
<tr class="spacer"><td colspan="7"></td></tr><tr class="spacer"><td colspan="7"></td></tr>
</tbody>
</table>
</div>

<div class="pg">
{{if .HasPrev}}<a data-dir="prev" href="/exp-2?page={{.PrevPage}}">← Newer</a>{{else}}<span data-dir="prev">← Newer</span>{{end}}
<span class="cur">Page {{.Page}}</span>
{{if .HasNext}}<a data-dir="next" href="/exp-2?page={{.NextPage}}">Older →</a>{{else}}<span data-dir="next">Older →</span>{{end}}
</div>
<script>
var lcCard=null;
function lcInit(){if(lcCard)return;var d=document.createElement('div');d.className='lcCard';d.innerHTML='<div class="lcHead"><span>Block card</span><span class="lcX" onclick="lcClose()">✕</span></div><div class="lcBody"></div>';document.body.appendChild(d);lcCard=d;}
function lcIsOpen(){return lcCard&&lcCard.style.display!=='none';}
function lcClose(){if(lcCard)lcCard.style.display='none';}
var lcSeq=0;
function lcLoad(h,retry){
  lcInit();
  var seq=++lcSeq;
  var d=lcCard,body=d.querySelector('.lcBody');
  fetch('/api/block/'+h).then(function(r){return r.json();}).then(function(b){
    if(seq!==lcSeq)return;
    if(!b||b.error){if(retry<2){setTimeout(function(){lcLoad(h,retry+1);},350);return;}body.innerHTML='<div style="color:#EF4444;padding:8px;">node busy — press Enter to retry</div>';return;}
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
      +'<div style="margin-top:12px"><a href="/block/'+b.height+'" target="_top">Open block page →</a></div>';
  }).catch(function(){if(seq===lcSeq&&retry<2)setTimeout(function(){lcLoad(h,retry+1);},350);});
}
function lcOpen(h){lcInit();lcCard.querySelector('.lcHead span').textContent='Block #'+h;lcCard.querySelector('.lcBody').innerHTML='<div style="color:var(--muted);padding:8px;">Loading…</div>';lcCard.style.display='flex';lcLoad(h,0);}
function lcRows(){return document.querySelectorAll('tbody tr:not(.spacer)');}
function lcSelRow(){return document.querySelector('tbody tr.sel');}
function lcScroll(r){
  var tw=r.closest('.tw');if(!tw)return;
  var th=tw.querySelector('thead');var oh=th?th.offsetHeight+2:0;
  var rowH=r.offsetHeight||36,pad=rowH*1.6;
  var y=0,el=r;while(el&&el!==tw){y+=el.offsetTop;el=el.offsetParent;if(!el)break;}
  var twH=tw.clientHeight;
  if(y<tw.scrollTop+oh+pad)tw.scrollTop=y-oh-pad;
  else if(y+rowH>tw.scrollTop+twH-pad)tw.scrollTop=y+rowH-twH+pad;
}
function lcSelect(r){
  var old=lcSelRow();if(old&&old!==r)old.classList.remove('sel');
  r.classList.add('sel');lcScroll(r);
  var t=r.closest('.tw');if(t)t.classList.remove('hdcur');
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
    if(e.key==='Escape'||e.keyCode===27){e.preventDefault();if(e.target.value){e.target.value='';}else e.target.blur();}
    return;
  }
  var k=e.key,c=e.keyCode;
  if(k==='1'||c===49||c===97){try{parent.document.getElementById('fB').src='/exp-1';}catch(x){}return;}
  if(k==='2'||c===50||c===98){try{parent.document.getElementById('fB').src='/exp-2';}catch(x){}return;}
  if(k==='Backspace'||c===8){try{var fi=parent.document.querySelector('.sf input');if(fi){e.preventDefault();fi.focus();}}catch(x){}return;}
  if(k==='Enter'||c===13){var sl=document.querySelector('tbody tr.sel[data-height]');if(sl){e.preventDefault();lcOpen(sl.getAttribute('data-height'));}return;}
  if(k==='Escape'||c===27){if(lcIsOpen()){lcClose();return;}return;}
  if(k==='ArrowDown'||c===40){e.preventDefault();lcNav(1,false);return;}
  if(k==='ArrowUp'||c===38){e.preventDefault();lcNav(-1,false);return;}
  if(k==='PageDown'||c===34){e.preventDefault();lcNav(1,true);return;}
  if(k==='PageUp'||c===33){e.preventDefault();lcNav(-1,true);return;}
  if(k==='Home'||c===36){e.preventDefault();lcEnd(false);return;}
  if(k==='End'||c===35){e.preventDefault();lcEnd(true);return;}
  if(k==='ArrowLeft'||c===37){var l=document.querySelector('[data-dir="prev"]');if(l&&l.href){e.preventDefault();location.href=l.href;}return;}
  if(k==='ArrowRight'||c===39){var l=document.querySelector('[data-dir="next"]');if(l&&l.href){e.preventDefault();location.href=l.href;}return;}
});
</script>`))

func (s *Server) handleExp2(w http.ResponseWriter, r *http.Request) {
	const perPage = 50

	data := struct {
		Tip      int64
		Blocks   []*Block
		Page     int64
		HasPrev  bool
		HasNext  bool
		PrevPage int64
		NextPage int64
		Error    string
	}{}

	if !s.rpc.Ping() {
		data.Error = "Cannot connect to legacycoind node"
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		exp2Template.Execute(w, data)
		return
	}

	tip, err := s.rpc.GetBlockCount()
	if err != nil {
		data.Error = err.Error()
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		exp2Template.Execute(w, data)
		return
	}
	data.Tip = tip

	// Страница (0 — самая свежая)
	page := int64(0)
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.ParseInt(p, 10, 64); err == nil && v >= 0 {
			page = v
		}
	}
	data.Page = page

	start := tip - page*perPage
	if start < 0 {
		start = 0
	}
	end := start - perPage + 1
	if end < 0 {
		end = 0
	}

	// Собираем блоки: hash → block (с заполнением confirmations)
	var blocks []*Block
	for h := start; h >= end; h-- {
		hash, err := s.rpc.GetBlockHash(h)
		if err != nil {
			continue
		}
		b, err := s.rpc.GetBlock(hash)
		if err != nil {
			continue
		}
		b.Confirmations = tip - h + 1
		blocks = append(blocks, b)
	}
	data.Blocks = blocks

	data.HasPrev = page > 0
	data.PrevPage = page - 1
	data.HasNext = end > 0
	data.NextPage = page + 1

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	exp2Template.Execute(w, data)
}
