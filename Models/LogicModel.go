package Models

import (
	"GameService/Code/ConstantCode"
	"GameService/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/threads"
	"github.com/buguang01/util"
)

/*
下面是跑在LogicMoudle协程下的模块
一般从别的模块收到消息后，要生成一个对应的Logic数据
然后扔到LogicMoudle协程上进行处理
你可以为不同消息来源去写对应的Logic模型
*/

type LogicBase struct {
	MemberID int //用户ID
	// user     *userobj.UserModel    //用户对象
	// wsmd     *event.WebSocketModel //连接对象
	JsMsg event.JsonMap //收到的消息
}

func (this *LogicBase) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)

			// event.WebSocketReplyMsg(this.wsmd, this.et, result, jsdata)
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
			// event.WebSocketReplyMsg(this.wsmd, this.et, result, nil)
		}, nil,
	)
	_ = result
}

func (this *LogicBase) KeyID() string {
	if Service.Sconf.LogicConf.InitNum == 0 {
		return util.NewStringInt(this.MemberID).ToString()
	}
	return util.NewStringInt(this.MemberID % Service.Sconf.LogicConf.InitNum).ToString()
}

//userLogicRun 用户逻辑运行
func (this *LogicBase) UserLogicRun(f func() int) int {
	result := ConstantCode.Success
	threads.Try(func() {
		result = f()
	}, func(err interface{}) {
		Logger.PFatal(err)
		result = ConstantCode.User_DB_Error
		// UserDbErrorHander(this.user)
		/*
			这里是如果传入的方法运行出错时，会给自己发一个同一的消息
			比如发一个重新加载用户数据的消息，用来清除用户内存的脏数据；
		*/

	}, nil)
	return result
}

//自动任务的逻辑模型
type LogicAuto struct {
	LogicBase
}

func (this *LogicAuto) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)
			if result == ConstantCode.Success {
				// this.user.PushWebSocket(jsdata)
				/*
					如果与用户的通信是建立在websocket模式下的，那可以使用这个方法给用户回复
				*/
			}
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
		}, nil,
	)
}

//不回复用户的逻辑处理模型
type LogicNotResult struct {
	LogicBase
}

func (this *LogicNotResult) LogicRun(f func(event.JsonMap) int) {
	jsdata := make(event.JsonMap)
	result := ConstantCode.NotLogic
	threads.Try(
		func() {
			result = f(jsdata)
			if result != ConstantCode.Success {
				// event.WebSocketReplyMsg(this.wsmd, this.et, result, jsdata)
			}
		}, func(err interface{}) {
			Logger.PFatal(err)
			result = ConstantCode.LOGIC_ERROR
			// event.WebSocketReplyMsg(this.wsmd, this.et, result, nil)
		}, nil,
	)

}
