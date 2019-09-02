FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/villains
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/villains

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/villains/villains.env /etc/sysconfig/villains

# propagate environment variables to the 'magic' container
RUN cat /etc/sysconfig/villains | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get .

# build project create binary file (./magic/magic)
RUN go build

CMD magic

EXPOSE 9090
