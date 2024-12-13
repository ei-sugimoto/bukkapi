.PHONY: dev
dev:
	@echo "runnning docker compose dev..."
	@docker compose -f docker-compose.dev.yml up

.PHONY: gen
gen:
	@echo "gen code..."
	@go generate ./...