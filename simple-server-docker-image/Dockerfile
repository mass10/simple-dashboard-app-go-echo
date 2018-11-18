FROM ubuntu:18.04

MAINTAINER mass10

# INSTALLATION
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update
RUN apt-get upgrade -y
RUN apt-get install apt-utils -y
RUN apt-get install git -y

# APPLICATION FILES
RUN mkdir -p /my/app
COPY . /my/app/

ENV PATH /my/app/go1.11.2.linux-amd64/bin:$PATH

RUN /my/app/go1.11.2.linux-amd64/bin/go get github.com/labstack/echo

ENV APP_ROOT /my/app

WORKDIR $APP_ROOT

EXPOSE 8080 8080

CMD /my/app/go1.11.2.linux-amd64/bin/go run main.go

