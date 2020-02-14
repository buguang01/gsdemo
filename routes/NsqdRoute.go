package routes

import (
	"GameService/actioncode"
	"GameService/events/nsqevents"

	"github.com/buguang01/bige/messages"
)

var (
	NsqdRoute = messages.JsonMessageHandleNew()
)

func init() {
	NsqdRoute.SetRoute(actioncode.Nsqd_ListenUser, &nsqevents.NsqdListenEvent{})
}
