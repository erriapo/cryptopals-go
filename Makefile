export PKG := github.com/erriapo/cryptopals-go
PKG_TEST := $(PKG)/set1

# Set up the development environment
env:
	@echo "PWD: $(PWD)"
	@echo "BASE_DIR: $(BASE_DIR)"
	@echo "GOPATH: $(GOPATH)"
	@echo "GOROOT: $(GOROOT)"
	@echo "DEST: $(DEST)"
	@echo "PKG: $(PKG)"

build:
	go build $(PKG)

test:
	go test -race -test.timeout 120s $(PKG_TEST)
