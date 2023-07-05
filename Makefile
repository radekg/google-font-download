build:
	go build -o $$GOPATH/bin/google-font-download ./cmd/download

.PHONY: test
test:
	go test ./... -count=1 -v
