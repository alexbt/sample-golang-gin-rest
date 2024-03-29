FROM golang:1.20.4-alpine AS builder
ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPATH=/go
ENV PATH=$PATH:$GOPATH/bin

RUN apk add git
RUN mkdir /app
WORKDIR /app
ADD ./src ./

RUN go mod tidy
RUN go build ./...
RUN go build ./cmd/...
RUN chmod 755 ./mytest

FROM alpine:latest
COPY --from=builder /app/mytest /tmp/mytest

EXPOSE 8080
ENTRYPOINT /tmp/mytest
