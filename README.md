# go-spring-gin-examples
go spring 整合 gin 框架使用案例

# 框架
```bash
# go-spring 框架
https://github.com/go-spring/spring-core
# gin 框架
https://github.com/gin-gonic/gin
```

# 导入数据库（Postgresql） sql
使用 Postgresql 数据库，其他数据库 请自行修改。
创建 数据库 demo，导入下面sql 文件
```bash
doc/sql/demo.sql
```
>注意: 如过不想用数据库测试案例，那么把配置文件`app.toml`中，`[database]`下,`enabled = true`改为`enabled = false`


# 运行
## 默认 单项目 启动
```bash
go run main.go
```
## 多项目运行 请自行更改端口
```bash
go run cmd/admin/main.go

go run cmd/api/main.go
```
## 多服务端口，统一一个项目启动
先 关闭本 案例项目中 所有已经启动的 项目
```bash
go run cmd/moreServer/main.go
```

#  项目结构
```bash
.
├── app
│   └── domainDemo                              模块
│       ├── controller                          控制器
│       ├── model                               模型实体
│       │   └── modUser                         用户模型实体
│       └── service                             服务
├── assets                                      web 静态资源
│   └── img                                     图片
├── cmd                                         命令文件目录，启动入口文件目录(多项目分别启动时使用)
│   ├── admin                                   admin项目
│   ├── api                                     api项目
│   └── moreServer                              多服务端口 启动
├── config                                      配置文件目录
├── doc                                         文档说明
├── infrastructure                              实体类，结构体，数据库表映射，资源映射
│   ├── entityDemo                              数据库表映射
│   └── repositoryDemo                          资源映射
├── middleware                                  中间件
│   ├── auth                                    权限认证
│   ├── db                                      数据库
│   │   └── dbPostgresql
│   ├── runner                                  应用启动后立即执行的一次性任务（初始化等）
│   ├── server                                  服务端，多服务端
│   │   ├── ginServer                           gin 服务端
│   │   └── httpServer                          自定义 http 服务端
│   └── validator                               验证器
├── pkg                                         可以被其他应用导入
│   ├── config                                  配置结构映射
│   └── tools                                   工具类
│       └── wrapperPg
│           └── rg
├── router                                      路由
└── scripts                                     构建、安装、分析等不同功能的脚本文件
└── test                                        测试目录


```



# 测试
## 首页
用浏览器 访问
```bash
http://127.0.0.1:9980/
```
## 查看端口
用浏览器 访问
```bash
http://127.0.0.1:9980/port
```
## 查看数据库配置信息
用浏览器 访问
```bash
http://127.0.0.1:9980/db-config
```

## json post 验证器 返回 中文 [错误情况]
```bash
curl -X POST --header "Content-Type: application/json" http://127.0.0.1:9980/admin/user/create?pass=ok \
-d '{"xxxx":"zhangsan"}'
```
结果 
```bash
{"code":"500","message":"操作失败","messageData":{"name":"名称为必填字段"},"data":""}
```

## json post 验证器 返回 中文 [正确情况]
```bash
curl -X POST --header "Content-Type: application/json" http://127.0.0.1:9980/admin/user/create?pass=ok \
-d '{"name":"zhangsan"}'
```
结果
```bash
{"code":"200","message":"操作成功","data":""}
```

## 权限认证 [失败情况]
```bash
curl -X POST --header "Content-Type: application/json" http://127.0.0.1:9980/admin/user/create \
-d '{"name":"zhangsan"}'
```
结果
```bash
{"code":"500","message":"权限验证失败","data":""}
```
## 权限认证 [正确情况]
```bash
curl -X POST --header "Content-Type: application/json" http://127.0.0.1:9980/admin/user/create?pass=ok \
-d '{"name":"zhangsan"}'
```
结果
```bash
{"code":"200","message":"操作成功","data":""}
```

## 图片
用浏览器 访问 
```bash
http://127.0.0.1:9980/assets/img/girl.jpg
```

## 数据库保存
```bash
curl -X POST --header "Content-Type: application/json" http://127.0.0.1:9980/admin/user-db/create?pass=ok \
-d '{"name":"zhangsan"}'
```
结果
```bash
{"code":"200","message":"操作成功","data":""}
```

## 查看数据库参数
用浏览器 访问
```bash
http://127.0.0.1:9980/db-config
```