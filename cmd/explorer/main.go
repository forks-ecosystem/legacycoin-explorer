// legacycoin-explorer — Block explorer for LegacyCoin (LBTC)
//
// Usage:
//
//	./explorer -nodehost=127.0.0.1 -nodeport=19556 -rpcuser=coin -rpcpassword=coin
//	./explorer -port=8084
//
// Then open http://localhost:8084 in your browser.
//
// All options can be overridden via environment variables (see .env):
//
//	EXPLORER_NODE_HOST, EXPLORER_NODE_PORT, EXPLORER_RPC_USER,
//	EXPLORER_RPC_PASSWORD, EXPLORER_COOKIE_FILE, EXPLORER_PORT, EXPLORER_DATA_DIR
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	explorer "github.com/legacycoin/explorer"
)

func envStr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	if v := os.Getenv(key); v != "" {
		if n, err := strconv.Atoi(v); err == nil {
			return n
		}
	}
	return def
}

func main() {
	nodeHost := flag.String("nodehost", envStr("EXPLORER_NODE_HOST", "127.0.0.1"), "legacycoind hostname")
	nodePort := flag.Int("nodeport", envInt("EXPLORER_NODE_PORT", 19556), "legacycoind RPC port")
	rpcUser := flag.String("rpcuser", envStr("EXPLORER_RPC_USER", ""), "RPC username (overrides cookie)")
	rpcPass := flag.String("rpcpassword", envStr("EXPLORER_RPC_PASSWORD", ""), "RPC password (overrides cookie)")
	cookieFile := flag.String("cookiefile", envStr("EXPLORER_COOKIE_FILE", "/home/coin/.legacycoin/.cookie"), "Path to .cookie file for RPC auth")
	httpPort := flag.Int("port", envInt("EXPLORER_PORT", 8084), "Explorer HTTP port")
	dataDir := flag.String("data", envStr("EXPLORER_DATA_DIR", "/data"), "Directory for persistent data (bookmarks)")
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

	srv := explorer.NewServer(rpc, *httpPort, *dataDir)
	srv.Start()
}
