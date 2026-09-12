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
<meta name="viewport" content="width=device-width,initial-scale=1.0">
<title>{{.Title}}</title>
<style>
* { box-sizing: border-box; }
body{ margin: 0; padding: 0; font-family: 'Arial', sans-serif; background-color: #010101; color: #e6e9ef; overflow: hidden;}
nav{  background:var(--dark);  border-bottom:1px solid var(--border);  padding:0 40px;  display:flex;  align-items:center;  gap:28px;  height:54px;  position:sticky;  top:0;  z-index:100;}
.brand{  font-size:17px;  font-weight:700;  color:var(--gold);  letter-spacing:1px;  white-space:nowrap;  text-decoration:none;}
.brand small{  font-size:11px;  font-weight:400;  color:var(--muted);  margin-left:5px;}
.navl{  display:flex;  gap:20px;  flex:1;}
.navl a{  color:var(--muted);  font-size:13px;  text-decoration:none;}
.navl a:hover{  color:var(--gold);  text-decoration:none;}
.sf{  display:flex;  margin-left:auto;}
.sf input{  background:var(--panel);  border:1px solid var(--border);  border-right:none;  color:var(--text);  padding:6px 12px;
  font-size:13px;  width:280px;  font-family:var(--mono);  outline:none;
}
.sf input:focus{  border-color:var(--gold);}
.sf button{  background:var(--gold);  color:var(--black);  border:none;  padding:6px 14px;
  cursor:pointer;  font-size:13px;  font-weight:700;
}
.brand{font-size:17px;font-weight:700;color:var(--gold);letter-spacing:1px;white-space:nowrap;}.brand.small{font-size:11px;font-weight:400;color:var(--muted);margin-left:5px;}
.navl{display:flex;gap:20px;flex:1;}.navl a{color:var(--muted);font-size:13px;}.navl a:hover{color:var(--gold);text-decoration:none;}
:root{
  --gold:#D4A017;
  --black:#080808;
  --dark:#0F0F0F;
  --panel:#040404;
  --panel2:#1a1a1a;
  --border:#222;
  --text:#E8E8E8;
  --muted:#888;
  --green:#22C55E;
  --red:#EF4444;
  --mono:'Courier New',monospace;
}
.DA { position:fixed;    top: 0px; right:0; bottom: 0px; left:  0px; }
.DB { position:fixed;    top:66px; right:0; bottom: 0px; left:220px; }
.DC { position:fixed;    top:40px; right:0; bottom: 0px; left:100%; background:linear-gradient(to right,#181818,#111111 10%);}
.DD { position:fixed; height: 0px; right:0; bottom:38px; left:0; z-index: 1000; }
.FA { height:100%;width:100%;border:none;}
.FC { height:100%;width:100%;border:none;}
.CP { position: absolute; bottom: 3px; left: 0; right: 0; display: flex; align-items: center; padding: 0 8px; }
.EX { position: fixed; top: 110px; right: 18px; bottom: 0; width: 40px; height: 110px; overflow: hidden; cursor: pointer; text-align: center; }
#PR { flex:1;height:0px;accent-color:#58a6ff;cursor:pointer; }
.TL { position:fixed;top:13px;left:3px;cursor:pointer}
.TR { position:fixed;top:13px;left:3px;cursor:pointer}
.content { position: fixed; top: 0; right: 0; bottom: 0; left: 250px; padding: 0; }
footer{position:fixed;bottom:0;left:220px;right:0;z-index:80;
background:linear-gradient(to right,#000,#333,#000);border-top:1px solid var(--border);padding:13px 28px;text-align:center;
font-size:12px;color:var(--muted);}footer span{color:var(--gold);}
</style>
</head>
<body>
<nav><a href="/" class="brand">⛓ LegacyCoin <small>EXPLORER</small></a>
<div class="navl"><a href="/">1. Home</a><a href="/blocks">2. All Blocks</a></div><form class="sf" action="/search".
method="GET"><input type="text" name="q" placeholder="Height, hash, txid or address…"><button type="submit">→</button></form>
</nav>
<div id=DA class=DA><iframe name=fA id=fA class=FA src=/exp-A></iframe></div>
<div id=DB class=DB><iframe name=fB id=fB class=FA src=/exp-1></iframe></div>
<div id=DC class=DC><iframe name=fC id=fC class=FC           ></iframe></div>
<div id="DD" class="DD">
  <div class="CP">
    <input type="range" id="PR" min="0" max="100" value="50">
    <div id="TL" class="TL"><img onclick="clHH()" src="http://185.253.219.51:8484/hollaex/v.php?ix=myimg-Get&p4=hollaex" style="height:26px" alt="menu"></div>
    <div id="TR" class="TR"><img onclick="clHH()" src="http://185.253.219.51:8484/hollaex/v.php?ix=myimg-Get&p4=hollaex-1" style="height:26px" alt="menu"></div>
  </div>
</div>
<a href="/" class="logo" title="Home">
<img src="/lbtc.png" alt="LBTC" style=position:fixed;top:60px;height:50px></a>
<script>
const DA=document.getElementById('DA');
const DB=document.getElementById('DB');
const fB=document.getElementById('fB');
const DC=document.getElementById('DC');
const PR=document.getElementById('PR');
const TL=document.getElementById('TL');
const TR=document.getElementById('TR');
DB.style.left='220px';
DC.style.left='100%';
TR.style.display='none';
let isDragging  =false;
function lf(url){fB.src=url;}
function clHH(){const db=DB.style.left
 if(DC.style.left=='100%'){
  if(db=='220px'){DB.style.left='46px'; TR.style.display='table-row';TL.style.display='none';}
  else           {DB.style.left='220px';TL.style.display='table-row';TR.style.display='none';}
 }else DC.style.left='100%'
}
function clS(      ){k=PR.value;             CA.style.left=k+'%'}
function setB(     ){k=PR.value;             DC.style.left=k+'%'                                          }
function setBleft( ){k=PR.value;PR.value=--k;DC.style.left=k+'%';localStorage.setItem('panelPos',PR.value)}
function setBright(){k=PR.value;PR.value=++k;DC.style.left=k+'%';localStorage.setItem('panelPos',PR.value)}
PR.addEventListener('input',setB);
PR.addEventListener('mousedown' ,function(e){isDragging=true; document.body.style.cursor='col-resize';});
PR.addEventListener('touchstart',function(){isDragging=true;});
document.addEventListener('mouseup',function(){
  if(isDragging){
     isDragging=false;
     document.body.style.cursor='';
     localStorage.setItem('panelPos',PR.value);
  }
});
document.addEventListener('touchend',function(){
  if(isDragging){isDragging=false;localStorage.setItem('panelPos',PR.value);}
});
document.addEventListener('keydown',function(e){
  var c=e.keyCode;
  if(e.key==='ArrowLeft' ){e.preventDefault();setBleft(); return;}
  if(e.key==='ArrowRight'){e.preventDefault();setBright();return;}
  if(e.key==='Escape'    ){e.preventDefault();clHH();     return;}
  if(e.key==='Backspace'||c===8){var fi=document.querySelector('.sf input');if(fi){e.preventDefault();fi.focus();}return;}
  if(e.key==='1'||c===49||c===97){fB.src='/exp-1';return;}
  if(e.key==='2'||c===50||c===98){fB.src='/exp-2';return;}
});
document.addEventListener('DOMContentLoaded',function(){
  if(DC.style.left==='' || DC.style.left==='0%'){DC.style.left='100%';}
  PR.value=parseInt(localStorage.getItem('panelPos'))
});
let startX=0, startLeft=0;
PR.addEventListener('mousedown',function(e){startX=e.clientX;startLeft=parseInt(DC.style.left)||0;});
document.addEventListener('mousemove',function(e){
  if(isDragging){
     const newPos=startLeft+(e.clientX-startX)/window.innerWidth*100;
     PR.value     =newPos;
     DC.style.left=newPos+'%';
  }
});
</script>
<footer><span>LegacyCoin (LBTC)</span> Block Explorer · CPU money for everyone</footer>
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
