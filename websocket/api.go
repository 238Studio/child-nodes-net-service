// Package childWebsocket
// 这个包是暴露给控制结构的API层
// 控制上位机的Websocket网络传输过程
package childWebsocket

import "github.com/238Studio/child-nodes-net-service/websocket/service"

// Init 初始化Websocket。作为暴露给外部的接口。
// 传入参数：websocket连接地址、ping的最大时间间隔、pong最大返回时间间隔
// 传出参数：websocket应用、错误信息
func Init(wsURL string, pingTimerMargin, pongWait int) (*service.WebsocketServiceApp, error) {
	return service.Init(wsURL, pingTimerMargin, pongWait)
}
