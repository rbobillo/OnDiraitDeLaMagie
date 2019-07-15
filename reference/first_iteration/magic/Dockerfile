FROM golang

COPY . /go/src/magic
WORKDIR /go/src/magic

# propagate environment variables to the 'magic' container
RUN cat /etc/sysconfig/magic | sed -e "s/^/export /g" >> .profile

# get go project external dependencies
RUN go get .

# build project create binary file (./magic/magic)
RUN go build

CMD magic

EXPOSE 9090
