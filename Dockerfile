# Build staged
FROM golang:1.23-alpine3.20 AS builder
WORKDIR /http-server
COPY ./vendor ./vendor
COPY . .
RUN go build -mod=vendor -o /http-server/build/server /http-server/cmd/server.go


FROM alpine:3.20
WORKDIR /http-server
COPY --from=builder /http-server/build/server ./core-server
COPY docker.env .env
COPY templates templates

EXPOSE 8080
CMD ["/http-server/core-server"]
