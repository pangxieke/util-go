fmt:
	go fmt ./...

test:
	go test ./... -v

bench:
	go test ./... -bench=.
