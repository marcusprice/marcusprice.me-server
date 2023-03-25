dev:
	air main.go

build:
	go build main.go

create-curl-card:
	bash ./bin/business-card-template.sh

docker-build:
	docker build -t marcusprice.me-server .

docker-run:
	docker run -it --rm --name running-marcusprice.me-server marcusprice.me-server
