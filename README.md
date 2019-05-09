# Simple Go webserver
This simple Go webserver prints all information about the networkinterfaces out. 
You can use the Dockerfile to build a docker image.

# Build on Mac OSX
GOOS=linux CGO_ENABLED=0 go build -a -x -o hello-webserver

