FROM golang:1.22-alpine

RUN apk update && apk upgrade && \
    apk add --no-cache make

RUN go install github.com/air-verse/air@latest

WORKDIR /app

CMD ["air", "-c", ".air.toml"]
