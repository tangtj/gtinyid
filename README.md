# gtinyid
## 简介
**gtinyid** 用golang开发的分布式id生成器，基于号段模式算法实现，模仿 [tinyid](https://github.com/didi/tinyid) 项目实现的golang版本。
生成全局唯一，趋势递增的64位整形id。

[Tinyid原理介绍](https://github.com/didi/tinyid/wiki/tinyid%E5%8E%9F%E7%90%86%E4%BB%8B%E7%BB%8D)

## 使用
### 拉取代码
```shell
git clone https://github.com/tangtj/gtinyid
```
### 建表
1. 使用 dataspirce.sql 创建表结构
2. 调整config.yaml配置

### 运行
```shell
go build ./main/main.go -o main
go run main
```
### 使用
- http api

    参考`gtinyid.postman_collection.json`postman导出文件

## TODO
- [X] 便于使用的http封装的sdk
- grpc
