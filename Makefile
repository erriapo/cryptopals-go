export PKG := github.com/erriapo/cryptopals-go
PKG_TEST := $(PKG)/set1

# Set up the development environment
.PHONY: env
env:
	@echo "PWD: $(PWD)"
	@echo "BASE_DIR: $(BASE_DIR)"
	@echo "GOPATH: $(GOPATH)"
	@echo "GOROOT: $(GOROOT)"
	@echo "GOBIN: $(GOBIN)"
	@echo "DEST: $(DEST)"
	@echo "PKG: $(PKG)"

.PHONY: build
build:
	go build $(PKG)

.PHONY: format
format:
	go fmt $(PKG_TEST)

.PHONY: test
test:
	go test -race -test.timeout 120s $(PKG_TEST)

.PHONY: lint
lint:
	${GOBIN}/golint $(PKG_TEST)
