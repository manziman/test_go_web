FROM golang

WORKDIR /go/src

RUN go get github.com/manziman/testWeb

COPY . .

RUN go install github.com/manziman/test_go_web

EXPOSE 9090

WORKDIR /go/bin
CMD ["test_go_web"]
