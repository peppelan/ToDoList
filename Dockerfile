# More info here: https://hub.docker.com/_/alpine/
FROM alpine:latest

ADD hello /usr/bin

ENTRYPOINT ["/usr/bin/hello"]