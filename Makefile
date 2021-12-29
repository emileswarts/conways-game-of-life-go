build:
	docker build -t app .

test: build
	docker run -t app go test -v

run: build
	docker run -t app 

shell: 
	docker run -t app /bin/bash

.PHONY: build test