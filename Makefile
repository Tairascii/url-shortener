BIN_DIR=${CURDIR}/bin
MIGRATION_DIR=db/migrations

bin-deps:
	$(info Installing binary dependencies...)
	GOBIN=$(BIN_DIR) go install github.com/pressly/goose/v3/cmd/goose@v3.24.1
	GOBIN=$(BIN_DIR) go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.28.0

create-migration:
	$(BIN_DIR)/goose -dir db/migrations create -s $(name) sql

compile-sql:
	$(BIN_DIR)/sqlc generate
