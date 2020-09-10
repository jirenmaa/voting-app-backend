MAIN_SRC := main.go

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

.PHONY: clean
clean: ## Deletes all compiled / executable files
	@find bin -type f -not -name '.gitkeep' -print0 | xargs -0 rm --

.PHONY: build
build: ## Compile the go files
	@echo "Building go files"
	@go build -o $(OUTFILE) $(MAIN_SRC)

.PHONY: build-all
build-all: ## Compile the go files for multiple OS
	@echo "Building go files for multiple OS"
	GOOS=linux GOARCH=arm go build -o bin/server-linux-arm $(MAIN_SRC)
	GOOS=linux GOARCH=arm64 go build -o bin/server-linux-arm64 $(MAIN_SRC)
	GOOS=linux GOARCH=386 go build -o bin/server-linux-386 $(MAIN_SRC)
	GOOS=freebsd GOARCH=386 go build -o bin/server-freebsd-386 $(MAIN_SRC)
	GOOS=windows GOARCH=386 go build -o bin/server-windows-386 $(MAIN_SRC)

.PHONY: run
run: ## Runs the server
	@APP_ENV=development go run $(MAIN_SRC)

.PHONY: start
start: ## Runs the compiled server
	@APP_ENV=production $(OUTFILE)

.PHONY: all
all: build start ## Build and Run the server
