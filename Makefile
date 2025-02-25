.PHONY: tidy
tidy:
	go mod tidy -v

.PHONY: migrate-create
migrate-create:
	./scripts/migrate.sh create $(name)

.PHONY: migrate-up
migrate-up:
	./scripts/migrate.sh up $(steps)

.PHONY: migrate-down
migrate-down:
	./scripts/migrate.sh down $(steps)

.PHONY: migrate-version
migrate-version:
	./scripts/migrate.sh version