run:
	go run cmd/main.go

test_coverage:
	go test -cover ./...

test:
	go test ./...