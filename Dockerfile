FROM golang:latest


RUN mkdir /app


RUN echo $GOPATH
RUN echo ls $GOPATH/src
RUN mkdir -p  $GOPATH/src/github.com/d-smith/docker-build
WORKDIR $GOPATH/src/github.com/d-smith/docker-build
ADD . $GOPATH/src/github.com/d-smith/docker-build
RUN go get -d -v
RUN go get github.com/lsegal/gucumber/cmd/gucumber
RUN GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o main .
RUN cp main /app
EXPOSE 8080
CMD /app/main
