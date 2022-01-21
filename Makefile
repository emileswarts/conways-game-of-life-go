build: stop
	docker-compose build

stop: 
	docker-compose down

test: build
	docker-compose exec -T gol "go test -v ./..."

serve: build
	docker-compose run gol sh -c "sleep infinity"

run: 
	docker-compose up

shell: 
	docker-compose run gol bash

.PHONY: build test
