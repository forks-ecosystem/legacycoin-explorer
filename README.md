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
make build
./explorer \
  -nodehost=127.0.0.1 \
  -nodeport=19556 \
  -rpcuser=coin \
  -rpcpassword=coin \
  -port=8084
```

Then open http://localhost:8084

## Configuration (.env)

All options are overridable via environment variables (with `-flag` taking
precedence). Runtime configuration lives in `.env` (gitignored; start from
`.env.example`). The config is used by Docker via `env_file`:

| Variable              | Default | Description                          |
|-----------------------|---------|--------------------------------------|
| `EXPLORER_NODE_HOST`  | `127.0.0.1` | legacycoind hostname             |
| `EXPLORER_NODE_PORT`  | `19556`  | legacycoind RPC port                 |
| `EXPLORER_RPC_USER`   | *(none)* | RPC username (overrides cookie)      |
| `EXPLORER_RPC_PASSWORD` | *(none)* | RPC password (overrides cookie)   |
| `EXPLORER_COOKIE_FILE` | `/home/coin/.legacycoin/.cookie` | cookie file when user/pass empty |
| `EXPLORER_PORT`       | `8084`   | Explorer HTTP port                   |
| `EXPLORER_DATA_DIR`   | `/data`  | Persistent dir for bookmarks         |

## Docker

```bash
docker compose up -d --build
```

- Listens on `0.0.0.0:8084`.
- Uses `network_mode: host` because legacycoind RPC is bound to `127.0.0.1`.
- Bookmarks persist in `./data/`.
- Configuration comes from `.env` (see above).

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
| Explorer | 8084 |
| Node RPC | 19556 (legacycoind) |
