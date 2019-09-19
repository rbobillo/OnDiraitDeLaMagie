FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/azkaban
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/azkaban

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/azkaban/azkaban.env /etc/sysconfig/azkaban

# propagate environment variables to the 'azkaban' container
RUN cat /etc/sysconfig/azkaban | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get github.com/satori/go.uuid
RUN go get github.com/streadway/amqp
# build project create binary file
RUN go build && mv ./azkaban /go/bin/azkaban

CMD azkaban