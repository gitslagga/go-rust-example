FROM golang

MAINTAINER Slagga "shuaijun"

WORKDIR $GOPATH/src/godocker

ADD . $GOPATH/src/godocker

RUN go build main.go

EXPOSE 8080

ENTRYPOINT ["./main"]