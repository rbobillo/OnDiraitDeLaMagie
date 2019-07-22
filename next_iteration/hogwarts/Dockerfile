FROM golang

COPY . /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts
WORKDIR /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts

RUN mkdir -p /etc/sysconfig/ && cp /go/src/github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/hogwarts.env /etc/sysconfig/hogwarts

# propagate environment variables to the 'hogwarts' container
RUN cat /etc/sysconfig/hogwarts | sed -e "s/^/export /g" >> ~/.bashrc

# get go project external dependencies
RUN go get .

# build project create binary file
RUN go build

CMD hogwarts

EXPOSE 9091