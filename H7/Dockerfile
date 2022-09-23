FROM golang:1.19-alpine

##set direktori utama
WORKDIR /app

##copy file go.mod dan go.sum
COPY go.mod go.sum ./

RUN go mod download

COPY . .

##buat executeable
RUN go build -o main .

EXPOSE 8080

##jalankan executeable
CMD ["/app/main"]