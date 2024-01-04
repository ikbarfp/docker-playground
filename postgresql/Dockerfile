FROM postgres:16-alpine AS builder
WORKDIR /app/

FROM alpine:3
WORKDIR /app/
COPY --from=builder /app/main /app/
CMD /app/main


