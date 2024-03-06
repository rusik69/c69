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
	GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-deploy-darwin-arm64 cmd/deploy/main.go
	GOARCH=arm64 GOOS=darwin go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-client-darwin-arm64 cmd/client/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-client-linux-amd64 cmd/client/main.go
	GOARCH=arm64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-client-linux-arm64 cmd/client/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-master-linux-amd64 cmd/master/main.go
	GOARCH=arm64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-master-linux-arm64 cmd/master/main.go
	GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-node-linux-amd64 cmd/node/main.go
	chmod +x bin/*

test:
	go test -timeout 30m -v ./...

deploy:
	/bin/govnocloud-deploy-linux-amd64 --master t440p.rusik69.lol --nodes x220.rusik69.lol,x230.rusik69.lol

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
	ssh t440p.rusik69.lol "cd govnocloud; make ansible get build deploy test logs"

rsync:
	rsync -avz . t440p.rusik69.lol:~/govnocloud

default: get build

