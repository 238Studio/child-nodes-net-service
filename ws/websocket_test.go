package ws_test

import (
	"testing"

	"github.com/238Studio/child-nodes-net-service/ws"
)

func TestWebsocket(t *testing.T) {
	wsAPP, err := ws.InitWebsocketService("ws://localhost:8080/ws", 100000, 100000)
	if err != nil {
		t.Fatal(err)
	}

	wsAPP.StartRead()

	app := wsAPP.InitModelMessageChan("test")

	app.StartWrite()
	for i := 1; i <= 10; i++ {
		app.WriteMessage("test", false, "测试消息")

		t.Log(<-app.ReadMessage)
	}
}

func TestPongHandel(t *testing.T) {
	wsAPP, err := ws.InitWebsocketService("ws://127.0.0.1:8080/ws", 1, 5)
	if err != nil {
		t.Fatal(err)
	}

	for {
		select {
		case err := <-wsAPP.ErrorThrower:
			t.Log(err)
		}
	}
}
