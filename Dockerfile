# 第一阶段：构建二进制文件
FROM golang:1.16 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o svc .

# 第二阶段：构建最终镜像
FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/svc .

CMD ["./svc"]