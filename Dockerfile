# 使用一个基础镜像
FROM golang:1.21.3

# 设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY="https://goproxy.cn,direct" \
    TZ=Asia/Shanghai

# 在容器内创建一个工作目录
WORKDIR /app

# 复制当前目录下的所有代码到容器中
COPY . .

# 进入工作目录并运行go mod download
WORKDIR /app
RUN go mod download

# 编译并运行 cmd/auto_migrate/main.go
WORKDIR /app/cmd/auto_migrate
RUN go build -o auto_migrate main.go
RUN mv /app/cmd/auto_migrate/auto_migrate /app/auto_migrate


# 编译并运行 cmd/main.go，并暴露8010端口
WORKDIR /app/cmd
RUN go build -o main main.go
# 移动 main 文件到上一层的根目录
RUN mv /app/cmd/main /app/main
WORKDIR /app
# 执行数据表初始化
RUN ./auto_migrate
CMD ["./main"]
EXPOSE 8010
