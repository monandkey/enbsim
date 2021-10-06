FROM golang:1.16.1
LABEL maintainer="monandkey"
RUN git clone https://github.com/ishidawataru/sctp.git

WORKDIR /go/src/example
RUN go build
