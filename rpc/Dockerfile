FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/zero

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
COPY rpc/etc /app/etc
RUN go build -ldflags="-s -w" -o /app/mag rpc/mag.go


FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/mag /app/mag
COPY --from=builder /app/etc /app/etc

CMD ["./mag", "-f", "etc/mag.yaml"]
