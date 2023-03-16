dev:
	air main.go

build:
	go build main.go

create-curl-card:
	bash ./bin/business-card-template.sh
