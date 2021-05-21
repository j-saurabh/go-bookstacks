FROM golang:1.16

LABEL maintainer="Saurabh Joshi"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN \
    go get -d -v \
    && go install -v \
    && go build -o bookstack .


EXPOSE 5000

CMD ["/app/bookstack"]

