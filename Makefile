MAIN := main.go

ifeq ($(shell echo "check_quotes"),"check_quotes")
OUTFILE := bin/server-windows-386
else
OUTFILE := bin/server-linux-386
endif

.PHONY: help
# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

.PHONY: build
build: ## Compile the go files
	go build -o $(OUTFILE) $(MAIN)

.PHONY: build-all
build-all: ## Compile the go files for multiple OS
	GOOS=linux GOARCH=arm go build -o bin/server-linux-arm $(MAIN)
	GOOS=linux GOARCH=arm64 go build -o bin/server-linux-arm64 $(MAIN)
	GOOS=linux GOARCH=386 go build -o bin/server-linux-386 $(MAIN)
	GOOS=freebsd GOARCH=386 go build -o bin/server-freebsd-386 $(MAIN)
	GOOS=windows GOARCH=386 go build -o bin/server-windows-386 $(MAIN)


.PHONY: run-server
run-server: ## Run the server
	$(OUTFILE)

.PHONY: dev
dev: ## Runs the development server
	go run $(MAIN)

all: ## Runs all the main commands
	make build && make run-server