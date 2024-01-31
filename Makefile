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
	scp deployments/docker-compose-master.yml govnocloud-master.rusik69.lol:~/
	ssh govnocloud-master.rusik69.lol "docker compose -f docker-compose-master.yml down"
	ssh govnocloud-master.rusik69.lol "docker compose -f docker-compose-master.yml up -d --quiet-pull"
	scp deployments/docker-compose-x220.yml x220.rusik69.lol:~/
	ssh x220.rusik69.lol "docker compose -f docker-compose-x220.yml down"
	ssh x220.rusik69.lol "docker system prune -a -f"
	ssh x220.rusik69.lol "sudo virsh destroy test; sudo virsh undefine test" || true
	ssh x220.rusik69.lol "echo {0..10} | xargs -I {} sudo virsh destroy test{} && echo {0..10} | xargs -I {} sudo virsh undefine test{}"	ssh x220.rusik69.lol "docker ps -aq | xargs docker stop | xargs docker rm" || true
	ssh x220.rusik69.lol "docker compose -f docker-compose-x220.yml up -d --quiet-pull"
	scp deployments/docker-compose-x230.yml x230.rusik69.lol:~/
	ssh x230.rusik69.lol "docker compose -f docker-compose-x230.yml down"
	ssh x230.rusik69.lol "docker system prune -a -f"
	ssh x230.rusik69.lol "sudo virsh destroy test; sudo virsh undefine test" || true
	ssh x230.rusik69.lol "echo {0..10} | xargs -I {} sudo virsh destroy test{} && echo {0..10} | xargs -I {} sudo virsh undefine test{}"	ssh x220.rusik69.lol "docker ps -aq | xargs docker stop | xargs docker rm" || true
	ssh x230.rusik69.lol "docker ps -aq | xargs docker stop | xargs docker rm" || true
	ssh x230.rusik69.lol "docker compose -f docker-compose-x230.yml up -d --quiet-pull"
	sleep 10

ansible:
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

