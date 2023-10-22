.DEFAULT_GOAL := build

shadow:
	shadow ./../... # this tool detects shadowing variables
.PHONY:shadow

fmt:
	go fmt ./../...
.PHONY:fmt

lint: fmt
	golint ./../...
.PHONY:lint

vet: fmt
	go vet ./../...
.PHONY:vet

build: vet
	go build -o gocker ../cmd/gocker/main.go
.PHONY:build

run: vet
	go run ../cmd/gocker/main.go
.PHONY:run

# docker commands below

image:
	docker build --platform linux/amd64 -t lifthus/gocker .
.PHONY:image

container:
	docker run -p 0.0.0.0:9595:9595 lifthus/gocker:${TAG}
.PHONY:container

# e.g. make tag TAG=v0.0.1
tag:
	docker tag gocker:latest lifthus/gocker:$(TAG)
.PHONY:tag

push:
	docker push lifthus/gocker:$(TAG)
.PHONY:push