FROM golang

# it is okay to leave user/GoDoRP as long as you do not want to share code with other libraries
COPY . /go/src/magic/api
WORKDIR /go/src/magic/api

# added vendor services will need to be included here
#RUN go get ./database

RUN go get ./
RUN go build

CMD api

EXPOSE 8080
