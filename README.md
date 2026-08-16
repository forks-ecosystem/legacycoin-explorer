# LegacyCoin Block Explorer

A lightweight block explorer for LegacyCoin (LBTC).
Connects to a running legacycoind node via JSON-RPC and serves a web UI.

## Build

```bash
go mod tidy
make build
```

## Run

```bash
./explorer -rpcpassword=yourpass

# All options:
./explorer \
  -nodehost=127.0.0.1 \
  -nodeport=19556 \
  -rpcuser=legacycoin \
  -rpcpassword=legacy123 \
  -port=8080
```

Then open http://localhost:8080

## Features

- Home page: node stats + last 20 blocks
- /blocks: paginated full block list
- /block/<height or hash>: full block detail with tx list + prev/next navigation
- /search: search by height or hash
- /api/stats: JSON stats API
- /api/blocks: JSON recent blocks
- /api/block/<id>: JSON single block
- 5-second RPC cache (safe to expose publicly)
- Shows "OFFLINE" gracefully when node is unreachable

## Ports

| Service  | Port |
|----------|------|
| Explorer | 8080 |
| Node RPC | 19556 (legacycoind) |
