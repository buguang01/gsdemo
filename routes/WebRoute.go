package routes

import (
	"net/http"

	"github.com/buguang01/bige/messages"
)

var (
	WebRoute = messages.HttpJsonMessageHandleNew()
)

func init() {

}

func WebTimeout(webmsg messages.IHttpMessageHandle,
	w http.ResponseWriter, req *http.Request) {
	//超时处理，这可以这里做统一处理
}
