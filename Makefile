GO_CMD = go
VARSION = -u

build:
	$(GO_CMD) build -o ./target/client ./cmd/client/main.go
	$(GO_CMD) build -o ./target/server ./cmd/server/main.go

fmt:
	$(GO_CMD) fmt ./...