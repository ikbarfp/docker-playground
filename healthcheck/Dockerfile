FROM golang:1.21-alpine

ENV APP_PORT=8080
ENV APP_VERSION=v1.0.0

RUN apk --no-cache add curl
RUN mkdir app

COPY main.go app

EXPOSE ${APP_PORT}

HEALTHCHECK --interval=5s --start-period=5s CMD curl -f http://localhost:${APP_PORT}/health
CMD go run app/main.go