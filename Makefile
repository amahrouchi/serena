.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: restart
restart: down up

.PHONY: shell
shell:
	docker-compose exec app sh

.PHONY: logs
logs:
	docker-compose logs -f app

.PHONY: db-cli
db-cli:
	docker-compose exec -it postgres psql "postgres://srn_user:testtest@localhost:5432/serena"

.PHONY: test
test:
	docker compose exec -e SRN_ENV=test app go test ./... -v -coverprofile=build/coverage.out -p 1

.PHONY: test-coverage
test-coverage:
	go tool cover -html=build/coverage.out

.PHONY: test-coverage-func
test-coverage-func:
	docker-compose exec app go tool cover -func build/coverage.out

.PHONY: ps
ps:
	docker-compose ps
