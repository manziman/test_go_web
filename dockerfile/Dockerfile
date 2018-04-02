FROM golang

WORKDIR /go/src

RUN go get https://github.com/manziman/testWeb.git

COPY . .

RUN go install testWeb

EXPOSE 9090

WORKDIR /go/bin
CMD ["testWeb"]
