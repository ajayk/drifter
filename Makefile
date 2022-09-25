export GOPROXY=direct
UNAME_S = $(shell uname -s)
UNAME_M = $(shell uname -m)
GO_INSTALL_FLAGS=-ldflags="-s -w"
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
TARGET := drifter

build: fmt
	go mod tidy
ifeq ($(UNAME_S), Darwin)
	time GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -o $(TARGET) $(GO_INSTALL_FLAGS) $V
else
	time GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o $(TARGET) $(GO_INSTALL_FLAGS) $V
endif
	@chmod +X $(TARGET)

fmt:
	@gofmt -l -w $(SRC)

clean:
	@rm drifter

.PHONY: build fmt clean
