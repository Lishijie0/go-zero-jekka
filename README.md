# Jekka-api-go

## 启动方式

**1. 以docker compose的方式：** 
- 直接执行 `make up` 即可

**2. 直接启动：** 
- 执行 `go install github.com/cortesi/modd/cmd/modd@latest` 安装 `modd`
- 保证`CGO_ENABLED=1`，如果不是，执行 `go env -w CGO_ENABLED=1`
- 执行 `modd` 即可

## 服务列表

### 队列服务

| server name   | port | path     | desc                                   |
|---------------|------|----------|----------------------------------------|
| mqueue        | -    | -        | 基于redis的队列服务，使用asynq实现                 |
| asynq monitor | 8901 | /monitor | asynq 监控 web 端（用户名：jekka 密码：jekka2025） |

### API

| server name | port | desc      |
|-------------|------|-----------|
| third api   | 5001 | 三方服务相关api |

### RPC
| server name | port | desc      |
|-------------|------|-----------|
| greet rpc   | 4001 | 演示服务，可删除  |
| third rpc   | 4002 | 三方服务相关rpc |



## 目录结构和作用
``` 
├── app
│   ├── greet 演示服务，可删除，运行在4001端口 
│   └── third  三方服务（单个服务目录，一般是某微服务名称）
│       ├── cmd 服务目录  
│       │   ├── api 三方服务相关API，运行在5001端口
│       │   │   ├── desc 路由描述文件, https://go-zero.dev/docs/tasks/dsl/api
│       │   │   ├── etc 静态配置文件目录（env配置文件）
│       │   │   └── internal 内部服务（单个服务内部文件，其可见范围仅限当前服务）
│       │   │       ├─ config 配置文件（env配置的结构体文件）
│       │   │       ├─ handler （handler 目录，可选，一般 http 服务会有这一层做路由管理，handler 为固定后缀）
│       │   │       ├─ logic 逻辑处理
│       │   │       ├─ middleware 中间件
│       │   │       ├─ svc 依赖注入目录，所有 logic 层需要用到的依赖都要在这里进行显式注入
│       │   │       └─ types 结构体存放目录
│       │   └── rpc 三方服务相关RPC，运行在4002端口
│       ├── model 数据库目录 
│       └── third.go 程序启动入口文件 
├── pkg
│   ├── cache 缓存相关
│   ├── constant 常量相关
│   ├── db 数据库相关
│   ├── response 响应相关
│   └── services 公共服务相关 
├── data
├── logs
├── batchkill.sh 杀掉modd, 及其运行的所有子进程
├── docker-compose.yml 服务容器配置
├── go.mod go依赖管理文件
├── makefile
├── modd.conf modd热重启配置文件
└── README.md 说明文档
```

### 记录
```
2025.2.24 
1. 微服务中的 svc目录，先放在各个微服务中，等后期项目逐渐多之后再抽离到公共目录
2. 增加新队列方式: 
  a. pkg/constant/asynq.go 配置队列名称， 例如：third_high
  b. app/third/mq/internal/jobpayload.go 定义队列消费时传递的参数 
  c. app/third/mq/internal/jobtype.go 定义队列名称的标识
  d. app/third/cmd/mq/internal/svc/aynqserver.go Queues: map[string]int增加队列名
  e. app/third/cmd/mq/internal/handler/routes.go 消费队列注册路由
  f. 复制 *handler.go 文件开发逻辑
3. 数据库连接方式：
  a. 查询用户信息
    var user model.JkUser //引入表模型
	JkUser := l.svcCtx.DB.Model(&model.JkUser{}).Select("id", "username").Where("id = ?", payload.JkUserId).First(&user)
	if JkUser.Error != nil {
		return fmt.Errorf("查询用户失败: %v", JkUser.Error)
	}
	fmt.Println(user.Username)
  b. 数据库映射
	DB          main
	DBThird     third
	DBThirdV2   third_v2 (已经使用sharding切片，curd时需要带入jk_user_id)
```
```
2025.2.24 
1. 配置文件无法再目前架构中提取出公共目录，难度较大，目前采用将配置文件放在每个微服务中，然后通过modd.conf进行热重启，实现配置文件热更新。
2. 后期需要使用etcd配置各个微服务的配置
```
```
2025.2.25
1. makefile 中的model 命令，需要安装：go install gorm.io/gen/tools/gentool@latest
2. 同一个包地址下，func 和 struct 不能重名
```
```
2025.2.27
1. 记录api开发流程：
    a. api/desc/*.api文件，其中，路由的 @handler 不能重复
    b. 结构体定义在api/types/*.go中，或者在/types/*.go 文件中
    c. 定义好路由和结构体后，通过 make api_gen third 生成api代码
    d. third/api/handler文件下生成 *handler.go和third/api/logic文件下生成 *logic.go
    e. 最终在logic中开发逻辑即可。 
```