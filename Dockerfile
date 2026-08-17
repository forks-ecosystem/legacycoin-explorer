FROM golang:1.21-alpine AS builder
WORKDIR /src
COPY go.mod ./
RUN go mod download || true
COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /explorer ./cmd/explorer

FROM alpine:3.19
RUN apk add --no-cache ca-certificates && mkdir -p /data
COPY --from=builder /explorer /usr/local/bin/explorer
EXPOSE 8084
ENTRYPOINT ["explorer"]