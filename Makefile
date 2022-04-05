server:
	go run cmd/main.go

build:
	go build -o bin/server cmd/main.go

d.up:
	docker-compose --env-file pkg/common/config/envs/.env up 

d.down:
	docker-compose down

d.up.build:
	docker-compose --env-file pkg/common/config/envs/.env --build up