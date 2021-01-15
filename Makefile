VERSION=0.0.1
LDFLAGS=-ldflags "-w -s -X main.version=${VERSION}"
GO111MODULE=on

all: dhtcli

.PHONY: dhtcli

dhtcli: main.go
	GOOS=linux GOARCH=arm go build $(LDFLAGS) -o dhtcli

mac: main.go
	go build $(LDFLAGS) -o dhtcli

clean:
	rm -rf dhtcli

check:
	go test ./...

tag:
	git tag v${VERSION}
	git push origin v${VERSION}
	git push origin master
