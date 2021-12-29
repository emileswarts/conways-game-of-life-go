build:
	docker build -t app .

test: build
	docker run -t app go test

.PHONY: build test