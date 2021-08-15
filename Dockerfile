FROM golang AS builder
MAINTAINER long2ice "long2ice@gmail.com"
# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on
ENV GOOS=linux
ENV GOARCH=$GOARCH
ENV CGO_ENABLED=0
# 移动到工作目录：/build
WORKDIR /build
# 复制项目中的 go.mod 和 go.sum文件并下载依赖信息
COPY go.mod .
COPY go.sum .
RUN go mod download
# 将代码复制到容器中
COPY . .
# 将我们的代码编译成二进制可执行文件app
RUN go build -o app ./examples
###################
# 接下来创建一个小镜像
###################
FROM scratch
# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /build/app /
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Shanghai
# 启动容器时运行的命令
ENTRYPOINT ["/app"]