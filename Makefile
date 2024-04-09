up:
	docker-compose up -d

down:
	docker-compose down

restart:
	down up

logs:
	docker-compose logs -f

test:
	docker compose exec app go test -v -cover ./... -coverprofile build/coverage.out

test-coverage:
	go tool cover -html build/coverage.out

test-coverage-func:
	docker-compose exec app go tool cover -func build/coverage.out
