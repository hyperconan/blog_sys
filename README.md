# Blog Microservices System

基于 go-zero 框架的微服务博客系统，包含用户服务、博客服务、评论服务和API网关。

## 架构

```
blog_sys/
├── gateway/          # HTTP网关服务 (端口8888)
├── user/            # 用户服务 (GRPC端口8080)
├── blog/            # 博客服务 (GRPC端口8081)
└── comment/         # 评论服务 (GRPC端口8082)
```

## 功能特性

### 用户服务 (User Service)
- 用户注册
- 用户登录
- JWT Token验证

### 博客服务 (Blog Service)
- 创建文章
- 更新文章
- 删除文章
- 获取所有文章
- 获取文章详情

### 评论服务 (Comment Service)
- 创建评论
- 根据文章ID获取评论

### 网关服务 (Gateway)
- RESTful API接口
- 统一认证
- 服务路由

## 技术栈

- **框架**: go-zero
- **通信**: gRPC + HTTP
- **服务发现**: ETCD
- **数据库**: MySQL + GORM

## 快速开始

### 前置条件

1. 安装 Go 1.25+
2. 安装 MySQL 8.0+
3. 安装 ETCD
4. 安装 goctl

### 启动ETCD

```bash
etcd
```

### 启动服务

按以下顺序启动服务：

1. 启动用户服务
```bash
cd user && ./user-server -f etc/user.yaml
```

2. 启动博客服务
```bash
cd blog && ./blog-server -f etc/blog.yaml
```

3. 启动评论服务
```bash
cd comment && ./comment-server -f etc/comment.yaml
```

4. 启动网关服务
```bash
cd gateway && go run gateway.go -f etc/gateway-api.yaml
```

## API接口

### 用户接口

- `POST /user/register` - 用户注册
- `POST /user/login` - 用户登录

### 博客接口

- `POST /blog` - 创建文章
- `GET /blog/all` - 获取所有文章
- `PUT /blog/:post_id` - 更新文章
- `DELETE /blog/:post_id` - 删除文章

### 评论接口

- `POST /blog/comment` - 创建评论 (需要JWT认证)
- `GET /blog/comment/:post_id` - 获取文章评论

## JWT 认证

以下接口需要 JWT 认证，在请求头中添加 `Authorization: Bearer <token>`：

- `POST /blog` - 创建文章
- `PUT /blog/:post_id` - 更新文章
- `DELETE /blog/:post_id` - 删除文章
- `POST /blog/comment` - 创建评论

**认证流程：**
1. 先调用 `POST /user/login` 获取 JWT token
2. 在后续请求的 `Authorization` 头中添加 `Bearer <token>`

## 注意事项

1. 现在使用 MySQL 数据库存储数据，重启服务后数据会持久化保存
2. 实现了完整的 JWT 认证机制
3. 需要先启动 MySQL 和 ETCD 服务才能正常运行
4. 服务端口已在配置文件中预设
5. 数据库连接配置在各服务的 etc/*.yaml 文件中

## 开发说明

每个服务都是独立的，可以单独开发和部署：

- `cd {service_name} && go build -o {service_name}-server .`
- 配置文件位于 `etc/` 目录
- 业务逻辑位于 `internal/logic/` 目录
- gRPC定义位于 `*.proto` 文件