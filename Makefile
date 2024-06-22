BIN_NAME=tu

build-windows:
	GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w -X 'github.com/mistweaverco/tu/cmd/tu.VERSION=$(VERSION)'" -o dist/$(BIN_NAME).exe main.go
build-linux:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags "-s -w -X 'github.com/mistweaverco/tu/cmd/tu.VERSION=$(VERSION)'" -o dist/$(BIN_NAME)-linux main.go
build-macos:
	GOOS=darwin GOARCH=arm64 CGO_ENABLED=0 go build -ldflags "-s -w -X 'github.com/mistweaverco/tu/cmd/tu.VERSION=$(VERSION)'" -o dist/$(BIN_NAME)-macos main.go

builds: build-linux build-macos build-windows

lint:
	golangci-lint run

test:
	if [ -n $(COLORS_ENABLED) ] then gotest ./... ; else go test ./...; fi

test-coverage:
	go test -race -covermode=atomic -coverprofile=coverage.out ./...

release:
	gh release create --generate-notes v$(VERSION) dist/$(BIN_NAME)-linux dist/$(BIN_NAME)-macos dist/$(BIN_NAME).exe

build-and-install-local-debug:
	go build -ldflags "-X 'github.com/mistweaverco/tu/cmd/tu.VERSION=development'" -o dist/$(BIN_NAME) main.go
	sudo mv dist/$(BIN_NAME) /usr/bin/$(BIN_NAME)

