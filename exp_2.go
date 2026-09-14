// exp_2.go
package explorer

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var exp2Template = template.Must(template.New("exp2").Funcs(template.FuncMap{
	"formatTime":  formatTime,
	"formatLBTC":  formatLBTC,
	"truncate":    truncate,
	"blockReward": blockRewardForHeight,
	"add":         func(a, b int64) int64 { return a + b },
}).Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;}
body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;display:flex;flex-direction:column;}
.pt{font-size:21px;font-weight:700;margin-bottom:14px;flex-shrink:0;padding-top: 10px}
.pt span{color:var(--gold);}
.pt small{font-size:13px;color:var(--muted);font-weight:400;}
.tw{flex:1;min-height:0;overflow:auto;}
.wrap{flex:1;min-height:0;display:flex;}
table{border-collapse:separate;border-spacing:0;}
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
.pg{display:flex;flex-direction:column;justify-content:center;gap:8px;margin-left:10px;flex-shrink:0;}
.pg a,.pg span{display:flex;align-items:center;justify-content:center;width:38px;height:38px;background:var(--panel);border:1px solid var(--border);font-size:16px;color:var(--muted);text-decoration:none;cursor:pointer;}
.pg a:hover{border-color:var(--gold);color:var(--gold);}
.pg .pi{width:auto;height:auto;padding:4px 0;background:none;border:none;font-family:var(--mono);font-size:11px;color:var(--muted);cursor:default;}
.sel{background:var(--panel2);box-shadow:inset 3px 0 0 #3B82F6;}
.tw.hdcur thead tr{box-shadow:inset 3px 0 0 #3B82F6;}
.curl{box-shadow:inset 3px 0 0 #3B82F6;}
</style>
<div class=pt><span style=color:#555>2.</span> All <span>Blocks</span> <small>Tip: {{.Tip}}</small></div>
{{if .Error}}<div class="offline">⚠ {{.Error}}</div>{{end}}
<div class=wrap>
<div class=tw>
<table>
<thead>
<tr>
  <th id="curhdr">Height</th>
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
<tr data-height="{{.Height}}" data-hash="{{.Hash}}">
  <td onclick="parent.lc('/exp-block/{{.Height}}')">{{.Height}}</td>
  <td class="hash"><a href="/exp-block/{{.Hash}}" onclick="parent.lc('/exp-block/{{.Hash}}');return false;">{{truncate .Hash 40}}</a></td>
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
{{if .HasPrev}}<a data-dir="prev" href="/exp-2?page={{.PrevPage}}" title="Newer ▴">▲</a>{{else}}<span data-dir="prev" title="Newer ▴">▲</span>{{end}}
<span class="pi">{{if gt .TotalPages 0}}{{add .Page 1}}/{{.TotalPages}}{{end}}</span>
{{if .HasNext}}<a data-dir="next" href="/exp-2?page={{.NextPage}}" title="Older ▾">▼</a>{{else}}<span data-dir="next" title="Older ▾">▼</span>{{end}}
</div>
</div>
<script>
function lcRows(){return document.querySelectorAll('tbody tr:not(.spacer)');}
function lcSelRow(){return document.querySelector('tbody tr.sel');}
function lcScroll(r){
  var tw=r.closest('.tw');if(!tw)return;
  var th=tw.querySelector('thead');var oh=th?th.offsetHeight+2:0;
  var rowH=r.offsetHeight||36;
  var padTop=rowH; if(oh>0)padTop=rowH*2;
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
  var tr=e.target&&e.target.closest?e.target.closest('tbody tr:not(.spacer)'):null;
  if(tr)lcSelect(tr);
});
document.addEventListener('keydown',function(e){
  if(e.target.tagName==='INPUT'||e.target.tagName==='TEXTAREA'){
    if(e.key==='Escape'||e.keyCode===27){e.preventDefault();if(e.target.value){e.target.value='';}else e.target.blur();}
    return;
  }
  var k=e.key
  if(k==='1'         ){e.preventDefault();parent.lf('/exp-1');return}
  if(k==='2'         ){e.preventDefault();parent.lf('/exp-2');return}
  if(k==='3'         ){e.preventDefault();parent.lf('/exp-3');return}
  if(k==='4'         ){e.preventDefault();top.location.href='/home';return}
  if(k==='Backspace' ){try{var fi=parent.document.querySelector('.sf input');if(fi){e.preventDefault();fi.focus();}}catch(x){}return;}
  if(k==='Enter'     ){var sl=document.querySelector('tbody tr.sel[data-height]');if(sl){e.preventDefault();parent.lc('/exp-block/'+sl.getAttribute('data-height'));}return;}
  if(k==='Escape'    ){parent.clkHH();return}
  if(k==='ArrowLeft' ){e.preventDefault();parent.setBL( );return}
  if(k==='ArrowRight'){e.preventDefault();parent.setBR( );return}
  if(k==='ArrowDown' ){e.preventDefault();lcNav(1,false );return}
  if(k==='ArrowUp'   ){e.preventDefault();lcNav(-1,false);return}
  if(k==='PageDown'  ){e.preventDefault();var l=document.querySelector('[data-dir="next"]');if(l&&l.href)location.href=l.href;return;}
  if(k==='PageUp'    ){e.preventDefault();var l=document.querySelector('[data-dir="prev"]');if(l&&l.href)location.href=l.href;return;}
  if(k==='Home'      ){e.preventDefault();lcEnd(false);return;}
  if(k==='End'       ){e.preventDefault();lcEnd(true);return;}
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

// exp2Page — последняя успешно собранная страница All Blocks
// (stale-fallback на случай, когда узел не отвечает / под rate-limit).
type exp2Page struct {
	Tip    int64
	Blocks []*Block
}

func (s *Server) getExp2Fallback(page int64) (exp2Page, bool) {
	s.exp2Mu.Lock()
	defer s.exp2Mu.Unlock()
	fb, ok := s.exp2Fallback[page]
	return fb, ok
}

func (s *Server) setExp2Fallback(page int64, tip int64, blocks []*Block) {
	s.exp2Mu.Lock()
	defer s.exp2Mu.Unlock()
	if s.exp2Fallback == nil {
		s.exp2Fallback = map[int64]exp2Page{}
	}
	s.exp2Fallback[page] = exp2Page{Tip: tip, Blocks: blocks}
}

func (s *Server) handleExp2(w http.ResponseWriter, r *http.Request) {
	const perPage = 50

	data := struct {
		Tip        int64
		Blocks     []*Block
		Page       int64
		TotalPages int64
		HasPrev    bool
		HasNext    bool
		PrevPage   int64
		NextPage   int64
		Error      string
	}{}

	// Страница (0 — самая свежая)
	page := int64(0)
	if p := r.URL.Query().Get("page"); p != "" {
		if v, err := strconv.ParseInt(p, 10, 64); err == nil && v >= 0 {
			page = v
		}
	}
	data.Page = page

	// Тёплая первая страница — без единого RPC-вызова (узел под rate-limit):
	// tip берётся из свежайшего блока кэша.
	if page == 0 {
		if blocks, e := s.cachedRecentBlocks(perPage); e == nil && len(blocks) > 0 {
			data.Tip = blocks[0].Height
			data.TotalPages = (data.Tip + perPage - 1) / perPage
			data.Blocks = blocks
			s.setExp2Fallback(0, data.Tip, data.Blocks)
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			if err := exp2Template.Execute(w, data); err != nil {
				log.Printf("handleExp2 template error: %v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			return
		}
	}

	tip, err := s.rpc.GetBlockCount()

	// Узел недоступен — отдаём последнюю успешно собранную страницу.
	if err != nil {
		if fb, ok := s.getExp2Fallback(page); ok {
			data.Tip = fb.Tip
			data.TotalPages = (fb.Tip + perPage - 1) / perPage
			data.Blocks = fb.Blocks
		} else {
			data.Error = "Cannot connect to legacycoind node"
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			exp2Template.Execute(w, data)
			return
		}
	} else {
		data.Tip = tip
		data.TotalPages = (tip + perPage - 1) / perPage
		if page == 0 {
			// Первая страница — через кэш/запасной набор (устойчиво к сбоям,
			// значительно снижает нагрузку на rate-limited узел).
			if blocks, e := s.cachedRecentBlocks(perPage); e == nil {
				data.Blocks = blocks
			}
		} else {
			data.Blocks = s.fetchExp2Page(tip, page, perPage)
		}
		if len(data.Blocks) == 0 {
			if fb, ok := s.getExp2Fallback(page); ok {
				data.Tip = fb.Tip
				data.TotalPages = (fb.Tip + perPage - 1) / perPage
				data.Blocks = fb.Blocks
			}
		}
		if len(data.Blocks) == 0 {
			data.Error = "Cannot connect to legacycoind node"
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			exp2Template.Execute(w, data)
			return
		}
		s.setExp2Fallback(page, data.Tip, data.Blocks)
	}

	start := data.Tip - page*perPage
	if start < 0 {
		start = 0
	}
	end := start - perPage + 1
	if end < 0 {
		end = 0
	}

	data.HasPrev = page > 0
	data.PrevPage = page - 1
	data.HasNext = end > 0
	data.NextPage = page + 1

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	exp2Template.Execute(w, data)
}

// fetchExp2Page собирает блоки страницы напрямую с узла.
func (s *Server) fetchExp2Page(tip, page, perPage int64) []*Block {
	start := tip - page*perPage
	if start < 0 {
		start = 0
	}
	end := start - perPage + 1
	if end < 0 {
		end = 0
	}
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
	return blocks
}
