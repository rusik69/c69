SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all build

BINARY_NAME=govnocloud
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

export TEST_MASTER_HOST := t440p.rusik69.lol
export TEST_MASTER_PORT := 7070
export TEST_NODE_NAME := x220
export TEST_NODE_HOST := x220.rusik69.lol
export TEST_NODE_PORT := 6969
export TEST_NODES := x220.rusik69.lol:6969,x230.rusik69.lol:6969

tidy:
	go mod tidy

get:
	go get -v ./...

build:
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-linux-amd64 main.go
	chmod +x bin/*

test:
	
	go test -v ./...

deploy:
	sudo systemctl stop govnocloud-master
	ETCDCTL_API=3 etcdctl del --prefix ""
	ssh x220.rusik69.lol "sudo systemctl stop govnocloud-node" || true
	ssh x230.rusik69.lol "sudo systemctl stop govnocloud-node" || true
	ssh x220.rusik69.lol "/usr/local/bin/cleanup.sh" || true
	ssh x230.rusik69.lol "/usr/local/bin/cleanup.sh" || true
	sudo cp bin/${BINARY_NAME}-linux-amd64 /usr/local/bin/govnocloud
	scp bin/${BINARY_NAME}-linux-amd64 root@x220.rusik69.lol:/usr/local/bin/govnocloud
	scp bin/${BINARY_NAME}-linux-amd64 root@x230.rusik69.lol:/usr/local/bin/govnocloud
	rsync deployments/ansible root@x220.rusik69.lol:/var/lib/libvirt/
	rsync deployments/ansible root@x230.rusik69.lol:/var/lib/libvirt/
	sudo systemctl start govnocloud-master
	ssh x220.rusik69.lol "sudo systemctl start govnocloud-node"
	ssh x230.rusik69.lol "sudo systemctl start govnocloud-node"

docker:
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME):$(IMAGE_TAG) -f build/Dockerfile .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) -f build/Dockerfile-test .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG) -f build/Dockerfile-front .
	docker tag $(ORG_PREFIX)/$(BINARY_NAME):$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME):latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-test:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-front:latest
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME):$(IMAGE_TAG)
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME)-test:$(IMAGE_TAG)
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME)-front:$(IMAGE_TAG)
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME):latest
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME)-test:latest
	docker push -q $(ORG_PREFIX)/$(BINARY_NAME)-front:latest

dockerdeploy:
	scp deployments/docker-compose-master.yml ~/
	docker compose -f ~/docker-compose-master.yml down
	docker compose -f ~/docker-compose-master.yml up -d --quiet-pull
	scp deployments/docker-compose-x220.yml x220.rusik69.lol:~/
	ssh x220.rusik69.lol "/usr/local/bin/cleanup.sh"
	ssh x220.rusik69.lol "docker compose -f ~/docker-compose-x220.yml up -d --quiet-pull"
	scp deployments/docker-compose-x230.yml x230.rusik69.lol:~/
	ssh x230.rusik69.lol "/usr/local/bin/cleanup.sh"
	ssh x230.rusik69.lol "docker compose -f ~/docker-compose-x230.yml up -d --quiet-pull"

ansible:
	ansible-playbook -i deployments/ansible/inventories/testing/hosts deployments/ansible/main.yml

composetest:
	docker compose -f deployments/docker-compose-test.yml up --abort-on-container-exit --exit-code-from test --quiet-pull

composelogs:
	ssh t440p.rusik69.lol "docker compose -f docker-compose-master.yml logs"
	ssh x220.rusik69.lol "docker compose -f docker-compose-x220.yml logs"
	ssh x230.rusik69.lol "docker compose -f docker-compose-x230.yml logs"

logs:
	journalctl _SYSTEMD_INVOCATION_ID=`systemctl show -p InvocationID --value govnocloud-master.service`
	ssh x220.rusik69.lol "get_logs.sh"
	ssh x230.rusik69.lol "get_logs.sh"

remotetest:
	rsync -avz . t440p.rusik69.lol:~/govnocloud
	ssh t440p.rusik69.lol "cd govnocloud; make docker; make deploy; make composetest; make composelogs"

rsync:
	rsync -avz . t440p.rusik69.lol:~/govnocloud

default: get build

