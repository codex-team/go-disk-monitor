BINARY_NAME=diskMonitor
BINARY_NAME_LINUX=$(BINARY_NAME)-linux
BINARY_NAME_DARWIN=$(BINARY_NAME)-darwin

export GO111MODULE=on

all: check lint build

build:
	go build -o $(BINARY_NAME) -v ./
	chmod +x $(BINARY_NAME)
test:
	go test ./...
lint:
	golint ./
clean:
	go clean
	rm -rf $(BINARY_NAME)
	rm -rf $(BINARY_NAME_LINUX)
	rm -rf $(BINARY_NAME_DARWIN)

build-all: build-linux build-darwin

build-linux:
	GOOS=linux GOARCH=amd64 go build -o $(BINARY_NAME_LINUX) -v ./

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o $(BINARY_NAME_DARWIN) -v ./