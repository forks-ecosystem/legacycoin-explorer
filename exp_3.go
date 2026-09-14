// exp_3.go — страница «3. API»: краткая документация публичного API
// для бирж/обменников: проверка транзакции и баланса.
package explorer

import (
	"html/template"
	"net/http"
)

var exp3Template = template.Must(template.New("exp3").Parse(`<style>
:root{--gold:#D4A017;--black:#080808;--dark:#0F0F0F;--panel:#040404;--panel2:#1a1a1a;--border:#222;--text:#E8E8E8;--muted:#888;--green:#22C55E;--red:#EF4444;--mono:'Courier New',monospace;}
*{box-sizing:border-box;margin:0;padding:0;}
html,body{height:100%;overflow:hidden;}
body{background:var(--black);color:var(--text);font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',sans-serif;font-size:15px;line-height:1.6;display:flex;flex-direction:column;}
.pt{font-size:21px;font-weight:700;margin-bottom:14px;flex-shrink:0;padding-top:10px}
.pt span{color:var(--gold);}
.wrap{flex:1;min-height:0;overflow:auto;padding:0 4px 60px 0;outline:none}
.card{background:var(--panel);border-radius:8px;padding:16px 18px;margin-bottom:14px;}
.card h3{font-size:14px;color:var(--gold);margin-bottom:8px;}
.card p{font-size:13px;color:var(--muted);margin:6px 0;}
code.k{display:inline-block;background:var(--panel2);border:1px solid var(--border);border-radius:4px;padding:2px 8px;font-family:var(--mono);font-size:12.5px;color:var(--text);}
pre.ex{border-radius:6px;padding:12px 14px;font-family:var(--mono);font-size:12px;color:#9FE88F;overflow-x:auto;white-space:pre;margin:8px 0;}
table.ep{border-collapse:separate;border-spacing:0;margin-top:6px;}
table.ep th{text-align:left;color:var(--muted);font-size:11px;text-transform:uppercase;letter-spacing:1px;padding:6px 10px;border-bottom:1px solid var(--border);}
table.ep td{padding:8px 10px;font-size:13px;border-bottom:1px solid var(--border);vertical-align:top;}
table.ep td code.k{white-space:nowrap;}
a{color:var(--gold);text-decoration:none;}
a:hover{text-decoration:underline;}
.g{color:var(--green);}
.r{color:var(--red);}
.m{color:var(--muted);}
</style>
<div class="pt"><span style=color:#555>3.</span> API <span>для бирж</span></div>
<div class="wrap">
<div class="card">
<h3>Базовый URL</h3>
<p>Все запросы — <code class="k">GET</code>, ответ — JSON. Отвечает любым origin (CORS <code class="k">Access-Control-Allow-Origin: *</code>).
Ошибки: <code class="k">404</code> с <code class="k">{"error":"..."}</code>.</p>
<pre class="ex">https://185.253.219.51:8084/api/...</pre>
</div>
<div class="card">
<h3>Проверка транзакции</h3>
<p>Наличие транзакции, блок и число подтверждений:</p>
<pre class="ex">GET /api/tx/{txid}
{
  "txid":"...",
  "blockhash":"...",
  "blockheight":12461,
  "confirmations":1,
  "time":1768000000,
  "blocktime":1768000000,
  "vin":[...],
  "vout":[{"value":50.0,"scriptPubKey":{"addresses":["L..."]}}]
}</pre>
<p>Правило для биржи: принимать перевод после <b class="g">N подтверждений</b>
(поле <code class="k">confirmations</code>, обычно N&nbsp;=&nbsp;6+). Если транзакция не найдена — ответ <code class="k">404</code>, значит она ещё не в блокчейне.</p>
</div>
<div class="card">
<h3>Баланс адреса</h3>
<pre class="ex">GET /api/balance/{address}
{
  "address":"L...",
  "balance":1.5,
  "balance_base_units":150000000,
  "received":2.0,
  "received_base_units":200000000
}</pre>
<p><code class="k">balance_base_units</code> — баланс в самых мелких единицах (сатоши).
<code class="k">balance</code> — то же в LBTC. Только подтверждённые выходы.</p>
</div>
<div class="card">
<h3>Адрес + история</h3>
<pre class="ex">GET /api/address/{address}
{
  "address":{"address":"L...","isvalid":true},
  "txs":[{...}, {...}]
}</pre>
<p>Валидация адреса и список транзакций адреса (из адрес-индекса ноды), каждая — с <code class="k">vin/vout</code> и <code class="k">confirmations</code>.</p>
</div>
<div class=card>
<h3>Блоки и статус</h3>
<table class=ep>
<tr><th>Endpoint</th><th>Что возвращает</th></tr>
<tr><td><code class=k>/api/block/{height|hash}</code></td><td>Блок: height, hash, time, tx, size, confirmations</td></tr>
<tr><td><code class=k>/api/blocks</code></td><td>Последние блоки (height, hash, txs, size)</td></tr>
<tr><td><code class=k>/api/stats </code></td><td>Сложность, hashrate, последние блоки</td></tr>
<tr><td><code class=k>/api/status</code></td><td>Статус ноды: блок, txs-в-мемпуле, транзакции и т.д.</td></tr>
</table>
</div>
</div>
<script>
document.addEventListener('keydown',function(e){
 var k=e.key
  if(k==='1'         ){e.preventDefault();parent.lf('/exp-1');return}
  if(k==='2'         ){e.preventDefault();parent.lf('/exp-2');return}
  if(k==='3'         ){e.preventDefault();parent.lf('/exp-3');return}
  if(k==='Escape'    ){e.preventDefault();parent.clkHH(     );return}
  if(k==='Backspace' ){e.preventDefault();parent.setSF(     );return}
  if(k==='ArrowLeft' ){e.preventDefault();parent.setBL(     );return}
  if(k==='ArrowRight'){e.preventDefault();parent.setBR(     );return}
});
const wrap=document.querySelector('.wrap');
wrap.tabIndex=-1;
wrap.focus();
</script>
`))

func (s *Server) handleExp3(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
    if err := exp3Template.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}
