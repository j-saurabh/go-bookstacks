FROM golang:1.16

LABEL maintainer="Saurabh Joshi <saurabh2226@gmail.com>"

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN \
    go get -d -v \
    && go install -v \
    && go build -o bookstack .


EXPOSE 5000

CMD ["/app/bookstack"]

