FROM golang:1.18
WORKDIR /go/src/github.com/hellokvn/go-fiber-api-docker
COPY . .
RUN go build -o bin/server cmd/main.go
CMD ["./bin/server"]