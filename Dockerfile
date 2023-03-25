FROM golang:1.20

WORKDIR /usr/src/marcusprice.me-backend

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go build -v -o usr/local/bin/marcusprice.me-backend ./...

CMD ["marcusprice.me-backend"]
