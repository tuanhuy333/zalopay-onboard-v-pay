PROJECT_ROOT = $(shell pwd)
MODULE_NAME ?= "order-service"
PROJECT_NAME ?= order-service
DOCKER_COMPOSE ?= cd devstack && docker-compose -p $(PROJECT_NAME)
GOOSE ?= goose -dir ./migrations mysql "root:root@tcp(localhost:3306)/v_pay_orders?parseTime=true"


help: 			## Show this help, targets are ordered by alphabet
	@echo 'usage: make [target] ...'
	@echo ''
	@echo 'targets:'
	@egrep '^(.+)\:\ .*##\ (.+)' ${MAKEFILE_LIST} | sort | sed 's/:.*##/#/' | column -t -c 2 -s '#'


run: ## Run app locally
	CONFIG_PATH=config/local.yaml go run main.go

test: ## Run unit test
	go test -cover -race ./...


test-cover: ## Run unit test with report coverage
	@go test -race -v -coverpkg=./... $(shell go list ./...) -coverprofile=cover.out
	@go tool cover -func=cover.out
	@rm -rf cover.out


##@ Development

up: ## Run local environment
	@$(DOCKER_COMPOSE) up -d
	@echo "Running database migration..."
#	@while ! make migrate-status &>/dev/null ; do \
#        echo "Database is not ready, sleep for 3s..." ; \
#        sleep 3 ; \
#    done
	@make migrate-up
	@echo "Migrate database successfully."

down: ## Shutdown local environment
	@$(DOCKER_COMPOSE) down



##@ Migration

fileName := "default"
migrate-create: ## migrate-create fileName=xxx fileType=[sql|go] Creates new migration file with the current timestamp
	@$(GOOSE) create ${fileName} sql

migrate-up: ## Migrate the DB to the most recent version available
	@$(GOOSE) up

migrate-down: ## Roll back the version by 1
	@$(GOOSE) down

migrate-status: ## Dump the migration status for the current DB
	@$(GOOSE) status