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

# 架构参考
```
> ├── cmd # 项目入口程序存放目录
> │ └── main.go # 项目入口点
> ├── config # 配置文件目录
> ├── controller # 控制器目录，用于路径控制
> │ └── middleware # 中间件目录
> ├── dao # 数据访问层
> │ └── mysql.go # MySQL 数据访问层
> | ── initialize # 项目初始化目录
> ├── docker-compose.yml # Docker Compose 配置文件
> ├── Dockerfile # Docker 构建文件，Dockerfile
> ├── docs # 文档目录
> ├── go.mod # Go 项目模块文件
> ├── go.sum # Go 项目依赖总结文件
> ├── models # 数据库映射层，用于各表的增删改查操作
> ├── README.md # 项目说明文档
> ├── service # 复杂逻辑层，处理业务逻辑
> ├── test # 测试层，用于编写项目测试
> └── utils # 工具目录，封装后的实用工具
>  └── jwt.go # JWT (JSON Web Token) 相关工具
```