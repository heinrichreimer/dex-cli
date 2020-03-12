DEX_VERSION=v2.22.0

build: build-linux build-windows

build-linux:
	go build -o bin/dex-cli -v -i main.go
	chmod +x bin/dex-cli

build-windows:
	GOOS=windows GOARCH=386 go build -o bin/dex-cli.exe -v -i main.go
