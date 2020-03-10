DEX_VERSION=v2.22.0

build: build-linux build-windows

build-linux: download-api
	go build -o bin/dex-cli -v -i main.go
	chmod +x bin/dex-cli

build-windows: download-api
	GOOS=windows GOARCH=386 go build -o bin/dex-cli.exe -v -i main.go

download-api:
	wget -nd https://raw.githubusercontent.com/dexidp/dex/${DEX_VERSION}/api/api.proto
