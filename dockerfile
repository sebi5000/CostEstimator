FROM golang:1.22
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN mkdir /usr/src/app/output
RUN go build -v -o ./output ./...

CMD ["app"]