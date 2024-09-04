### 项目介绍
...

### 开始

```shell
go run cmd/server.go

# or 

air
```

### 构建

```shell
go build cmd/server.go
```

### 目录结构

```text

.
├── app                         //主 app
│   ├── app.go                  //注册 app
│   ├── handler                 //handler 处理方法
│   ├── middleware              //中间件
│   ├── model                   //模型
│   ├── repo                    //存储库
│   └── service                 //服务
├── cmd                 
│   └── server.go               //启动服务
├── config                      //配置
│   └── config.go
├── config.yaml                 //配置文件
├── constants                   //全局变量 数据库链接和配置
│   └── constants.go
├── core                        
│   ├── app.go                  // app 结构定义
│   ├── curd                    // curd 通用方法
│   └── response                // response 通用
├── dist
│   └── server
├── docker
│   ├── data
│   └── docker-compose.yaml
├── docs                        //swag文档
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── driver                      //数据库链接
│   └── db.go
├── go.mod
├── go.sum
├── libs                        //三方依赖
│   └── dtk
├── modules                     //多 app 模式，类似 django
│   └── post
├── storage                     //上传文件存储目录
├── tmp
└── utils                       //通用的工具
    └── recover.go

```