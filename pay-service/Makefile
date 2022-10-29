help: 			## Show this help, targets are ordered by alphabet
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sort | sed 's/:.*##/#/' | column -t -c 2 -s '#'

gen-swag: ## Generate swagger docs
	swag init

run: ## Run app locally
	CONFIG_PATH=config/local.yaml go run main.go

test: ## Run unit test
	go test -cover -race ./...

build-web: ## Build web
	cd ./web && yarn install && yarn build
	rm -rf views static
	cd ./web && mv build/ ../
	mv build views
	mv ./views/static ./

test-cover: ## Run unit test with report coverage
	@go test -race -v -coverpkg=./... $(shell go list ./...) -coverprofile=cover.out
	@go tool cover -func=cover.out
	@rm -rf cover.out