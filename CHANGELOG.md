# CHANGELOG

## 1.0.0

1. 实现静态数据采集
2. 实现usage、process、connections、temperatures数据采集

## 1.0.1

1. 修正amd64打包文件中不包含配置文件的问题

## 1.0.2

1. 去除注册系统服务时的报错逻辑
2. 增加start和stop命令行参数
3. 修改默认配置文件内容

## 1.0.3

1. 修正ubuntu12系统下注册的系统服务无法启动的问题

## 1.0.4

1. 修正用户名被设置成用户描述的问题
2. 修正第一次获取CPU使用率时特别高的问题
3. 修改打包程序增加epoch

## 1.0.5

1. 修正windows下获取系统安装时间的问题

## 1.0.6

1. go版本升级到1.19.2
2. 升级第三方库版本

## 1.0.7

1. 升级第三方库版本
2. 优化打包脚本

## 1.0.8

1. go版本升级到1.19.3
2. 升级第三方库版本

## 1.0.9

1. 修正deb和rpm包中的epoch问题
2. 增加manifest.yaml配置信息描述文件

## 1.0.10

修改命令行交互方式

## 1.0.11

修正deb和rpm包安装时注册系统服务的问题

## 1.0.12

修改manifest.yaml增加配置项的最小值设置

## 1.0.13

1. 修正windows下的系统服务启动问题
2. 去除msi安装包

## 1.0.14

windows下的exe安装包支持静默卸载

## 1.0.15

1. 修改启动、停止、重启、注册、反注册系统服务失败时的状态码
2. 修正windows系统下未启动服务无法重启的问题

## 1.0.16

1. manifest.yaml中增加required字段
2. 修正golint问题

## 1.0.17

修改manifest.yaml中的basic.id字段类型为uuid

## 1.1.0

1. 支持arm架构
2. go版本升级到1.19.5

## 1.2.0

1. 新增日志文件上报功能
2. 升级第三方库版本
3. go版本升级到1.20

## 1.2.1

1. 新增nameserver相关字段
2. go版本升级到1.20.1

## 1.2.2

1. 新增CPU的load1、load5、load15负载信息字段 [#31](https://github.com/jkstack/metrics-agent/issues/31)
2. 新增磁盘的read_per_second、write_per_second、iops_in_progress速率信息字段 [#30](https://github.com/jkstack/metrics-agent/issues/30)
3. 调整打包脚本，新增oss上传逻辑

## 1.2.3

1. 修正websocket连接断开问题
2. 修正无法正常获取日志文件列表的问题
3. go版本升级到1.20.3

## 1.2.4

1. 修正监控频率设置过低时CPU耗尽的问题
2. go版本升级到1.20.7