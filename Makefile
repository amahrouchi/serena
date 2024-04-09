up:
	docker-compose up -d

down:
	docker-compose down

restart:
	down up

logs:
	docker-compose logs -f

test:
	docker compose exec app go test -v -cover ./...
