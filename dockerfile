FROM golang:1.20-alpine

WORKDIR /cmd
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /costestimator

EXPOSE 8080

CMD ["/costestimator"]