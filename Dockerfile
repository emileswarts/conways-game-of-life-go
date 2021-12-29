FROM golang:1.17.5-alpine3.15

WORKDIR /app
COPY src/ /app/

RUN apk update && apk add build-base
RUN go get github.com/stretchr/testify/assert
RUN go build -o welcome .

CMD ["./welcome"]