FROM ubuntu:18.04

MAINTAINER mass10

RUN mkdir -p /my/app


RUN apt-get update

# WGET https://dl.google.com/go/go1.11.2.linux-amd64.tar.gz .

COPY . /my/app/
# COPY go1.11.2.linux-amd64 /my/app/

ENV PATH /my/app/go1.11.2.linux-amd64/bin

CMD go get github.com/labstack/echo

ENV APP_ROOT /my/app

WORKDIR $APP_ROOT

EXPOSE 8080 8080

# CMD ["/my/app/go1.11.2.linux-amd64/bin/go", "run", "main.go"]
CMD /my/app/go1.11.2.linux-amd64/bin/go run main.go
