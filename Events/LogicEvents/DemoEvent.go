package LogicEvents

import (
	"GameService/Code/ActionCode"
	"GameService/Code/ConstantCode"
	"GameService/Models"
	"GameService/Service"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/util"
)

type Exped_Logic_Action struct {
	Models.LogicBase
	ActionID int //消息号
	// AID      int    //冒险槽ID
	// PtypeID  int    //冒险类型1：一般挂机；2：英雄挂机;3:BOSS挂机
	// HeroLi   []int  //英雄列表
	UseID int //花钱类型
	// HookID   int    //挂机ID
	// Num      int    //编号
	// TeamID   int    //编队ID
	// TeamV    string //编队信息
}

func NewExped_Logic_Action(mid int, et event.JsonMap) { //(user *userobj.UserModel, wsmd *event.WebSocketModel, et event.JsonMap) {
	logicmd := new(Exped_Logic_Action)
	logicmd.MemberID = mid
	// logicmd.user = user
	// logicmd.wsmd = wsmd
	logicmd.JsMsg = et
	logicmd.ActionID = et.GetAction()
	switch logicmd.ActionID {
	case ActionCode.Ws_User_Test:
		logicmd.UseID = util.NewStringAny(et["UseID"]).ToIntV()
	default:
		// event.WebSocketReplyMsg(wsmd, et, ConstantCode.NotAction, nil)
		return
	}
	Service.LogicEx.AddMsg(logicmd)
}
func (this *Exped_Logic_Action) Run() {
	switch this.ActionID {
	case ActionCode.Ws_User_Test:
		this.LogicRun(this.exped_Slot_Open)
	default:
		Logger.PError(nil, "Not Action:%d", this.ActionID)
	}
}

func (this *Exped_Logic_Action) exped_Slot_Open(jsuser event.JsonMap) (result int) {
	// user := this.user
	result = ConstantCode.Success
	//检查消息的处理条件

	// user.ObjLock.Lock()
	// defer user.ObjLock.Unlock()
	// if user.IsUserStatus(userobj.USER_STATUS_UNLOAD) {
	// 	return ConstantCode.User_STATUS_NOT
	// }
	// args := NewUserArgs(user, util.PrintMyName())
	// uid := len(user.AvgList) + 1
	// avgconf, ok := Conf.ConfExample.GetAvgByUID(uid)
	// if !ok {
	// 	return ConstantCode.Exped_OpenMax
	// }
	// delbc := util.NewBaseDataString("")
	// if this.UseID == 0 {
	// 	delbc.UpDataBc(avgconf.OpenBc1, nil)
	// 	if avgconf.UserLv > user.UserData.Lv {
	// 		return ConstantCode.Player_LvEnough
	// 	}
	// } else {
	// 	delbc.UpDataBc(avgconf.OpenBc2, nil)
	// }
	// if ok, code := user.ContainsItems(delbc); !ok {
	// 	return code
	// }

	//开始处理，当然 里面也有可能还有检查，但如果在这里面出现异常应该要重新加载数据等操作
	result = this.UserLogicRun(func() int {
		// args.DelItem(delbc)
		// avgmd := new(Dal.AvgSlotMD)
		// avgmd.Init(user.MemberID(), uid)
		// user.AvgList[avgmd.AvgID] = avgmd
		// args.GetAvgSlotByID(uid)
		// if tmd, ok := user.AvgTeamList[uid]; !ok {
		// 	tmd = new(Dal.AvgTeamItemMD)
		// 	tmd.Init(user.MemberID(), uid)
		// 	user.AvgTeamList[tmd.TeamID] = tmd
		// 	args.GetAvgTeamByID(tmd.TeamID)
		// }
		// args.UpData()
		return result
	})
	if result == ConstantCode.Success {
		//填充要回复的数据信息

		// args.ToJson(jsuser)
	}
	return result
}
