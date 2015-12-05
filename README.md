This SOAP example was adapted from [xavi-sample](https://github.com/xtracdev/xavi-sample). The
provided Dockerfile will build the application in a Docker container, that can then be 
run via docker run:

<pre>
docker build .
docker run -d -p 8080:8080 209727a65d8b
</pre>

Note the project is setup to build in codeship's docker build infrastructure. Refer to the codeship-services.yml
and codeship-steps.yml files.

This example is a 4 stage CI pipeline:

1. The code is compiled and the unit tests run, and the binary is copied to a shared data volume
2. An acceptance test image is created for the system under test, linked with dependent services (a
mountebank reflector in this case),
3. The acceptance tests are compiled and executed against the test image
4. The binary is packaged into very small container and uploaded to docker hub.

A future refinement should seek to run the tests against the image that will be uploaded to docker hub to
verify both the application functionality and the packaging of the application.

One thing to note is the use of the GOOS and GOARCH variables in the compile command. When pulling
and running the final image on Mac OS X, the binary produced in stage one would not execute, giving
a '/bin/sh: main not found' error or something to that effect. Including the two options produced
an binary that ran correctly in the Mac OS X docker runtime.