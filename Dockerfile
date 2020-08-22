FROM golang:latest

WORKDIR /go/src/multi-lang-microservice/users

COPY . .

RUN go get -v github.com/codegangsta/gin
RUN go get -d -v ./src/...

RUN go install -v ./src/...

# CMD ["src"]

# Live reload

WORKDIR /go/src/multi-lang-microservice/users/src

CMD ["gin", "--all", "-i", "run", "main.go"]