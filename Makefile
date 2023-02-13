MODULE = $(shell go list -m)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || echo "1.0.0")
PACKAGES := $(shell go list ./... | grep -v /vendor/)
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
PROTO_CMD := protoc -I./api -I./api/proto/grpc-gateway/third_party/googleapis -I./api/proto/grpc-gateway-v2 --go_out ./api/proto --go-grpc_out ./api/proto --grpc-gateway_out ./api/proto --openapiv2_out ./api/ --openapiv2_opt allow_merge=true,merge_file_name=api ./api/proto/*.proto
PROTO_EXT_CMD := protoc -I./api -I./api/proto/grpc-gateway/third_party/googleapis -I./api/proto/grpc-gateway-v2 --go_out ./api/proto/external --go-grpc_out ./api/proto/external ./api/proto/external/*.proto

#ifndef APP_ENV
#	ifneq (,$(wildcard .env))
#		include .env
#		export
#		. .env
#	endif
#endif

APP_DSN=postgres://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_DATABASE}?sslmode=disable
MIGRATE := migrate -path=./migrations/ -database "$(APP_DSN)"

PID_FILE := './.pid'
FSWATCH_FILE := './fswatch.cfg'

.PHONY: default
default: help

# generate help info from comments: thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help: ## help information about make commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test
test: ## run unit tests
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: test-cover
test-cover: test ## run unit tests and show test coverage information
	go tool cover -html=coverage-all.out

.PHONY: run
run: ## run the API server
	go run ${LDFLAGS} cmd/loto/main.go

.PHONY: run-restart
run-restart: ## restart the API server
	@pkill -P `cat $(PID_FILE)` || true
	@printf '%*s\n' "80" '' | tr ' ' -
	@echo "Source file changed. Restarting server..."
	@go run ${LDFLAGS} cmd/loto/main.go & echo $$! > $(PID_FILE)
	@printf '%*s\n' "80" '' | tr ' ' -

.PHONY: run-live
run-live: ## run the API server with live reload support (requires fswatch)
	@go run ${LDFLAGS} cmd/loto/main.go & echo $$! > $(PID_FILE)
	@fswatch -x -o --event Created --event Updated --event Renamed -r internal pkg cmd config | xargs -n1 -I {} make run-restart

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o /loto -v ./cmd/loto

.PHONY: docker-build
docker-build: ## build the API server as a docker image
	docker-compose -f deployments/docker-compose.yml -p loto build --no-cache

.PHONY: docker-up
docker-up: ## run the API server as a docker image
	docker-compose -f deployments/docker-compose.yml -p loto up -d

.PHONY: docker-stop
docker-stop: ## stop the API server
	docker-compose -f deployments/docker-compose.yml -p loto stop

.PHONY: docker-down
docker-down: ## remove docker images
	docker-compose -f deployments/docker-compose.yml -p loto down

.PHONY: docker-rebuild
docker-rebuild: ## re-build the API server as a docker image
	docker-compose -f deployments/docker-compose.yml -p loto down && docker-compose -f deployments/docker-compose.yml -p loto up -d --build

.PHONY: app-shell
app-shell: ## bash into app container
	docker-compose -f deployments/docker-compose.yml -p loto exec loto bash

.PHONY: clean
clean: ## remove temporary files
	rm -rf loto coverage.out coverage-all.out

.PHONY: version
version: ## display the version of the API server
	@echo $(VERSION)

.PHONY: db-start
db-start: ## start the database server
	@mkdir -p testdata/postgres
	docker run --rm --name db -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=loto -e POSTGRES_DB=loto -d -p 5432:5432 loto

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop db

.PHONY: testdata
testdata: ## populate the database with test data
	make migrate-reset
	@echo "Populating test data..."
	@docker exec -it db psql "$(APP_DSN)" -f /testdata/testdata.sql

.PHONY: lint
lint: ## run golint on all Go package
	@golint $(PACKAGES)

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1

.PHONY: migrate-new
migrate-new: ## create a new database migration
	@read -p "Enter the name of the new migration: " name; \
	$(MIGRATE) create -ext sql -dir ./migrations/ $${name}
#	$(MIGRATE) create -ext sql -dir /migrations/ $${name// /_}

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: swagger
swagger: ## build swagger.yml file
	swagger generate spec -o ./api/swagger.json --scan-models

.PHONY: proto
proto: ## Generate proto
ifneq (,$(wildcard api/proto/external/*.proto))
	@$(PROTO_CMD)
	@$(PROTO_EXT_CMD)
else
	@$(PROTO_CMD)
endif