echo "Cross-compiling"

go build

echo "Regular binary compiled"

GOOS=windows
GOARCH=amd64
go build -o pokedex.exe

echo ".exe compiled"
