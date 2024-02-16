package ws

import (
	"github.com/238Studio/child-nodes-error-manager/errpack"
)

// InitModelMessageChan 模块管道初始化
// 传入参数：模块名
// 传出参数：模型消息通道结构体
func (app *WebsocketServiceApp) InitModelMessageChan(moduleName string) *ModelMessageChan {
	// 初始化模型消息通道结构体
	modelMessageChan := &ModelMessageChan{
		writeMessage: make(chan WebsocketMessage),
		ReadMessage:  make(chan WebsocketMessage),
		ErrorChan:    make(chan error),
		stop:         make(chan struct{}),
		conn:         app.conn,
		ctx:          app.ctx, //共享上下文
	}

	// 将模型消息通道结构体添加到模型消息通道结构体map中
	app.modelMessageChannels[moduleName] = modelMessageChan

	return modelMessageChan
}

// StartRead 开始websocket读
// 传入:无
// 传出:无
func (app *WebsocketServiceApp) StartRead() {
	go app.read()
}

// StopRead 停止websocket读
// 传入:无
// 传出:无
func (app *WebsocketServiceApp) StopRead() {
	app.stop <- struct{}{}
}

// CloseConn 关闭websocket连接
// 传入:无
// 传出:错误
func (app *WebsocketServiceApp) CloseConn() error {
	app.stopReadAndWrite() //阻塞等待读写goroutine全部退出
	err := app.conn.Close()
	if err != nil {
		return errpack.NewError(errpack.TrivialException, errpack.Network, err)
	}

	return nil
}

// StartWrite 开始写
// 传入:无
// 传出:无
func (app *ModelMessageChan) StartWrite() {
	go app.write()
}

// StopWrite 停止写
// 传入:无
// 传出:无
func (app *ModelMessageChan) StopWrite() {
	app.stop <- struct{}{}
}

// WriteMessage 写消息
// 传入:消息结构
// 传出:无
func (app *ModelMessageChan) WriteMessage(moduleName string, isBytes bool, data interface{}) {
	app.writeMessage <- WebsocketMessage{
		ModuleName: moduleName,
		IsBytes:    isBytes,
		Data:       data,
	}
}
