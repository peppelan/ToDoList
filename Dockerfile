# More info here: https://hub.docker.com/_/alpine/
FROM alpine:latest

ADD todolist /usr/bin

ENTRYPOINT ["/usr/bin/todolist"]