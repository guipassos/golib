.PHONY: help mocks test docs

help:
	@echo
	@echo "| Commands        | Descriptions                                                                                       |"
	@echo "---------------------------------------------------------------------------------------------------------------------- |"
	@echo "| help            | Lists all commands.                                                                                |"
	@echo "| install         | Install external dependencies                                                                      |"
	@echo "| mocks           | Generate all go mocks based on ports path.                                                         |"
	@echo "| mocks           | Generate all go mocks based on ports path.                                                         |"
	@echo "| test            | Run all tests and returns results and coverage to the console                                      |"
	@echo

test: 
	go test -v -p 1 -cover -failfast ./... -coverprofile=coverage.out
	@go tool cover -func coverage.out | awk 'END{print sprintf("coverage: %s", $$3)}'

mocks:
	go generate ./...

install:
	go install github.com/golang/mock/mockgen@latest
