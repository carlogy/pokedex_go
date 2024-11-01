run:
	go build -o pokedex && ./pokedex

compile:

	go build -o pokedex

	GOOS=windows GOARCH=amd64 go build -o pokedex.exe

test:
	go test ./...

test-v:
	go test ./... -v
