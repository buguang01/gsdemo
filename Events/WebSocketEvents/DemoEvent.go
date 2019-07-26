package WebSocketEvents

import (
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
)

//WsDomeEvent
// et 是收到消息的Json对象
// wsmd 是websocket的连接对象
// runobj 为这个连接建的协程管理器，会在连接关闭时关闭
func WsDomeEvent(et event.JsonMap, wsmd *event.WebSocketModel, runobj *threads.ThreadGo) {
	//WebSocket的消息处理

}
