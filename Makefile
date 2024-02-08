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
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) -f build/Dockerfile-client .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) -f build/Dockerfile-node .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) -f build/Dockerfile-test .
	docker build --progress=plain -t $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG) -f build/Dockerfile-front .
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-master:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-client:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-node:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-test:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-front:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-master:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-client:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-node:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-test:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-front:latest

deploy:
	scp deployments/docker-compose-master.yml ~/
	docker compose -f ~/docker-compose-master.yml down
	docker compose -f ~/docker-compose-master.yml up -d --quiet-pull
	docker system prune -a -f
	scp deployments/docker-compose-x220.yml x220.rusik69.lol:~/
	ssh x220.rusik69.lol "/usr/local/bin/cleanup.sh"
	ssh x220.rusik69.lol "docker compose -f ~/docker-compose-x220.yml up -d --quiet-pull"
	scp deployments/docker-compose-x230.yml x230.rusik69.lol:~/
	ssh x230.rusik69.lol "/usr/local/bin/cleanup.sh"
	ssh x230.rusik69.lol "docker compose -f ~/docker-compose-x230.yml up -d --quiet-pull"

ansible:
	sudo dnf install python3-devel libxml2-devel libxslt-devel redhat-rpm-config gcc
	pip3 install -r deployments/ansible/requirements.txt
	ansible-playbook -i deployments/ansible/inventories/testing/hosts deployments/ansible/main.yml

composetest:
	docker compose -f deployments/docker-compose-test.yml up --abort-on-container-exit --exit-code-from test --quiet-pull

composelogs:
	ssh govnocloud-master.rusik69.lol "docker compose -f docker-compose-master.yml logs"
	ssh x220.rusik69.lol "docker compose -f docker-compose-x220.yml logs"
	ssh x230.rusik69.lol "docker compose -f docker-compose-x230.yml logs"

remotetest:
	rsync -avz . govnocloud-master.rusik69.lol:~/govnocloud
	ssh govnocloud-master.rusik69.lol "cd govnocloud; make docker; make deploy; make composetest; make composelogs"

default: get build

