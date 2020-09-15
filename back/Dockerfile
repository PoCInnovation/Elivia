FROM golang:1.12-alpine

RUN apk add git gcc libc-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ./bin .

CMD ["./bin"]