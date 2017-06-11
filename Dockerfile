FROM golang

ADD . /go/src/github.com/pdu/presents

RUN go get golang.org/x/tools/cmd/present

WORKDIR /go/src/github.com/pdu/presents

ENTRYPOINT ["/go/bin/present"]
CMD []

EXPOSE 3999
