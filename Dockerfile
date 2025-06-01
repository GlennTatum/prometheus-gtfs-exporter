FROM golang:1.24

WORKDIR /app

COPY . /app/

RUN go build -o run -v cli/main.go

EXPOSE 9091

CMD ["/app/run"]