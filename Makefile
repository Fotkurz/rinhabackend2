

all: 
	- docker compose -f deployment/docker-compose.yml down
	- docker compose -f deployment/docker-compose.yml up

build:
	docker build -t rinhabackend2 . -f deployment/Dockerfile

