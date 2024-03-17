all: help

.PHONY:
help:
	@echo "courses aggregator"
	@echo ""
	@echo "run                       Build docker and run application"
	@echo "test                      Run tests locally"
	@echo "fmt                       Run go fmt locally"
	@echo ""

run:
	@docker-compose build
	@docker-compose up

test:
	go test ./...

fmt:
	gofmt -l -s -w .
