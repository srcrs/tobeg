FROM golang:alpine AS builder

# 设置环境变量
ENV GO111MODULE=on \
  CGO_ENABLED=1 \
  GOOS=linux

# 工作目录
WORKDIR /build

# 安装GCC
RUN apk add --no-cache \
  gcc \
  musl-dev

# 下载依赖包
COPY go.mod .
COPY go.sum .
RUN go mod download

# 复制代码到容器
COPY . .

# 编译代码
RUN go build -ldflags='-s -w -extldflags "-static"' -o tobeg .

# 构建小镜像
FROM scratch

COPY ./templates /templates
COPY ./public/ /public
COPY config.yml config.yml
COPY flow.db flow.db

COPY --from=builder /build/tobeg /

ENTRYPOINT ["/tobeg"]
