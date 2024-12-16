FROM golang:1.23.2-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./app/main/main.go

CMD ["./main"]