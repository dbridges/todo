.PHONY: run run-with-race clean install uninstall release test

VERSION=0.1.0
GO_SRC = $(shell find . -iname '*.go')
BINDIR?=/usr/local/bin
BINNAME?=todo

all: dist/$(BINNAME)

dist/$(BINNAME): $(GO_SRC) dist
	go build -ldflags "-X main.Version=$(VERSION)" -o $@

run:
	@go run $(GO_SRC)

run-with-race:
	@GORACE="log_path=race_log" go run -race *.go

clean:
	rm -f dist/*
	rm -f race_log.*

install: all
	mkdir -p $(BINDIR)
	install dist/$(BINNAME) $(BINDIR)/$(BINNAME)

dist:
	mkdir dist

uninstall:
	rm -f $(BINDIR)/$(BINNAME)

test:
	go test ./...
