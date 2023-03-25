FROM golang:1.20

WORKDIR /usr/src/marcusprice.me-server

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN make create-curl-card

RUN go build -v -o /usr/local/bin/marcusprice.me-server ./...

EXPOSE 6969

CMD ["marcusprice.me-server"]
