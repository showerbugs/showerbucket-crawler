from golang

ADD . /app

RUN go get "github.com/badoux/goscraper"

WORKDIR /app

ENTRYPOINT go run main.go

EXPOSE 9090
