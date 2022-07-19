GOCMD:=go
GOBUILD:=$(GOCMD) build
GOTEST=$(GOCMD) test

BINARY_NAME:=mend
PACKAGE_PATH:=./main.go

generate:
	if ! which mockery ; echo "installing mockery" & then go install github.com/vektra/mockery/v2@v2.12.0 ; fi
	go generate ./...

build:
	GOARCH="amd64" GOOS="linux" $(GOBUILD) -o $(BINARY_NAME) $(PACKAGE_PATH)


# golangci-lint needed:
# curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.47.0
lint:
	golangci-lint run

test:
	$(GOTEST) ./...