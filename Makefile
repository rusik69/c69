SHELL := /bin/bash
.DEFAULT_GOAL := default
.PHONY: all

BINARY_NAME=govnocloud
IMAGE_TAG=$(shell git describe --tags --always)
GIT_COMMIT=$(shell git rev-parse --short HEAD)
ORG_PREFIX := loqutus

tidy:
	go mod tidy

build:
	CGO_ENABLED=1 GOARCH=amd64 GOOS=linux go build -ldflags "-X main.version=$(GIT_COMMIT)" -o bin/${BINARY_NAME}-linux-amd64 main.go
	chmod +x bin/*

docker:
	set -x
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) -f Dockerfile-master .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG) -f Dockerfile-web .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) -f Dockerfile-client .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) -f Dockerfile-node .
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-master:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-web:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-client:latest
	docker tag $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG) $(ORG_PREFIX)/$(BINARY_NAME)-node:latest
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG)
	docker push $(ORG_PREFIX)/$(BINARY_NAME)-node:$(IMAGE_TAG)

compose:
	docker-compose up

default: tidy build