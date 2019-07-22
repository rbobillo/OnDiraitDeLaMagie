FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/ministry
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/ministry

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/ministry/ministry.env /etc/sysconfig/ministry

# propagate environment variables to the 'ministry' container
RUN cat /etc/sysconfig/ministry | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get github.com/satori/go.uuid
RUN go get github.com/streadway/amqp

# build project create binary file
RUN go build && mv ./ministry /go/bin/ministry

CMD ministry