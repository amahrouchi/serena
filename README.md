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
```
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
postgresql://srn_user:testtest@localhost:5432/postgres
```
