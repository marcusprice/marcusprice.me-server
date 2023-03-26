dev:
	air main.go

tmux-session:
	bash ./scripts/tmux-session.sh

build:
	go build main.go

create-curl-card:
	bash ./bin/business-card-template.sh

docker-build:
	docker build -t marcusprice.me-server .

docker-run:
	docker run -it -p 6969:6969 --rm --name running-marcusprice.me-server marcusprice.me-server
