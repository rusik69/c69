SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all build

BINARY_NAME=govnocloud
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

tidy:
	go mod tidy

get:
	go get -v ./...

build:
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-linux-amd64 main.go
	chmod +x bin/*

test:
	go test -v ./...

docker:
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) -f build/Dockerfile-master .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG) -f build/Dockerfile-web .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) -f build/Dockerfile-client .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) -f build/Dockerfile-node .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) -f build/Dockerfile-test .
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-master:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-web:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-client:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-node:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-test:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-master:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-web:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-client:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-node:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-test:latest

deploy:
	ssh master "docker compose -f docker-compose-master.yml down"
	ssh master "docker system prune -a -f"
	scp deployments/docker-compose-master.yml master:~/
	ssh master "docker pull $(ORG_PREFIX)/$(BINARY_NAME)-master"
	ssh master "docker compose -f docker-compose-master.yml up -d"
	scp deployments/docker-compose-node0.yml node0:~/
	ssh node0 "docker compose -f docker-compose-node0.yml down"
	ssh node0 "docker system prune -a -f"
	ssh node0 "docker compose -f docker-compose-node0.yml up -d"
	scp deployments/docker-compose-node1.yml node1:~/
	ssh node1 "docker compose -f docker-compose-node1.yml down"
	ssh node1 "docker system prune -a -f"
	ssh node1 "docker compose -f docker-compose-node1.yml up -d"
	sleep 10

prune:
	docker system prune -a -f

ansible:
	ansible-playbook -i deployments/ansible/inventories/testing/hosts deployments/ansible/main.yml

composetest:
	docker compose -f deployments/docker-compose-test.yml up --abort-on-container-exit --exit-code-from test

composelogs:
	ssh master "docker compose -f docker-compose-master.yml logs"
	ssh node0 "docker compose -f docker-compose-node.yml logs"
	ssh node1 "docker compose -f docker-compose-node.yml logs"

default: get build
