server:
	go run framework/cmd/server/server.go

test:
	go test ./...

.PHONY: test server