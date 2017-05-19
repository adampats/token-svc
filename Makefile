SHELL := /bin/bash
TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

.PHONY: all build image run

all: build image run

build: $(TARGET)
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(TARGET) .

image: $(TARGET)
	docker build -t $(TARGET) .

run: $(TARGET)
	docker run -it -p 8080:8080 --rm $(TARGET)
