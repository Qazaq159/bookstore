FROM golang:alpine

MAINTAINER yerke

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/bookapp

COPY . /go/src/bookapp

RUN go get .

CMD ["go", "run", "."]