.PHONY: docker.build
docker.build:
	docker build -t proxy:latest .

.PHONY: docker.run
docker.run:
	docker rm proxysvc
	docker run --rm -d --name proxysvc proxy 
	
.PHONY: docker.stop
docker.stop:
	docker stop proxysvc 

.PHONY: install
install:
	go install ./...

.PHONY: run
run:
	proxy