.PHONY: docker.build
docker.build:
	docker build -t router1:latest .

.PHONY: docker.run
docker.run:
	docker rm router1svc
	docker run --rm -d --name router1svc router1 
	
.PHONY: docker.stop
docker.stop:
	docker stop router1svc 

.PHONY: install
install:
	go install ./...

.PHONY: run
run:
	router1