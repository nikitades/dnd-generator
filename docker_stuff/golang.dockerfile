FROM golang:alpine
RUN apk add git && apk add bash
ENV GOPATH /go
RUN mkdir /go/src/dnd-generator -p
ADD ./ /go/src/dnd-generator
WORKDIR /go/src/dnd-generator/GoApp
EXPOSE 80
RUN go get -u github.com/golang/dep/cmd/dep && dep ensure && go build -o ../main
RUN chmod +x /go/src/dnd-generator/docker_stuff/wait-for-it.sh
CMD ["../docker_stuff/wait-for-it.sh", "db:5432", "--", "/go/src/dnd-generator/main"]