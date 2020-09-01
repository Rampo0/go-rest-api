FROM golang:1.13.5

WORKDIR /go/src/github.com/rampo0/multi-lang-microservice/users

COPY . .

RUN go get -v github.com/codegangsta/gin
RUN go get -d -v ./src/...

RUN go install -v ./src/...

# CMD ["src"]

# Live reload

WORKDIR /go/src/github.com/rampo0/multi-lang-microservice/users/src

CMD ["gin", "--all", "-i", "run", "main.go"]