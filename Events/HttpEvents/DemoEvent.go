package HttpEvents

import (
	"net/http"

	"github.com/buguang01/bige/event"
)

func DemoEvent(et event.JsonMap, w http.ResponseWriter) {
	//这里写收到http消息的处理逻辑
}
