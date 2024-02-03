FROM golang:1.21-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o bankService ./src

RUN chmod +x /app/bankService

FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/bankService /app

CMD ["/app/bankService"]
