package service

import (
	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	"github.com/gorilla/websocket"
)

// Init 初始化websocket服务
// 传入参数：websocket连接地址、ping的最大时间间隔、pong最大返回时间间隔
// 传出参数：websocket应用、错误信息
func Init(wsURL string, pingTimerMargin, pongWait int) (*WebsocketServiceApp, error) {
	// 初始化websocket服务
	app := &WebsocketServiceApp{
		wsURL:           wsURL,
		pingTimerMargin: pingTimerMargin,
		pongWait:        pongWait,
	}

	// 初始化模型消息通道结构体
	app.modelMessageChannels = make(map[string]*ModelMessageChan)

	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

	if err != nil {
		return nil, util.NewError(_const.TrivialException, _const.Network, err)
	}

	app.conn = conn

	return app, nil
}
