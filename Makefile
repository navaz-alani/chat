GXX=go
BINS=chat client
CHAT_SOURCES=$(wildcard cmd/service/*.go service/*.go payload/*.go)
CLIENT_SOURCES=$(wildcard cmd/client/*.go payload/*.go)

all:
	make chat client

chat: $(CHAT_SOURCES)
	$(GXX) build -o chat ./cmd/service

client: $(CLIENT_SOURCES)
	$(GXX) build -o client ./cmd/client

.PHONY: clean tidy

tidy:
	# tidy up dependencies
	go mod tidy
	# format all source code
	go fmt ./...

clean:
	rm $(BINS)
