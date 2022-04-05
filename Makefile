server:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

d.up:
	docker-compose up 

d.down:
	docker-compose down

d.up.build:
	docker-compose --build up