package service

import (
	"context"
	"sync"

	"github.com/gorilla/websocket"
)

// WebsocketServiceApp websocket服务模型
type WebsocketServiceApp struct {
	wsURL string // websocket连接地址

	modelMessageChannels map[string]*ModelMessageChan // 模型消息通道结构体
	pingTimerMargin      int                          // ping的最大时间间隔 ms
	pongWait             int                          // pong最大返回时间间隔 ms
	conn                 *websocket.Conn              // 网络连接

	ErrorThrower chan error // 错误抛出器。读写goroutine发生panic时，将错误利用该管道抛出给上层错误管理器。

	stop chan struct{} //停止通道。利用方法进行封装。不暴露给外部。
	ctx  *ctx          //上下文。利用方法进行封装。不暴露给外部。
}

// ModelMessageChan 模型消息通道结构体
type ModelMessageChan struct {
	// 管道，用于模块(子结点)向核心传递数据 核心的模块名为键值 数组0为目标模块名 数组isByte为bool,表示是否为原始二进制数据 字段Data为数据
	//利用接口进行封装，不暴露给外部
	writeMessage chan WebsocketMessage

	// 管道，用于核心向模块传递数据 模块名为键值 字段isBytes为bool 是否为原始二进制数据 字段Data为数据
	//暴露给外部
	ReadMessage chan WebsocketMessage

	ErrorChan chan error // 错误通道

	stop chan struct{} //停止通道。利用方法进行封装。不暴露给外部。
	ctx  *ctx          //上下文。利用方法进行封装。不暴露给外部。

	conn *websocket.Conn // 网络连接
}

type ctx struct {
	ctx  context.Context
	stop context.CancelFunc

	wg sync.WaitGroup //等待所有goroutine退出
}

// WebsocketMessage websocket消息结构体
type WebsocketMessage struct {
	ModuleName string      `json:"module_name"` //模块名
	IsBytes    bool        `json:"is_bytes"`    //是否为原始二进制数据
	Data       interface{} `json:"data"`        //数据
}
