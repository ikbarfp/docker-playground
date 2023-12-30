FROM golang:1.21-alpine

ENV APP_PORT=8080
ENV APP_VERSION=v1.0.0

RUN mkdir app
COPY main.go app

EXPOSE ${APP_PORT}

CMD go run app/main.go