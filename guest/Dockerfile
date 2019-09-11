FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/guest
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/guest

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/guest/guest.env /etc/sysconfig/guest

# propagate environment variables to the 'guest' container
RUN cat /etc/sysconfig/guest | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get .

# build project create binary file (./guest/guest)
RUN go build

CMD families

EXPOSE 9092