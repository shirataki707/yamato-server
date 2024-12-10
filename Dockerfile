FROM golang:1.23.4 as builder

WORKDIR /app
COPY . .

