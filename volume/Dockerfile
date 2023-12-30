FROM golang:1.21-alpine

ENV APP_PORT=8080
ENV APP_VERSION=v1.0.0
ENV TEMP_DIR=/logs

RUN mkdir ${TEMP_DIR}

RUN mkdir app
COPY main.go app

EXPOSE ${APP_PORT}

VOLUME ${TEMP_DIR}

CMD go run app/main.go