# 简介

gin web 开发脚手架，拿来即用。

## 配置

- toml

## 组件

- worker

## 日志

- glog

## middleware

- identity: 身份认证，认证逻辑需自己实现
- req_resp_log: 请求响应日志输出，可定制输出需要的日志
- request_id: 绑定请求响应唯一id
- error_hanlder: 统一异常处理，本人没有用到，如需使用可自己实现异常处理逻辑

## 包

- common: 公共包，包含常量，错误，请求响应体格式定义...
- component: 组件包，包含worker，用户可自行添加项目中需使用的其他组件
- conf: 配置包
- controller: 接口实现包
- middleware: 中间件
- router: 路由包，统一在该包定义路由规则
- starter: 项目启动器，项目中需要启动的功能统一在该包中定义启动规则
- util: 工具包

## 说明

目前未加入orm，cache等相关功能，如有需要可自行添加。

