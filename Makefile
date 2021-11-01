lint:
	golangci-lint run ./...

test:
	go test ./...

gen:
	go generate ./...

pre-commit: lint test

install-tools:
	go get -u golang.org/x/tools/cmd/stringer
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(shell go env GOPATH)/bin v1.42.1