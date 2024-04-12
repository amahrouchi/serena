.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart: down up

.PHONY: logs
logs:
	docker-compose logs -f

.PHONY: db-cli
db-cli:
	docker-compose exec -it postgres psql "postgres://srn_user:testtest@localhost:5432/serena"

.PHONY: test
test:
	docker compose exec app go test -v -cover ./... -coverprofile build/coverage.out

.PHONY: test-coverage
test-coverage:
	go tool cover -html build/coverage.out

.PHONY: test-coverage-func
test-coverage-func:
	docker-compose exec app go tool cover -func build/coverage.out
