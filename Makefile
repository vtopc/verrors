.PHONY: deps
deps:
	go mod download

.PHONY: test
test:
	go test `go list ./... | grep -v '/gen' | grep -v '/protos'` -cover -count=1 -coverprofile=coverage.txt -covermode=count

# linter:
GOLINT = $(GOPATH)/bin/golangci-lint
$(GOLINT):
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin v1.33.0

.PHONY: lint
lint: $(GOLINT)
	$(GOLINT) run
