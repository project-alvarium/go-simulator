.PHONY: build clean run test

build:
	GOCGO=CGO_ENABLED=1 go build -o go-simulator

clean:
	rm ./go-simulator

run:
	cd bin && ./launch.sh

test:
	go test -coverprofile=coverage.out ./...
	go vet ./...
	gofmt -l .
	[ "`gofmt -l .`" = "" ]