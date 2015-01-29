DEFAULT: all

deps:
	go get

build: deps
	go build

install: deps
	go install

all: install
