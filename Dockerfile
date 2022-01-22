FROM golang:1.16

ENV GO111MODULE=on

ADD . /app
WORKDIR /app

RUN go build -o app/web_market_analyzer ./cmd/web_market_analyzer

RUN cd app

CMD ["./app/web_market_analyzer"]