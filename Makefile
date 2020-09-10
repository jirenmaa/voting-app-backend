MAIN_SRC := cmd/server/main.go

ifeq ($(shell echo "check_quotes"),"check_quotes")
OUTFILE := bin/server-windows-386
else
OUTFILE := bin/server-linux-386
endif

.PHONY: help clean build build-all docs run start all
# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## Deletes all compiled / executable files
	@find bin -type f -not -name '.gitkeep' -print0 | xargs -0 rm --
	@echo "Deleted all build files"

docs: ## Generates api documentation
	@cd cmd/server && $(GOPATH)/bin/swag init && cd ../..

build: ## Compile the go files
	@echo "Building go file..."
	@go build -o $(OUTFILE) $(MAIN_SRC)

build-all: ## Compile the go files for multiple OS
	@echo "Building go files for multiple OS..."
	@GOOS=linux GOARCH=arm go build -o bin/server-linux-arm $(MAIN_SRC)
	@GOOS=linux GOARCH=arm64 go build -o bin/server-linux-arm64 $(MAIN_SRC)
	@GOOS=linux GOARCH=386 go build -o bin/server-linux-386 $(MAIN_SRC)
	@GOOS=freebsd GOARCH=386 go build -o bin/server-freebsd-386 $(MAIN_SRC)
	@GOOS=windows GOARCH=386 go build -o bin/server-windows-386 $(MAIN_SRC)
	@echo "Finished building"

run: ## Runs the server
	@APP_ENV=development GIN_MODE=debug go run $(MAIN_SRC)

start: ## Runs the compiled server in production mode
	@APP_ENV=production GIN_MODE=release $(OUTFILE)

all: build start ## Build and Run the server

.DEFAULT_GOAL := help