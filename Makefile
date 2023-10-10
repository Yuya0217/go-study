DOCKER_COMPOSE := docker-compose -f ./docker-compose.yml -p go-layered-architecture-sample
TARGET := go-layered-architecture-sample
BINARY_DIR := ./bin
MIGRATE_VERSION := 4.16.2
MIGRATE=$(BINARY_DIR)/migrate
DB_MIGRATIONS_PATH=./db/migrations
GOLANGCI_LINT_VERSION := 1.54.2
GOLANGCI_LINT= ./bin/golangci-lint
GOFILES=$(shell find . -type f -name '*.go')
GOFMT=gofmt
DATABASE_URL='mysql://$(DATABASE_USER):$(DATABASE_PASSWORD)@tcp($(DATABASE_PRIMARY_HOST):$(DATABASE_PRIMARY_PORT))/$(DATABASE_NAME)'

.PHONY: tool-download
tool-download: ## 必要なツール類をダウンロード
	curl -L https://github.com/golang-migrate/migrate/releases/download/v$(MIGRATE_VERSION)/migrate.darwin-arm64.tar.gz | tar xz -C ./bin migrate
	curl -L https://github.com/golangci/golangci-lint/releases/download/v$(GOLANGCI_LINT_VERSION)/golangci-lint-$(GOLANGCI_LINT_VERSION)-darwin-arm64.tar.gz | tar -xz --strip=1 -C ./bin golangci-lint-$(GOLANGCI_LINT_VERSION)-darwin-arm64/golangci-lint

.PHONY: test
test: ## testを実行します
	go test ./...

.PHONY: lint
lint: ## コードのリントを実行
	${GOLANGCI_LINT} run

.PHONY: fmt
fmt: ## コードフォーマット
	$(GOFMT) -w $(GOFILES)

.PHONY: build
build: ## アプリケーションのビルド
	go build -o "$(TARGET)" "./cmd/$(TARGET)"

.PHONY: run
run: ## アプリケーションの実行
	${MAKE} build
	go run ./$(TARGET)

.PHONY: ci
ci: lint test build ## CIで実行されるタスクを実行

.PHONY: up
up: ## docker-compose up でサービスを開始
	$(DOCKER_COMPOSE) up --build -d database

.PHONY: down
down: ## docker-compose down でサービスを停止・削除
	$(DOCKER_COMPOSE) down

.PHONY: golden-update
golden-update: ## goldenファイルの生成
	go test ./internal/presentation/rest/handler/... -update-golden

.PHONY: generate
generate: ## モックやgoldenファイルを生成　
	openapi-generator generate -i ./docs/openapi.yaml -g go -o ./generated --global-property models,supportingFiles="utils.go",modelDocs=false --type-mappings integer=int
	go generate ./...
	${MAKE} golden-update

.PHONY: migrate-up
migrate-up: ## マイグレーションを適用
	$(MIGRATE) -path=$(DB_MIGRATIONS_PATH) -database=$(DATABASE_URL) up

.PHONY: migrate-down
migrate-down: # マイグレーションを取消
	$(MIGRATE) -path=$(DB_MIGRATIONS_PATH) -database=$(DATABASE_URL) down

.PHONY: migrate-status
migrate-status: ## 最新のマイグレーションの状態を確認
	$(MIGRATE) -path=$(DB_MIGRATIONS_PATH) -database=$(DATABASE_URL) version

.PHONY: create-migration
create-migration: ## マイグレーションファイルの生成
	@if [ "$(name)" = "" ]; then \
		echo "Error: マイグレーションの名前を指定してください (e.g., make create-migration name=create_users_table)"; \
		exit 1; \
	fi
	$(MIGRATE) create -ext sql -dir $(DB_MIGRATIONS_PATH) -seq $(name)

help:
	@echo "利用可能なコマンド:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
