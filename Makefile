build-docker-kafka:
	docker build -t kafka-kaste
run-up-docker-compose:
	docker-compose up -t
exec-kafka-kaste-container:
	docker exec -it kafka-kaste sh

tests:
	go test -v ./...
build:
	go buil -v ./...

.PHONY: build-docker-kafka run-up-docker-compose exec-kafka-kaste-container tests build