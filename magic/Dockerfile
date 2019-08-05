FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/magic
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/magic

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/magic/magic.env /etc/sysconfig/magic

# propagate environment variables to the 'magic' container
RUN cat /etc/sysconfig/magic | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get .

# build project create binary file (./magic/magic)
RUN go build

CMD magic

EXPOSE 9090
