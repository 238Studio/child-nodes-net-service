# universal robot drive child node net service

通用机器人子节点网络服务。包含HTTP和Websocket两大模块。

## 使用
`go get github.com/238Studio/child-nodes-database-service`

## HTTP

封装了`GET`、`POST`、`PUT`、`DELETE`的请求实现。

## Websocket

对Websocket进行了封装。包含消息读取、消息发送、心跳（ping定时发送和pong超时检测）。