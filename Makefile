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
	docker system prune -a -f
	#docker buildx create --name multiarch --use || true
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-master:$(IMAGE_TAG) -f Dockerfile-master --push .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-web:$(IMAGE_TAG) -f Dockerfile-web --push .
	docker build -t $(ORG_PREFIX)/$(BINARY_NAME)-client:$(IMAGE_TAG) -f Dockerfile-client --push .

default: tidy build