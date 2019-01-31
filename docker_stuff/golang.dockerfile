FROM golang:alpine
RUN apk add git
RUN mkdir /go/bin -p
ADD ./ /go
WORKDIR /go/GoApp
RUN go get -u github.com/golang/dep/cmd/dep && \
    dep ensure && \
    go build -o ../bin/main
CMD ["/go/bin/main"]