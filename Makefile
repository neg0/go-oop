.PHONY: all
build:
	docker-compose -f docker/docker-compose.yml build

build-force:
	docker-compose -f docker/docker-compose.yml build --no-cache

up:
	docker-compose -f docker/docker-compose.yml up -d

down:
	docker-compose -f docker/docker-compose.yml down

ssh:
	docker-compose -f docker/docker-compose.yml exec golang bash