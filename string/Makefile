.PHONY: docker.build
docker.build:
	docker build -t string:latest .

.PHONY: docker.run
docker.run:
	docker rm stringsvc
	docker run --rm -d --name stringsvc string 
	
.PHONY: docker.stop
docker.stop:
	docker stop stringsvc 

.PHONY: install
install:
	go install ./...

.PHONY: run
run:
	string