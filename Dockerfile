FROM golang:1.22.5-alpine3.19 AS builder

WORKDIR /go/src/kubecenter.com/server

COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go mod tidy \
    && go build -o server .

FROM alpine:latest

LABEL MAINTAINER="benzhangchi@gmail.com"

WORKDIR /go/src/kubecenter.com/server

COPY --from=0 /go/src/kubecenter.com/server/.env ./.env
COPY --from=0 /go/src/kubecenter.com/server/.kube/config ./.kube/config
COPY --from=0 /go/src/kubecenter.com/server/server ./

EXPOSE 7080

ENTRYPOINT ["./server"]

