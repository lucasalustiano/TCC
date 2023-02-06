FROM golang:latest

WORKDIR /app

COPY . .

RUN go get -d -v ./...

RUN go build ./api/main.go

EXPOSE 8080

CMD ["./main"]