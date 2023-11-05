package service

import (
	"errors"

	_const "github.com/UniversalRobotDriveTeam/child-nodes-assist/const"
	"github.com/UniversalRobotDriveTeam/child-nodes-assist/util"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

// WebsocketServiceApp Websocket 读消息
// 传入:无
// 传出:无
// 利用管道进行消息传递
func (app *WebsocketServiceApp) read() {
	//获取websocket连接
	conn := app.conn

	//先捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, err.(error))
		}
	}()

	//循环读取消息
	for {
		select {
		case <-app.stop:
			return
		default:
			_, message, err := conn.ReadMessage()
			if err != nil {
				app.ErrorThrower <- util.NewError(_const.CommonException, _const.Network, err)
				continue //出错后继续读取
			}
			//消息进一步处理
			go app.handelMessage(message)
		}
	}
}

// handelMessage 消息处理
// 传入参数：消息
// 传出参数：无
func (app *WebsocketServiceApp) handelMessage(message []byte) {
	//捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, err.(error))
		}
	}()

	//json格式化
	json := jsoniter.ConfigCompatibleWithStandardLibrary
	var websocketMessage WebsocketMessage
	err := json.Unmarshal(message, &websocketMessage)
	if err != nil {
		app.ErrorThrower <- util.NewError(_const.CommonException, _const.Network, err)
		return
	}

	//获取消息通道，进行消息分发
	modelMessageChan, ok := app.modelMessageChannels[websocketMessage.ModuleName]
	if !ok {
		app.ErrorThrower <- util.NewError(_const.TrivialException, _const.Network, errors.New("对应消息通道不存在"))
		return
	}

	//将消息写入消息通道
	modelMessageChan.ReadMessage <- websocketMessage
}

// WebsocketServiceApp Websocket 写消息
// 传入:无
// 传出:无
// 利用管道进行消息传递
func (app *ModelMessageChan) write() {
	//获取websocket连接
	conn := app.conn

	//先捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorChan <- util.NewError(_const.FatalException, _const.Network, err.(error))
		}
	}()

	for {
		select {
		case message := <-app.writeMessage:
			if message.IsBytes {
				err := writeBinary(conn, message.Data)
				if err != nil {
					app.ErrorChan <- util.NewError(_const.CommonException, _const.Network, err)
				}
			} else {
				err := writeJSON(conn, message.Data)
				if err != nil {
					app.ErrorChan <- util.NewError(_const.CommonException, _const.Network, err)
				}
			}

		//停止通道
		case <-app.stop:
			return
		}
	}
}

// writeJSON 写JSON
// 传入参数：websocket连接、数据
// 传出参数：错误信息
func writeJSON(conn *websocket.Conn, data interface{}) error {
	return conn.WriteJSON(data)
}

// writeBinary 写二进制
// 传入参数：websocket连接、数据
// 传出参数：错误信息
func writeBinary(conn *websocket.Conn, data interface{}) error {
	return conn.WriteMessage(websocket.BinaryMessage, data.([]byte))
}
