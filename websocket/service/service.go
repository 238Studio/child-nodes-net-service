package service

import (
	"errors"
	"time"

	_const "github.com/238Studio/child-nodes-assist/const"
	"github.com/238Studio/child-nodes-assist/util"
	"github.com/gorilla/websocket"
	jsoniter "github.com/json-iterator/go"
)

// WebsocketServiceApp Websocket 读消息
// 传入:无
// 传出:无
// 利用管道进行消息传递
func (app *WebsocketServiceApp) read() {
	app.ctx.wg.Add(1)       //wait group +1
	defer app.ctx.wg.Done() //退出时 wait group -1

	//获取websocket连接
	conn := app.conn

	//先捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, errors.New(err.(string)))
		}
	}()

	//循环读取消息
	for {
		select {
		//停止通道
		case <-app.stop:
			return

		//上下文停止通道（全局读写停止）
		case <-app.ctx.ctx.Done():
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
			app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, errors.New(err.(string)))
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
	app.ctx.wg.Add(1) //wait group +1
	defer app.ctx.wg.Done()

	//获取websocket连接
	conn := app.conn

	//先捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorChan <- util.NewError(_const.FatalException, _const.Network, errors.New(err.(string)))
		}
	}()

	for {
		select {
		case message := <-app.writeMessage:
			//检查消息类型。如果是二进制类型，则调用二进制写。否则按json格式传输。
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

		//上下文停止通道
		case <-app.ctx.ctx.Done():
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

// WebsocketServiceApp Websocket ping 定时发送ping并处理断线
// 传入:无
// 传出:无。通过管道抛出问题。
func (app *WebsocketServiceApp) ping() {
	app.ctx.wg.Add(1) //wait group +1
	defer app.ctx.wg.Done()

	//获取websocket连接
	conn := app.conn

	//先捕获异常，防止程序崩溃
	defer func() {
		if err := recover(); err != nil {
			//panic错误，定级为fatal
			app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, errors.New(err.(string)))
		}
	}()

	//设置ping定时器 单位：秒
	pingTimer := time.NewTicker(time.Duration(app.pingTimerMargin) * time.Second)
	//设置pong最大等待定时器 单位：秒
	pongMaxWait := time.NewTicker(time.Duration(app.pongWait) * time.Second)
	//退出时释放定时器
	defer func() {
		pingTimer.Stop()
		pongMaxWait.Stop()
	}()

	//设置pong handel函数
	conn.SetPongHandler(func(message string) error {
		//捕获panic
		defer func() {
			if err := recover(); err != nil {
				//panic错误，定级为fatal
				app.ErrorThrower <- util.NewError(_const.FatalException, _const.Network, errors.New(err.(string))) //FIXME：这么做是否合适？有待验证。
			}
		}()
		//重置pongMaxWait
		pongMaxWait.Reset(time.Duration(app.pongWait) * time.Second)
		return nil
	})

	//循环发送ping
	for {
		select {
		//停止通道
		case <-app.stop:
			return

		//上下文停止通道
		case <-app.ctx.ctx.Done():
			return

		//定时发送ping
		case <-pingTimer.C:
			err := conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				app.ErrorThrower <- util.NewError(_const.CommonException, _const.Network, err)
				continue //出错后继续发送
			}

		//pong超时状态
		case <-pongMaxWait.C:
			//注：所有common以上级别的错误都会启动离线机制。会试图进行断线重连。
			app.ErrorThrower <- util.NewError(_const.CommonException, _const.Network, errors.New("pong超时"))
		}
	}
}

// WebsocketServiceApp stopReadAndWrite 停止读写 阻塞等待所有goroutine退出
// 传入:无
// 传出:无
func (app *WebsocketServiceApp) stopReadAndWrite() {
	app.ctx.stop()    //停止所有goroutine
	app.ctx.wg.Wait() //等待所有goroutine退出
}
