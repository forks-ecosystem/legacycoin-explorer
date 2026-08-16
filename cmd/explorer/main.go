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

	explorer "github.com/legacycoin/explorer"
)

func main() {
	nodeHost := flag.String("nodehost", "127.0.0.1", "legacycoind hostname")
	nodePort := flag.Int("nodeport", 19556, "legacycoind RPC port")
	rpcUser  := flag.String("rpcuser", "coin", "RPC username")
	rpcPass  := flag.String("rpcpassword", "coin", "RPC password")
	httpPort := flag.Int("port", 8084, "Explorer HTTP port")
	flag.Parse()

	if *rpcPass == "" {
		fmt.Fprintln(os.Stderr, "ERROR: -rpcpassword is required")
		fmt.Fprintln(os.Stderr, "Usage: ./explorer -rpcpassword=yourpass")
		os.Exit(1)
	}

	rpc := explorer.NewRPCClient(*nodeHost, *nodePort, *rpcUser, *rpcPass)
	if !rpc.Ping() {
		log.Printf("WARNING: Cannot connect to legacycoind at %s:%d — explorer will show offline state", *nodeHost, *nodePort)
	} else {
		log.Printf("Connected to legacycoind at %s:%d", *nodeHost, *nodePort)
	}

	srv := explorer.NewServer(rpc, *httpPort)
	srv.Start()
}
