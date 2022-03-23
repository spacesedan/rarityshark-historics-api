.PHONY: build clean deploy run

build:
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/historics historics/main.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose


run:
	go run historics/main.go
