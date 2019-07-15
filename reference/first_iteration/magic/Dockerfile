FROM golang

COPY . /go/src/magic
WORKDIR /go/src/magic

RUN go get .
RUN go build

CMD api

EXPOSE 9090
