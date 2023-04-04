GOCMD=go
CODE_COVERAGE=code-coverage
build: ${BINARY_DIR} ## Compile the code, build Executable File
	$(GOCMD) build -o $(BINARY_DIR) -v ./cmd/api

run: ## Start application
	$(GOCMD) run ./cmd/api

test: ## Run tests
	$(GOCMD) test ./... -cover

test-coverage: ## Run tests and generate coverage file
	$(GOCMD) test ./... -coverprofile=$(CODE_COVERAGE).out
	$(GOCMD) tool cover -html=$(CODE_COVERAGE).out

wire: ## Generate wire_gen.go
	cd pkg/di && go run github.com/google/wire/cmd/wire@latest

# swag: ## Generate swagger docs
# 	cd cmd/api && swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs
swag :## generate swagger docs
	swag init -g pkg/api/handler/adminHandler.go -o ./cmd/api/docs

mock :#generate mock data
	mockgen -source=pkg/repository/interfaces/authInterface.go -destination=pkg/mock/authmockRepo.go -package=mock