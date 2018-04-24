.PHONY: container export

run:
	python2 -m SimpleHTTPServer 8000

build:
	GOOS=linux go build

container:
	docker build -t ubuntu:bvale -f docker/ubuntu/Dockerfile ./docker/ubuntu

export:
	docker export $(docker create ubuntu:bvale) -o ubuntu.tar
