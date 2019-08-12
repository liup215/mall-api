# mall-api

## 项目概述

人人商城的golang实现。人人商城是基于微擎的微信分销商城。本项目基于人人商城的数据库结构和业务逻辑，用golang重写的后台系统。本项目仅用于学习，不能用于线上应用，不可线上使用。

本项目使用微服务架构，系统间采用grpc调用。项目部署推荐docker+traefik方式部署。

本项目暂不涉及前端页面部分，前端的开发也在计划中。

## 目录

```
├── app	业务代码，所有的业务代码都在这个目录下
│   ├── api	对外暴露业务接口，restful api
│   │   ├── admin	系统管理后台接口，面向系统管理员
│   │   ├── mobile	移动端接口，面向C端用户
│   │   └── web	商家端管理后台接口，面向使用系统的商家
│   │       ├── account	账户模块
│   │       ├── finance	财务模块
│   │       ├── goods	商品模块
│   │       ├── groups	拼团模块
│   │       ├── member	会员模块
│   │       ├── order	订单模块
│   │       ├── pay	支付模块
│   │       ├── perm	
│   │       ├── sale	促销模块
│   │       ├── shop	商城模块
│   │       ├── system	系统管理模块
│   │       ├── user	商家用户模块
│   │       └── wechat	微信相关接口
│   ├── job	任务系统，后台运行的定时任务、队列任务等
│   │   └── main
│   │       └── groups 拼团模块
│   └── service	gRPC调用接口，供业务接口或任务系统调用，所有负责数据处理的操作都在这里进行
│       └── main
│           ├── account
│           ├── adv
│           ├── finance
│           ├── goods
│           ├── groups
│           ├── member
│           ├── order
│           ├── payment
│           ├── perm
│           ├── sale
│           ├── shop
│           ├── sysset
│           ├── system
│           ├── user
│           └── wechat
├── lib	项目依赖，封装的包都在这里
│   ├── address	地址管理
│   ├── database	数据库操作
│   │   ├── nosql	nosql
│   │   │   └── mongo	mongo操作封装
│   │   ├── orm	关系型数据库（mysql）
│   │   └── redis	redis
│   ├── mq	消息队列
│   │   └── nsq
│   ├── net
│   │   ├── http
│   │   └── rpc
│   ├── platform	第三方平台
│   │   └── wechat
│   ├── strings	字符串操作
│   ├── time	时间戳封装
│   └── utils	其他工具
├── main.go
└── model	人人商城数据模型，根据人人商城数据库生成的golang struct。
```

说明：

- app目录下为所有的业务代码，api下为对外暴露的restful接口；job下为系统运行的后台任务；service为grpc服务端，共api和job下的应用调用；service下的应用也是唯一可以操作数据库的应用（api下暂时还未使用grpc调用，之后的升级将重点完成）；
- 业务代码中，每个系统服务都是单独的文件夹，cmd文件夹为服务入口；

## 依赖

1. gin ( gin-jwt, gin-cors )
2. grpc 
3. nsq
4. redis ( redigo )
5. mongodb ( mgo )
6. mysql ( gorm )

## 功能

（未完成）

## 启动部署

由于此项目基于人人商城，使用前请自行获取人人商城系统；本项目不会透露任何人人商城相关源代码。

1. 配置最小开发环境
   - Mysql
   - Golang (v1.1以上)
   - Redis
   - mongo(部分功能使用)
   - Nsq ( 系统内支付回调的消息推送 )
2. 依次启动service服务
3. 依次启动job服务
4. 按需启动api服务

## 警告

1. 本项目仅用于学习使用。
2. 本项目不完善，仍处于并将长期处于开发中，不承担任何使用后果。

## 联系方式

微信： liup215

QQ： 2256935429

## License

不懂，随便用吧！！！