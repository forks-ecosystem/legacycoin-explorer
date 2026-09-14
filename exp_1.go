// exp_1.go
package explorer

import (
    "html/template"
    "log"
    "net/http"
)

var exp1Template = template.Must(template.New("exp1").Funcs(template.FuncMap{
    "formatTime":  formatTime,
    "formatLBTC":  formatLBTC,
    "truncate":    truncate,
    "blockReward": blockRewardForHeight,
}).Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;}
body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;display:flex;flex-direction:column;}
.pt{font-size:21px;font-weight:700;margin-bottom:14px;flex-shrink:0;padding-top:10px}
.pt span{color:var(--gold);}
.tw{flex:1;min-height:0;overflow:auto;}
table{border-collapse:separate;border-spacing:0;}
thead th{position:sticky;top:0;z-index:5;background:linear-gradient(to right,var(--panel2),#000);color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:10px 13px;text-align:left;border-bottom:1px solid var(--border);white-space:nowrap;}
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
.sel{background:var(--panel2);box-shadow:inset 3px 0 0 #3B82F6;}
.tw.hdcur thead tr{box-shadow:inset 3px 0 0 #3B82F6;}
.curl{box-shadow:inset 3px 0 0 #3B82F6;}
</style>
<div class=pt><span style=color:#555>1.</span> Latest <span>Blocks</span></div>
{{if .Error}}<div class="offline">⚠ {{.Error}}</div>{{end}}
<div class="tw">
<table>
<thead>
<tr>
  <th id="curhdr">
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
{{range .RecentBlocks}}
<tr data-height="{{.Height}}" data-hash="{{.Hash}}">
  <td>
  <td              onclick="parent.lc('/exp-block/{{.Height}}')">{{.Height}}</td>
  <td class="hash" onclick="parent.lc('/exp-block/{{.Hash  }}')">{{truncate .Hash 32}}</td>
  <td style="color:var(--muted);font-size:12px;">{{formatTime .Time}}</td>
  <td>{{len .Tx}}</td>
  <td style="color:var(--gold);font-family:var(--mono);font-size:12px;">{{formatLBTC (blockReward .Height)}}</td>
  <td style="color:var(--muted);">{{.Size}} B</td>
  <td><span class="bge bg-g">{{.Confirmations}}</span></td>
</tr>
{{else}}
<tr><td colspan="8" style="text-align:center;color:var(--muted);padding:26px;">No blocks yet.</td></tr>
{{end}}
<tr class=spacer><td colspan="8"></td></tr><tr class="spacer"><td colspan="8"></td></tr>
<tr><td><tr><td></tbody>
<thead><tr><td><th>END</div></thead>
</table>
</div>
<script>
function lcRows(){return document.querySelectorAll('tbody tr[data-height]');}
function lcSelRow(){return document.querySelector('tbody tr.sel');}
function lcScroll(r){
  var tw=r.closest('.tw');if(!tw)return;
  var th=tw.querySelector('thead');var oh=th?th.offsetHeight+2:0;
  var rowH=r.offsetHeight||36;
  var padTop=rowH; if(oh>0)padTop=rowH*2+10;
  var padBot=rowH*1.6;
  var y=0,el=r;while(el&&el!==tw){y+=el.offsetTop;el=el.offsetParent;if(!el)break;}
  var twH=tw.clientHeight;
  if(y<tw.scrollTop+oh+padTop)tw.scrollTop=y-oh-padTop;
  else if(y+rowH>tw.scrollTop+twH-padBot)tw.scrollTop=y+rowH-twH+padBot;
}
function lcSelect(r){
  var old=lcSelRow();if(old&&old!==r)old.classList.remove('sel');
  r.classList.add('sel');lcScroll(r);
  var c=document.getElementById('curhdr');if(c)c.classList.remove('curl');
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
(function(){var c=document.getElementById('curhdr');if(c)c.classList.add('curl');})();
document.addEventListener('click',function(e){
  var c=document.getElementById('curhdr');if(c)c.classList.remove('curl');
  var t=document.querySelector('.tw');if(t)t.classList.remove('hdcur');
  var tr=e.target&&e.target.closest?e.target.closest('tbody tr[data-height]'):null;
  if(tr)lcSelect(tr);
});
document.addEventListener('keydown',function(e){
  var k=e.key
  if(e.target.tagName==='INPUT'||e.target.tagName==='TEXTAREA'){
    if(k==='Escape'){e.preventDefault();if(e.target.value){e.target.value='';}else e.target.blur();}
    if(k==='ArrowLeft'  && e.ctrlKey){e.preventDefault();setBL();return}
    if(k==='ArrowRight' && e.ctrlKey){e.preventDefault();setBR();return}
    return;
  }
  if(k==='Enter'     ){var sl=document.querySelector('tbody tr.sel[data-height]');if(sl){
                       var sd='/exp-block/'+sl.getAttribute('data-height')
                       e.preventDefault();parent.lc(     sd)};return}
  if(k==='1'         ){e.preventDefault();parent.lf('/exp-1');return}
  if(k==='2'         ){e.preventDefault();parent.lf('/exp-2');return}
  if(k==='3'         ){e.preventDefault();parent.lf('/exp-3');return}
  if(k==='4'         ){e.preventDefault();top.location.href='/home';return}
  if(k==='Backspace' ){e.preventDefault();parent.setSF(     );return}
  if(k==='Escape'    ){e.preventDefault();parent.clkHH(     );return}
  if(k==='ArrowLeft' ){e.preventDefault();parent.setBL(     );return}
  if(k==='ArrowRight'){e.preventDefault();parent.setBR(     );return}
  if(k==='ArrowDown' ){e.preventDefault();lcNav( 1,false    );return}
  if(k==='ArrowUp'   ){e.preventDefault();lcNav(-1,false    );return}
  if(k==='PageDown'  ){e.preventDefault();lcNav( 1,true     );return}
  if(k==='PageUp'    ){e.preventDefault();lcNav(-1,true     );return}
  if(k==='Home'      ){e.preventDefault();lcEnd(false       );return}
  if(k==='End'       ){e.preventDefault();lcEnd(true        );return}
});
(function(){
  if(document.querySelector('.offline')||!document.querySelector('tbody tr[data-height]')){
    var key=location.pathname.replace(/[^a-z0-9]+/gi,'-')+':retry';
    var n=(parseInt(localStorage.getItem(key)||'0',10)||0)+1;
    if(n<=4){localStorage.setItem(key,n);setTimeout(function(){location.reload();},n*4000);}
    else {localStorage.removeItem(key);}
  }else{
    try{localStorage.removeItem(location.pathname.replace(/[^a-z0-9]+/gi,'-')+':retry');}catch(x){}
  }
})();
window.focus();
document.body.tabIndex=-1;
document.body.focus();
</script>`))

func (s *Server) handleExp1(w http.ResponseWriter, r *http.Request) {
	data := struct {
		RecentBlocks []*Block
		Error        string
	}{}
	if blocks, err := s.cachedRecentBlocks(20); err == nil {
		data.RecentBlocks = blocks
	} else {
		data.Error = "Cannot connect to legacycoind node"
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := exp1Template.Execute(w, data); err != nil {
		log.Printf("handleExp1 template error: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
