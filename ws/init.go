package ws

import (
	"context"

	"github.com/238Studio/child-nodes-error-manager/errpack"
	"github.com/gorilla/websocket"
)

// InitWebsocketService 初始化websocket服务
// 传入参数：websocket连接地址、ping的最大时间间隔、pong最大返回时间间隔
// 传出参数：websocket应用、错误信息
func InitWebsocketService(wsURL string, pingTimerMargin, pongWait int) (*WebsocketServiceApp, error) {
	// 初始化websocket服务
	app := &WebsocketServiceApp{
		wsURL:           wsURL,
		pingTimerMargin: pingTimerMargin,
		pongWait:        pongWait,
		ErrorThrower:    make(chan error),    //错误抛出器
		stop:            make(chan struct{}), //停止通道
	}

	// 初始化模型消息通道结构体
	app.modelMessageChannels = make(map[string]*ModelMessageChan)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		return nil, errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	app.conn = conn

	// 初始化上下文
	app.ctx = new(ctx)
	app.ctx.ctx, app.ctx.stop = context.WithCancel(context.Background()) //goroutine全局退出机制

	//开始ping
	go app.ping()

	return app, nil
}
