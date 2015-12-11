# CodeShip Docker Build Sample

## Overview

This SOAP example was adapted from [xavi-sample](https://github.com/xtracdev/xavi-sample). 

This project shows how to use Docker containers to build and run the tests for an application, then package
it for deployment. The project is setup to build in Codeship's docker build infrastructure. Refer to the codeship-services.yml
and codeship-steps.yml files.

This example is a 4 stage CI pipeline:

1. The code is compiled and the unit tests run, and the binary is copied to a shared data volume
2. An acceptance test image is created for the system under test, linked with dependent services (a
mountebank reflector in this case),
3. The acceptance tests are compiled and executed against the test image
4. The binary is packaged into very small container and uploaded to docker hub.

## Local Build

To make the build process efficient for scenarios where dependencies and environment change less frequently
than code or tests, we layer in dependencies and environment as early as possible to allow caching. As a prequisite
to running `jet steps`, manually build the gomb and dockerbuildprojectdeps images first, e.g.

<pre>
docker build -t dasmith/gomb -f Dockerfile.gomb
docker push dasmith/gomb
docker build -t dasmith/dockerbuildprojectdeps -f Dockerfile.projectdeps .
docker push dasmith/dockerbuildprojectdeps
</pre>

With the prerequisite images in place, we run the pipeline via `jet steps`. Note that the final push to Dockerhub is 
not done in the local build; a push to github triggers the codeship build, which will do the final push if
all the stanges succeed.

## Running the Image

Once the Codeship build completes with a push to Dockerhub, the sample can be run thusly. First, start the containers:

<pre>
docker pull dasmith/quote
docker pull dasmith/mb-server-alpine
docker run -d -p 2525:2525 --name mountebank --label 'xt-container-type=atest-mb' dasmith/mb-server-alpine
docker run -d -p 8080:8080 --name main --link mountebank:mb dasmith/quote:latest
</pre>

Note on the mac you need to have port forwarding at the virtual box level configured to allow localhost access on ports
2525 and 8080.

Next, set up the imposter endpoint:

<pre>
curl -i -X POST -H 'Content-Type: application/json' -d@imposter.json localhost:2525/imposters
</pre>

Finally, curl for a stock quote (which is fixed, don't panic or call your broker when you see the response)

<pre>
curl localhost:8080/quote/xxx
</pre>



## Futures and Notes

A future refinement should seek to run the tests against the image that will be uploaded to docker hub to
verify both the application functionality and the packaging of the application.

One thing to note is the use of the GOOS and GOARCH variables in the compile command. When pulling
and running the final image on Mac OS X, the binary produced in stage one would not execute, giving
a '/bin/sh: main not found' error or something to that effect. Including the two options produced
an binary that ran correctly in the Mac OS X docker runtime.