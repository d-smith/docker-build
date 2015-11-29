FROM golang:latest
RUN mkdir /app

RUN curl -sL https://deb.nodesource.com/setup_4.x | bash -
RUN apt-get install -y nodejs
RUN npm install -g mountebank

RUN echo $GOPATH
RUN echo ls $GOPATH/src
RUN mkdir -p  $GOPATH/src/github.com/d-smith/docker-build
WORKDIR $GOPATH/src/github.com/d-smith/docker-build
ADD . $GOPATH/src/github.com/d-smith/docker-build
RUN go get -d -v
RUN go get github.com/lsegal/gucumber/cmd/gucumber
RUN go build -o main .
RUN cp main /app
RUN cp run.sh /app
EXPOSE 8080
EXPOSE 2525
CMD /app/run.sh
