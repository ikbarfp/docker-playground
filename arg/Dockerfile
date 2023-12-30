FROM golang:1.21-alpine

ENV APP_PORT=8080
ENV APP_VERSION=v1.0.0

ARG app=main

RUN mkdir app
COPY main.go app
RUN mv app/main.go app/${app}.go

EXPOSE ${APP_PORT}

# Little tricks
#   so when it comes to run time, we need to do some 'tricks' to store to a ENV (MY_APP) value
#   from ARG value (app), because ARG instruction only running on build time not on run time
ENV MY_APP=${app}
CMD go run app/${MY_APP}.go