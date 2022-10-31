up-env: ## Set up env
	cd ./order-service && make up

run-order: 		## Run order service
	cd ./order-service && CONFIG_PATH=config/local.yaml go run main.go

run-pay: 		## Run pay service
	cd ./pay-service && CONFIG_PATH=config/local.yaml go run main.go

run-admin: 		## Run admin page
	cd ./admin/admin-panel && npm install && npm start