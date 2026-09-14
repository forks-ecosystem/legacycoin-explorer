package explorer

import (
	_ "embed"
	"net/http"
	"strconv"
)

//go:embed lbtc.png
var lbtcPNG []byte

//go:embed left.png
var leftPNG []byte

//go:embed right.png
var rightPNG []byte

//go:embed favicon.ico
var faviconICO []byte

type staticAsset struct {
	data []byte
	ct   string
}

var staticAssets = map[string]staticAsset{
	"/left.png":    {leftPNG, "image/png"},
	"/right.png":   {rightPNG, "image/png"},
	"/favicon.ico": {faviconICO, "image/x-icon"},
}

// handleLogo serves the sidebar LBTC logo.
func (s *Server) handleLogo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(lbtcPNG)))
	w.Write(lbtcPNG)
}

// handleStatic serves embedded static assets (toggle icons, favicon).
func (s *Server) handleStatic(w http.ResponseWriter, r *http.Request) {
	a, ok := staticAssets[r.URL.Path]
	if !ok {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", a.ct)
	w.Header().Set("Content-Length", strconv.Itoa(len(a.data)))
	w.Write(a.data)
}
