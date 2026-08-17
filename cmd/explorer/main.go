// legacycoin-explorer — Block explorer for LegacyCoin (LBTC)
//
// Usage:
//   ./explorer -node=127.0.0.1:19556 -rpcuser=legacycoin -rpcpassword=yourpass
//   ./explorer -port=8080
//
// Then open http://localhost:8080 in your browser.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	explorer "github.com/legacycoin/explorer"
)

func main() {
	nodeHost := flag.String("nodehost", "127.0.0.1", "legacycoind hostname")
	nodePort := flag.Int("nodeport", 19556, "legacycoind RPC port")
	rpcUser  := flag.String("rpcuser", "", "RPC username (overrides cookie)")
	rpcPass  := flag.String("rpcpassword", "", "RPC password (overrides cookie)")
	cookieFile := flag.String("cookiefile", "/home/coin/.legacycoin/.cookie", "Path to .cookie file for RPC auth")
	httpPort := flag.Int("port", 8084, "Explorer HTTP port")
	flag.Parse()

	user := *rpcUser
	pass := *rpcPass

	if user == "" && pass == "" && *cookieFile != "" {
		data, err := os.ReadFile(*cookieFile)
		if err == nil {
			parts := strings.SplitN(strings.TrimSpace(string(data)), ":", 2)
			if len(parts) == 2 {
				user = parts[0]
				pass = parts[1]
				log.Printf("Loaded RPC credentials from %s", *cookieFile)
			}
		} else {
			log.Printf("Cookie file %s not found, using defaults", *cookieFile)
			user = "coin"
			pass = "coin"
		}
	}

	if pass == "" {
		fmt.Fprintln(os.Stderr, "ERROR: RPC password required (-rpcpassword or -cookiefile)")
		os.Exit(1)
	}

	rpc := explorer.NewRPCClient(*nodeHost, *nodePort, user, pass)
	if !rpc.Ping() {
		log.Printf("WARNING: Cannot connect to legacycoind at %s:%d — explorer will show offline state", *nodeHost, *nodePort)
	} else {
		log.Printf("Connected to legacycoind at %s:%d", *nodeHost, *nodePort)
	}

	srv := explorer.NewServer(rpc, *httpPort)
	srv.Start()
}
