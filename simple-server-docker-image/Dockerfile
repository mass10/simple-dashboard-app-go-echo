FROM ubuntu:18.04

MAINTAINER mass10

RUN mkdir -p /my/app

RUN apt-get update

COPY . /my/app/

ENV PATH /my/app/go1.11.2.linux-amd64/bin

CMD /my/app/go1.11.2.linux-amd64/bin/go get github.com/labstack/echo

ENV APP_ROOT /my/app

WORKDIR $APP_ROOT

EXPOSE 8080 8080

CMD /my/app/go1.11.2.linux-amd64/bin/go run main.go
