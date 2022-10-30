build:
	go build -o bin/safetymap

run: build
	go run ./bin/safetymap

test:
	go test -v ./...