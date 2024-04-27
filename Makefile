.PHONY: run
run: build
	@echo "running..."
	@./bin/go-user-api

.PHONY: build
build:
	@echo "building..."
	@go build -o ./bin/go-user-api ./cmd/go-user-api/main.go

.PHONY: watch
watch:
	@${HOME}/go/bin/air

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	@echo "removing bin/ files"
	@rm ./bin/*
	@echo "removing tmp/ files"
	@rm ./tmp/*

.PHONY: clean-scylla
clean-scylla:
	@echo "cleaning local scylladb..."
	@sudo docker exec -it go-user-api-db-1 cqlsh -e "DROP KEYSPACE go_api; CREATE KEYSPACE go_api WITH replication = {'class': 'SimpleStrategy', 'replication_factor': '1'};"
