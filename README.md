```
.
├── cmd/                     # 应用程序入口
│   └── main.go              # 主程序入口
├── internal/                # 内部业务逻辑
│   ├── user/                # 用户模块
│   │   ├── handler.go       # HTTP 请求处理
│   │   ├── service.go       # 业务逻辑
│   │   └── repository.go    # 数据访问层
│   ├── post/                # 文章模块
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── comment/             # 评论模块
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   ├── tag/                 # 标签模块
│   │   ├── handler.go
│   │   ├── service.go
│   │   └── repository.go
│   └── auth/                # 认证模块
│       ├── middleware.go    # 身份验证中间件
│       └── jwt.go           # JWT 处理逻辑
├── pkg/                     # 可复用组件
│   ├── config/              # 配置管理
│   │   └── config.go        # 配置加载
│   ├── logger/              # 日志记录
│   │   └── logger.go        # 日志工具
│   ├── storage/             # 文件存储
│   │   └── storage.go       # 文件上传下载
│   └── validator/           # 输入验证
│       └── validator.go     # 参数校验工具
├── models/                  # 数据模型
│   ├── user.go              # 用户模型
│   ├── post.go              # 文章模型
│   ├── comment.go           # 评论模型
│   └── tag.go               # 标签模型
├── routes/                  # 路由定义
│   └── routes.go            # HTTP 路由注册
└── utils/                   # 工具函数
    └── errors.go            # 统一错误处理
```