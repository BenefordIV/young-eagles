FROM golang:1.26.1 AS builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o young_eagles ./

RUN chmod +x /app/young_eagles

#build tiny image
FROM alpine:latest

RUN mkdir /app

copy --from=builder /app/young_eagles /app

cmd [ "/app/young_eagles" ]

