package webevents

import (
	"GameService/events"
	"GameService/msgmodels"
	"net/http"
)

type WebListenEvent struct {
	actionID uint32 `json:"ActionID"`
	Data     events.LogicListen
}

func (msg *WebListenEvent) GetAction() uint32 {
	return msg.actionID
}

//HTTP的回调
func (msg *WebListenEvent) HttpDirectCall(w http.ResponseWriter, req *http.Request) {
	msgmodels.WebTryRun(msg, msg.Data.Hander, w, req)
}
