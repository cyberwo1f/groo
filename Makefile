
.PHONY: all macos

all: macos

macos:
	@echo "Compiling macos binaries"
	env GOOS=darwin GOARCH=amd64 go build -o dist/macos/amd64/groo
	chmod a+x ./dist/macos/amd64/groo
