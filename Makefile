default: out/example

clean:
	rm -rf out

test: *.go
	go test ./...

out/example: implementation.go cmd/example/main.go
	mkdir out
	go build -o out/example ./cmd/example