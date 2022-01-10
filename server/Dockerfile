FROM golang:alpine as builder

WORKDIR /root/zga/zga-client-restful/server
COPY . .

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w CGO_ENABLED=0
RUN go env
RUN go mod tidy
RUN go build -o manage .

FROM alpine:latest
LABEL MAINTAINER="sun.ji@zenlayer.com"

WORKDIR /root/zga/zga-client-restful/server

COPY --from=0 /root/zga/zga-client-restful/server ./

EXPOSE 8888

ENTRYPOINT ./manage
