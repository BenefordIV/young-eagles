# base go image
FROM golang:1.23.4-alpine3.20 AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o young-eagles-app .

RUN chmod +x /app/young-eagles-app

#build tiny image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/young-eagles-app /app

CMD [ "/app/young-eagles-app" ]