FROM golang:1.23.3

WORKDIR /go/src/app

COPY  . .

RUN go build -o main main.go

EXPOSE 4040

CMD [ "./main" ]