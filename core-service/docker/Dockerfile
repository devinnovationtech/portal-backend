# Builder
FROM harbor.digitalservice.id/proxyjds/library/golang:1.16-alpine as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app

COPY . .

RUN make engine
RUN make migrater

# Distribution
FROM harbor.digitalservice.id/proxyjds/library/alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 7070

COPY --from=builder /app/src/engine /app
COPY --from=builder /app/src/migrater /app
COPY --from=builder /app/src/database/migrations /app/migrations

CMD /app/engine
