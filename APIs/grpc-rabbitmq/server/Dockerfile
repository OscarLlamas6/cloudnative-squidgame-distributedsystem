# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
RUN touch .env
RUN chmod 777 .env

COPY . .

RUN go build -o /grpc-server

EXPOSE 6001

CMD [ "/grpc-server" ]