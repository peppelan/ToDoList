# More info here: https://hub.docker.com/_/alpine/
FROM alpine:latest

ADD todolist /usr/bin

EXPOSE 8080

ENTRYPOINT ["/usr/bin/todolist"]
