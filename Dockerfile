FROM golang:1.17-alpine as binary

ENV TZ=Europe/Prague

WORKDIR /go/src

ADD . /go/src

RUN CGO_ENABLE=0 go build -o main








FROM alpine:latest

RUN echo "http://147.32.232.215/alpine/v3.14/main" > /etc/apk/repositories
RUN echo "http://147.32.232.215/alpine/v3.14/community" >> /etc/apk/repositories

RUN apk update && apk upgrade
RUN apk add tzdata bluez

EXPOSE 80/tcp

RUN mkdir /app
WORKDIR /app

COPY --from=binary /go/src/main /app

CMD ["/app/main"]