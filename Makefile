build:
	go build -ldflags="-s -w" gen.go
	$(if $(shell command -v upx), upx gen)

mac:
	GOOS=darwin go build -ldflags="-s -w" -o gen-darwin gen.go
	$(if $(shell command -v upx), upx gen-darwin)

win:
	GOOS=windows go build -ldflags="-s -w" -o gen.exe gen.go
	$(if $(shell command -v upx), upx gen.exe)

linux:
	GOOS=linux go build -ldflags="-s -w" -o gen-linux gen.go
	$(if $(shell command -v upx), upx gen-linux)

