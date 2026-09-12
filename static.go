package explorer

import (
	_ "embed"
	"net/http"
	"strconv"
)

//go:embed lbtc.png
var lbtcPNG []byte

// handleLogo serves the sidebar LBTC logo.
func (s *Server) handleLogo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/png")
	w.Header().Set("Content-Length", strconv.Itoa(len(lbtcPNG)))
	w.Write(lbtcPNG)
}
