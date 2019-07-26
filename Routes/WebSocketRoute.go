package Routes

import (
	"GameService/Code/ActionCode"
	"GameService/Events/WebSocketEvents"
	"GameService/Service"

	"github.com/buguang01/bige/event"
	"golang.org/x/net/websocket"
)

func init() {
	WSRoutelist = make(map[int]event.WebSocketCall)
	WSRoutelist[ActionCode.Ws_User_Test] = WebSocketEvents.WsDomeEvent
}

var (
	WSRoutelist map[int]event.WebSocketCall
)

func WebSocketInit() {
	Service.WebSocketEx.RouteFun = WebSocketRoute
	Service.WebSocketEx.WebSocketOnlineFun = WebSocketOnlineFun
}

func WebSocketRoute(code int) event.WebSocketCall {
	f, ok := WSRoutelist[code]
	if ok {
		return f
	}
	return nil
}

func WebSocketOnlineFun(conn *websocket.Conn) string {
	req := conn.Request()
	// Logger.PInfo("%v", req.Header)
	//这个方法是用来拿IP的，因为会被https代理，所以RemoteAddr不一定拿到客户的IP；
	//所以与你自己的运营沟通一下看看在哪里可以拿到IP；
	if ips, ok := req.Header["X-Forwarded-For"]; ok {
		if len(ips) > 0 {
			return ips[0]
		}
	}
	return req.RemoteAddr
}
