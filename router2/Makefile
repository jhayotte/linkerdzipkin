.PHONY: docker.build
docker.build:
	docker build -t router2:latest .

.PHONY: docker.run
docker.run:
	docker rm router2svc
	docker run --rm -d --name router2svc router2 
	
.PHONY: docker.stop
docker.stop:
	docker stop router2svc 

.PHONY: install
install:
	go install ./...

.PHONY: run
run:
	router2