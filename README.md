# 快速开始

本项目入口点在cmd/main.go若要手动运行go build后移到根目录运行

## 使用docker(未安装mysql)

1. ```
   git clone https://github.com/rn-consider/compuBackend.git
   ```

2. ```
   cd compuBackend
   ```

3. ```
   docker-compose up -d
   ```
# 健康检查
  ```
  curl -f http://127.0.0.1:8010/health
  ```
# 手动部署

1. 修改config/config.yaml中的文件

2. 进入cmd目录手动编译并运行auto_migrate/main.go,main.go