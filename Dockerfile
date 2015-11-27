FROM golang:latest
RUN mkdir /app
#ADD . /app/
#WORKDIR /app
RUN echo $GOPATH
RUN echo ls $GOPATH/src
RUN mkdir -p  $GOPATH/src/github.com/d-smith/docker-build
WORKDIR $GOPATH/src/github.com/d-smith/docker-build
ADD . $GOPATH/src/github.com/d-smith/docker-build
RUN go get -d -v
RUN go build -o main .
RUN cp main /app
CMD ["/app/main"]