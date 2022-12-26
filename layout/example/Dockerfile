FROM golang:1.19.4-alpine3.17 AS builder

WORKDIR /app
ADD . /app
ENV GOPROXY httos://goproxy.cn,direct
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main main.go

FROM alpine AS runner

WORKDIR /app

# 复制编译阶段编译出来的运行文件到目标目录
COPY --from=builder /app/main .
# 将时区设置为东八区
RUN echo "https://mirrors.aliyun.com/alpine/v3.17/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.17/community/" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone

EXPOSE 7788
EXPOSE 7789

# 可外挂的目录
#VOLUME /opt/singsen/conf /app/conf

ARG configfile
ENTRYPOINT ./main -c ${configfile}

