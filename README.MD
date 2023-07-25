# 服务介绍

## 提供生产级别的，秒杀 saas 服务

## 服务功能规划

- pluto/orderService：秒杀功能的核心服务，负责库存相关操作
- ganymede/userService：用户注册、登录、认证服务
- europa/apiService: amusing-x 服务的 web api 接口
- risk： 用户的准入、限流, 风制
- product: 商品服务核心功能
- admin: admin 服务的 api 接口

## 依赖的软件包版本

```shell
protoc-gen-go v1.26.0
protoc        v3.12.3
```

## 数据库部署

## proto 编译

```shell
cd protos && make all
```
