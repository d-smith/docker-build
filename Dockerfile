FROM dasmith/dockerbuildprojectdeps


RUN mkdir /app

WORKDIR $GOPATH/src/github.com/d-smith/docker-build
ADD . $GOPATH/src/github.com/d-smith/docker-build
RUN GOOS=linux GOARCH=386 CGO_ENABLED=0 go build -o main .
RUN cp main /app
EXPOSE 8080
CMD /app/main
