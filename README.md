## 项目结构说明

本项目采用 Gin 框架构建，遵循清晰的目录结构，以便于维护和扩展。以下是各个目录和文件的详细说明：

## 目录结构

.
├── cmd
│ ├── main.go # 应用程序的主入口文件，负责启动 Gin 服务器
│ └── wire.go # 用于依赖注入的配置文件，通常使用 Google 的 Wire 工具生成依赖注入代码
├── config
│ ├── dev.yaml # 开发环境的配置文件
│ └── prod.yaml # 生产环境的配置文件
├── internal
│ ├── config
│ │ ├── config.go # 配置文件的解析和加载逻辑
│ │ ├── database.go # 数据库连接和配置
│ │ ├── log.go # 日志配置
│ │ └── redis.go # Redis 连接和配置
│ ├── handlers # HTTP 请求的处理逻辑，通常对应视图层
│ ├── models # 数据模型定义，通常对应数据库表结构
│ ├── repositories # 数据库操作的封装，通常对应数据访问层
│ ├── requests # 请求参数的验证和处理
│ ├── router # 路由定义，负责将请求映射到相应的处理函数
│ └── services # 业务逻辑的封装，通常对应服务层
├── pkg
│ ├── db # 数据库相关的工具和辅助函数
│ ├── log # 日志相关的工具和辅助函数
│ ├── middlewares # Gin 中间件的定义和实现
│ ├── response # HTTP 响应的封装和处理
│ ├── server # 服务器相关的配置和启动逻辑
│ ├── utils # 通用的工具函数
│ └── validation # 请求参数的验证逻辑
└── storage
└── logs # 日志文件的存储目录

## 日志

日志文件存储在 storage/logs 目录下，日志配置在 internal/config/log.go 中定义。

## 配置文件

配置文件位于 config 目录下，支持开发环境和生产环境的不同配置。配置文件使用 YAML 格式。

## 依赖注入

wire cmd/wire.go

## 启动main

go build gin-api/cmd