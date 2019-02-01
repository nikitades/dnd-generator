FROM golang:alpine
RUN apk add git
ENV GOPATH /go
RUN mkdir /go/src/dnd-generator -p
ADD ./ /go/src/dnd-generator
WORKDIR /go/src/dnd-generator/GoApp
EXPOSE 80
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure && go build -o ../main
CMD ["/go/src/dnd-generator/main"]