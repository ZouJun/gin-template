#打包镜像并输出可执行文件
FROM golang:1.16-alpine as build
ARG GO_ROOT=/opt/app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
	GOPROXY="https://goproxy.cn,direct"

WORKDIR $GO_ROOT

COPY . $GO_ROOT
RUN go build -o gin-template .

EXPOSE 50000
ENTRYPOINT ["./gin-template"]