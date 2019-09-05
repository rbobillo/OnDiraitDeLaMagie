FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/families
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/families

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/families/families.env /etc/sysconfig/families

# propagate environment variables to the 'magic' container
RUN cat /etc/sysconfig/families | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get .

# build project create binary file (./magic/magic)
RUN go build

CMD families

EXPOSE 9092