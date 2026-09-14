// ex.go
package explorer

import (
	"html/template"
	"net/http"
)

var exTemplate = template.Must(template.New("ex").Parse(`
<!DOCTYPE html>
<html lang=ru>
<head>
<meta charset=UTF-8>
<meta name=viewport content="width=device-width,initial-scale=1.0">
<link rel=icon href=/favicon.ico>
<title>{{.Title}}</title>
<style>
:root {
  --gold:#D4A017;
  --black:#080808;
  --dark:#0F0F0F;
  --panel:#040404;
  --border:#222;
  --text:#E8E8E8;
  --muted:#888;
  --mono:'Courier New', monospace;
}
*{box-sizing:border-box; }
body{ margin:0;  padding:0;  font-family:Arial, sans-serif;  background:#010101;  color:var(--text);  overflow:hidden;}
nav {  background:var(--dark);  padding:0 40px;  display:flex;  align-items:center;  gap:28px;
  height:54px;  position:sticky;  top:0;  z-index:100;
}
#TL, #TR{position:fixed; top:14px; left:3px; cursor:pointer; z-index:150; }
.brand{ font-size:17px;  font-weight:700;  color:var(--gold);  letter-spacing:1px;  white-space:nowrap;  text-decoration:none;}
.brand small{ font-size:11px;  font-weight:400;  color:var(--muted);  margin-left:5px;}
.navl{ display:flex;  gap:20px;  flex:1;}
.navl a{ color:var(--muted);  font-size:13px;  text-decoration:none;}
.navl a:hover{color:var(--gold); }
.sf{ display:flex;  margin-left:auto;}
.sf input{ background:var(--panel);  border:1px solid var(--border);  border-right:none;  color:var(--text);
  padding:6px 12px;  font-size:13px;  width:280px;  font-family:var(--mono);  outline:none;
}
.sf input:focus{border-color:var(--gold); }
.sf button {
  background:var(--gold);  color:var(--black);  border:none;  padding:6px 14px;  cursor:pointer;
  font-size:13px;  font-weight:700;
}
#DA{position:fixed; inset:0}
#DB{position:fixed; top:56px; right:0; bottom:40px; left:220px}
#DC{position:fixed; top:56px; right:0; bottom:40px; left:100%; background:linear-gradient(to right, #181818, #111 10%); }
#DD{position:fixed; bottom:33px; left:0; right:0; height:18px; z-index:1000; cursor:ew-resize; }
#KN {
  position:absolute; top:50%; left:0; width:18px; height:18px; margin-top:-9px; border-radius:50%;
  background:radial-gradient(circle at 35% 30%, #7fc2ff, #2a76d4 55%, #14335e);
  box-shadow:0 0 12px rgba(88,166,255,.75), 0 0 3px rgba(0,0,0,.9), inset 0 0 5px rgba(255,255,255,.35);
  border:1px solid #9cc6ff; cursor:ew-resize;
}
#DS{position:fixed; top:4px; right:40px; width:330px; height:50px; z-index:120; }
.FA, .FC, .FS{height:100%; width:100%; border:none; }
.CP{position:absolute; bottom:3px; left:0; right:0; display:flex; align-items:center; padding:0 8px; }
footer {
  position:fixed;  bottom:0;  left:220px;  right:0;  background:linear-gradient(to right, #000, #333, #000);
  padding:13px 28px;  text-align:center;  font-size:12px;  color:var(--muted);
}
footer span{color:var(--gold); }
</style>
</head>
<body>
<nav>
  <a href=/ class=brand>⛓ LegacyCoin <small>EXPLORER</small></a>
  <div class=navl>
    <a href=/exp-1 target=fB>1. Latest</a>
    <a href=/exp-2 target=fB>2. All Blocks</a>
    <a href=/exp-3 target=fB>3. API</a>
    <a href=/home  target=_top>4. Home</a>
  </div>
</nav>
<div id=DA><iframe name=fA id=fA class=FA src=/exp-a></iframe></div>
<div id=DB><iframe name=fB id=fB class=FA src=/exp-1></iframe></div>
<div id=DC><iframe name=fC id=fC class=FC            ></iframe></div>
<div id=DD><div            id=KN></div></div>
<div id=DS><iframe name=fS id=fS class=FS src=/exp-s></iframe></div>

<a href=/ title="Home"><img src=/lbtc.png alt="LBTC" style=position:fixed;top:60px;height:50px></a>
<img id=TL src=/left.png  alt="menu" title="Меню" onclick=clkHH() style=height:27px;display:inline>
<img id=TR src=/right.png alt="menu" title="Меню" onclick=clkHH() style=height:27px;display:none>
<footer><span>LegacyCoin (LBTC)</span> Block Explorer · CPU money for everyone</footer>
<script>
const DB=document.getElementById('DB')
const DC=document.getElementById('DC')
const KN=document.getElementById('KN')
const fB=document.getElementById('fB')
const fC=document.getElementById('fC')
let _v=50, _drag=false, _x='100%', _z='220px', _n='46px'
function lim(  n){ n=Math.round(n); return Math.max(0,Math.min(100,n))}
function place(v){ var W=document.documentElement.clientWidth, d=KN.offsetWidth; KN.style.left=Math.round((v/100)*(W-d))+'px'; }
function getPR( ){ return _v; }
function setPR(n){ _v=lim(n); place(_v); }
function fromX(x){ var W=document.documentElement.clientWidth, d=KN.offsetWidth; return lim((x-d/2)/(W-d)*100); }
DB.style.left=_z;
DC.style.left=_x;
function setCard(val) {
  var n=parseFloat(val);
  const k=Math.max(0, Math.min(100,isFinite(n)?n:50));
  DC.style.left=k+'%';
  localStorage.setItem('panelPos',k);
  setPR(k);
}
function setBB(){setCard(_v  )}
function setBL(){setCard(_v-1)}
function setBR(){setCard(_v+1)}
function lf(url){fB.src=url;DC.style.left='100%'}
function lc(url){fC.src=url;setCard(_v)}
function clkHH(){
  if(  DC.style.left!==_x){DC.style.left=_x}else{
    if(DB.style.left===_z){DB.style.left=_n}else{
       DB.style.left  =_z
    }
  }
  setMenu(DB.style.left===_z);
}
function setMenu(open){
  var tl=document.getElementById('TL');
  var tr=document.getElementById('TR');
  if(tl) tl.style.display=open?'inline':'none';
  if(tr) tr.style.display=open?'none':'inline';
}
function setSF() {
  var S=document.getElementById('fS');
  if(S && S.contentWindow && S.contentWindow.focusSearch) S.contentWindow.focusSearch();
}
(function(){
  var dd=document.getElementById('DD');
  dd.addEventListener('mousedown', function(e){
    _drag=true; setPR(fromX(e.clientX)); setCard(_v);
    if(e.preventDefault) e.preventDefault();
  });
  window.addEventListener('mousemove', function(e){
    if(!_drag) return; setPR(fromX(e.clientX)); setCard(_v);
  });
  window.addEventListener('mouseup',   function(){ _drag=false; });
})();
document.addEventListener('keydown', function(e) {
  const k=e.key;
  const isInput=e.target.tagName==='INPUT' || e.target.tagName==='TEXTAREA';
  if(isInput){
    if(k==='Escape'                 ){e.preventDefault(); if(e.target.value) e.target.value='';  else e.target.blur(); }
    if(k==='ArrowLeft'  && e.ctrlKey){ e.preventDefault(); setBL()}
    if(k==='ArrowRight' && e.ctrlKey){ e.preventDefault(); setBR()}
    return;
  }
  if(k==='ArrowLeft') {e.preventDefault(); setBL(); return}
  if(k==='ArrowRight'){e.preventDefault(); setBR(); return}
  if(k==='Escape')    {e.preventDefault(); clkHH(); return}
  if(k==='Backspace') {e.preventDefault(); setSF(); return}
  if(k==='1')         {e.preventDefault(); lf('/exp-1'); return}
  if(k==='2')         {e.preventDefault(); lf('/exp-2'); return}
  if(k==='3')         {e.preventDefault(); lf('/exp-3'); return}
  if(k==='4')         {top.location.href='/home'; return}
});
document.addEventListener('DOMContentLoaded',function(){
  try{ var s=parseInt(localStorage.getItem('panelPos'),10); if(!isNaN(s)) _v=lim(s); }catch(e){}
  DC.style.left='100%';   // карточка всегда закрыта при загрузке
  setPR(_v);
  setMenu(true);
});
</script>
</body>
</html>
`))

func (s *Server) handleHomeEx(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	data := struct {
		Title string
	}{Title: "LBTC Explorer"}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := exTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
