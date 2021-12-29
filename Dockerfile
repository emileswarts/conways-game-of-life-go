FROM golang:1.17.5-alpine3.15

WORKDIR /app
COPY src /app/src

CMD ["go", "run", "/app/src/main.go"]