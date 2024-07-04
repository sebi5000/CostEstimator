FROM golang:latest

RUN mkdir /app
WORKDIR /app

RUN cd /app && git clone https://github.com/sebi5000/CostEstimator.git
RUN cd /app/CostEstimator && go mod download

RUN mkdir /app/CostEstimator/build
RUN cd /app/CostEstimator && go build -o /app/CostEstimator/build/main cmd/*.go

EXPOSE 8080
ENTRYPOINT [ "/app/CostEstimator/build/main" ]