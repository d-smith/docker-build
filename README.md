This SOAP example was adapted from [xavi-sample](https://github.com/xtracdev/xavi-sample). The
provided Dockerfile will build the application in a Docker container, that can then be 
run via docker run:

<pre>
docker build .
docker run -d -p 8080:8080 209727a65d8b
</pre>

Note the project is setup to build in codeship's docker build infrastructure. Refer to the codeship-services.yml
and codeship-steps.yml files.

Note: next step is to add mountebank to the image, configure and run it so the quote service returns something.


