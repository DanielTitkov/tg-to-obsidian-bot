NAME:=bot
BUILD_CMD ?= CGO_ENABLED=0 go build -o bin/${NAME} -ldflags '-v -w -s' ./main.go

SHELL = /bin/sh
CURRENT_UID := $(shell id -u)
CURRENT_GID := $(shell id -g)

include .env
export

.PHONY: run
run:
	go run main.go \
		-timeout=30 \
		-token=$(TG_TOKEN) \
		-template=$(TEMPLATE_PATH) \
		-path=$(NOTES_PATH)

.PHONY: build
build:
	echo "building"
	${BUILD_CMD}
	echo "build done"

.PHONY: up
up:
	UID="${CURRENT_UID}" GID="${CURRENT_GID}" docker-compose up --build -d

# not to fail if .env is not present
.env:
	touch $@