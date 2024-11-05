一. 目录说明
1. internal/models: 存放数据库模型（Model）定义。
2. internal/services: 存放业务逻辑服务（Service）,服务层通常包含应用程序的核心业务逻辑，并且可以被控制器（Handler）调用
3. internal/repositories: 存放数据访问层（Repository），仓库层通常负责与数据库进行交互，执行 CRUD 操作
4. internal/handlers: 存放控制器（Handler）,控制器负责处理 HTTP 请求，并调用相应的服务层逻辑
5. internal/requests: 存放请求参数的结构体定义，这些结构体通常用于接收和验证 HTTP 请求的输入数据
6. internal/validator: 存放验证器（Validator）的实现, 验证器负责验证请求参数的合法性
7. internal/middleware: 存放中间件（Middleware）的实现,中间件通常用于处理请求的预处理和后处理逻辑，例如身份验证、日志记录等