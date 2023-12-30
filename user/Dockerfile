FROM golang:1.21-alpine

ENV APP_PORT=8080
ENV APP_VERSION=v1.0.0

RUN mkdir /app

# 1. Create the user group
RUN addgroup -S devgroup

# 2. Create new user then add to specified user group
RUN adduser -S -D -h /app rambokonguser devgroup

# 3. Change ownership to desired directory
RUN chown -R rambokonguser:devgroup /app

# 4. Switch to specified user
USER rambokonguser

COPY main.go /app

EXPOSE ${APP_PORT}

CMD go run /app/main.go