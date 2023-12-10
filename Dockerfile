# 第一阶段：构建阶段
FROM golang:1.21 as builder

WORKDIR /app

# 先复制go.mod和go.sum并下载依赖，利用Docker缓存层
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码并构建可执行文件
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o main .

# 第二阶段：运行阶段
FROM alpine:latest

# 为了安全性，添加非root用户
RUN adduser -S -D -H -h /app appuser
USER appuser

WORKDIR /app

# 从构建阶段复制构建好的二进制文件
COPY --from=builder /app/main .

# 从构建阶段复制配置文件，保留配置文件的目录结构
COPY --from=builder /app/config ./config

# 暴露端口
EXPOSE 8080

# 运行应用
CMD ["./main"]