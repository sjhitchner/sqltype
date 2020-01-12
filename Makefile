PROJECT_NAME=sqltype
PROJECT_PATH=github.com/sjhitchner/sqltype

BRANCH=`git rev-parse --abbrev-ref HEAD`
SHA=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X github.com/sjhitchner/sqltype.Sha=$(SHA) -X github.com/sjhitchner/sqltype.Branch=$(BRANCH)"

CMD_MAIN="./cmd/sqltype"
CMD_BINARY_NAME="sqltype"

BIN_DIR="./bin"

GOFILES := $(wildcard *.go)

export GO111MODULE=on

default: build

all: test build

build: 
	go build $(LDFLAGS) -o $(BIN_DIR)/$(CMD_BINARY_NAME) $(CMD_MAIN)

install: build
	go install $(CMD_MAIN)

vet:
	go vet ./...

test: vet
	go test -v -cover ./...

clean:
	rm -f $(BIN_DIR)/$(CMD_BINARY_NAME)
	go clean ./...

run: build
	$(BIN_DIR)/$(CMD_BINARY_NAME)

.PHONY: all build test clean run
