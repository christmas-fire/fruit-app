.PHONY: build
build:
	@go build -o ./bin/fruit-app cmd/main.go

.PHONY: run
run: build
	@./bin/fruit-app
