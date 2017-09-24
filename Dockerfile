# More info here: https://hub.docker.com/_/alpine/
FROM alpine:latest

ADD bin/todolist /usr/bin

EXPOSE 8080
EXPOSE 8081

ENTRYPOINT ["/usr/bin/todolist"]
