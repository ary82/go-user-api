.PHONY: run
run: build
	./bin/go-user-api

.PHONY: build
build:
	go build -o ./bin/go-user-api ./cmd/go-user-api/main.go

.PHONY: watch
watch:
	@${HOME}/go/bin/air

.PHONY: test
test:
	go test -v ./...
