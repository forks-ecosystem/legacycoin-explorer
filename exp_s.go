// exp_s.go — блок поиска в отдельном iframe (#DS, src=/exp-s).
// Поиск вынесен в свой фрейм, чтобы главная страница оставалась без <input>
// и визирь (круг-кнопка в #DD) мог жить в основном блоке без конфликта фокуса.
// Форма целится в родительский фрейм карточки fC (target="fC").
package explorer

import (
    "html/template"
    "net/http"
)

var expSTemplate = template.Must(template.New("expS").Parse(`<style>
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;background:transparent;}
body{display:flex;align-items:center;}
.sf{display:flex;width:100%;}
.sf input{
  flex:1;min-width:0;background:#040404;border:1px solid #222;border-right:none;
  color:#E8E8E8;padding:8px 12px;font-size:13px;font-family:'Courier New',monospace;outline:none;
}
.sf input:focus{border-color:#D4A017;}
.sf button{background:#D4A017;color:#080808;border:none;padding:8px 14px;cursor:pointer;font-size:13px;font-weight:700;}
</style>
<form target=fC class="sf" action="/searchx" method="GET" autocomplete="off">
  <input  type="text" id="q" name="q" placeholder="Height, hash, txid or address…" onfocus="parent.setBB()">
  <button type="submit">→</button>
</form>
<script>
function focusSearch(){ var i=document.getElementById('q'); if(i) i.focus(); }
document.getElementById('q').addEventListener('keydown',function(e){
  if(e.key==='Escape'){
    e.preventDefault();
    var i=document.getElementById('q');
    if(i){ if(i.value){ i.value=''; } else { i.blur(); } }
  }
});
document.addEventListener('DOMContentLoaded',function(){
  var frm=document.querySelector('form');
  frm.addEventListener('submit',function(e){
    e.preventDefault();
    var i=document.getElementById('q');
    var q=(i.value||'').trim();
    if(!q) return;
    var top=window.top;
    if(top && top.document){
      var c=top.document.getElementById('fC');
      if(c) c.src='/searchx?q='+encodeURIComponent(q);
    }
  });
});
</script>
`))

func (s *Server) handleExpS(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := expSTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
