package childWebsocket_test

import (
	"testing"

	childWebsocket "github.com/UniversalRobotDriveTeam/child-nodes-net-service/websocket"
)

func TestWebsocket(t *testing.T) {
	wsAPP, err := childWebsocket.Init("ws://127.0.0.1:8080/ws", 1000, 1000)
	if err != nil {
		t.Fatal(err)
	}

	wsAPP.StartRead()

	app := wsAPP.InitModelMessageChan("test")

	app.StartWrite()
	app.WriteMessage("test", false, "测试消息")
}
