FROM golang:1.19-alpine

RUN mkdir /app

WORKDIR /app

ADD ../day7/. /app

RUN go build -o main .

CMD ["/app/main"]