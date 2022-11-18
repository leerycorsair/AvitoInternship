run:
	docker-compose down -v
	service postgre stop
	docker build --no-cache --network host -f ./docker/Dockerfile . --tag application
	docker-compose up --build

build:
	go build -o server.out -v ./cmd/server/main.go