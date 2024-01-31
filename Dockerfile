#################
# Base image
#################
FROM alpine:3.19.1 as echo-postgresql-base

USER root

RUN addgroup -g 10001 echo-postgresql && \
    adduser --disabled-password --system --gecos "" --home "/home/echo-postgresql" --shell "/sbin/nologin" --uid 10001 echo-postgresql && \
    mkdir -p "/home/echo-postgresql" && \
    chown echo-postgresql:0 /home/echo-postgresql && \
    chmod g=u /home/echo-postgresql && \
    chmod g=u /etc/passwd
RUN apk add --update --no-cache alpine-sdk curl

ENV USER=echo-postgresql
USER 10001
WORKDIR /home/echo-postgresql

#################
# Builder image
#################
FROM golang:1.21-alpine AS echo-postgresql-builder
RUN apk add --update --no-cache alpine-sdk
WORKDIR /app
COPY . .
RUN make build

#################
# Final image
#################
FROM echo-postgresql-base

COPY --from=echo-postgresql-builder /app/bin/echo-postgresql /usr/local/bin

# Command to run the executable
ENTRYPOINT ["echo-postgresql"]
