# Serena

Serena is a simple blockchain implementation in Go.

## Requirements

- Go 1.22
- Docker
- Docker Compose
- Make

## How to run the app

### Configuration
Copy the `.env.template` file to `.env` and fill it with the required values. It is already
filled with the default values for a local environment.

### Run the app
Build and start the app with the following command:
```bash
make up
```

### Health check
```
GET http://localhost:8080/healthz
```

### Write to the blockchain
```
POST http://localhost:8080/write
Accept: application/json
Content-Type: application/json
{
    "author": "Angelo M.",
    "data": {
        "transaction": {
            "uuid": "b1b9b1b9-1b9b-1b9b-1b9b-1b9b1b9b1b9b",
            "amount": 100.0,
            "currency": "USD",
            "status": "approved"
        }
    }
}
```
**N.B.:** The data field can be any JSON object.

### Storage
The blocks are stored in a PostgreSQL database running in a Docker container.
```
postgresql://srn_user:testtest@localhost:5432/serena
```

## Tests

### Database
Some functional and integration tests need a specific database to be created. So, before running the tests, please connect to the Postgres server and run the following SQL queries:
```sql
DROP DATABASE IF EXISTS serena_test;
CREATE DATABASE serena_test OWNER srn_user;
```

### Run the tests
The whole suite can be run with the following command:
```bash
make test
```
This command will run the test and then display a coverage report telling you which functions in the project are not covered properly (80% min).
