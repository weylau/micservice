# 微服务实践demo

### 架构
- 开发语言：golang
- rpc框架：thrift
- 其他：MySQL

### demo功能
主要通过两个服务来实现了用户登录功能，这里没有实现gatewayapi层

- user-edge-service
> 主要用于实现登录逻辑，包括校验密码，校验谷歌码，生成token

- user-service
> 主要负责数据库查询，也可叫db下沉服务
