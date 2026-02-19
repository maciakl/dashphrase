BINARY_NAME := "dashphrase"

# defaults to build
all: build

# build the project
[group('build')]
build: tidy
	go build

# clean build artifacts
[group('util')]
clean:
	go clean
	-rm -rf dist
	-rm -rf coverage*

# run the project
[group('build')]
run: build
	./{{BINARY_NAME}}

# tidy up the go.mod and go.sum files
[group('build')]
tidy:
	go mod tidy
	go fmt ./...
	go vet ./...
	go mod verify

# run tests
[group('test')]
test: build
	clrz go test -v

# calculate test coverage
[group('test')]
coverage:
    go test -coverprofile=coverage ./...
    go tool cover -html=coverage -o coverage.html

# check coverage in a browser
[macos]
[group('test')]
check: coverage
	open coverage.html

# check coverage in a browser
[windows]
[group('test')]
check: coverage
   pwsh -c Start-Process coverage.html 

# release the project and generate scoop file
[group('release')]
release: build
	gopher release
	gopher scoop
