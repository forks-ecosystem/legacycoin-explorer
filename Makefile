BINARY = explorer

.PHONY: all build clean

all: build

build:
	go build -v -o $(BINARY) ./cmd/explorer

clean:
	rm -f $(BINARY)

run: build
	./$(BINARY) -rpcpassword=yourpass
