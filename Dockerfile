FROM golang

ADD . /go/src/github.com/service-b-go
WORKDIR /go/src/github.com/service-b-go
RUN go build
ENTRYPOINT ./service-b-go

EXPOSE 8080

