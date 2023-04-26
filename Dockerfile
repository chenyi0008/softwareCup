FROM golang:1.20.1
MAINTAINER ChenYi
WORKDIR /go/src/docker_go
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

CMD ["docker_go"]

