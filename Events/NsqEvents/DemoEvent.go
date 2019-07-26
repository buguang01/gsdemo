package NsqEvents

import (
	"GameService/Code/ConstantCode"
	"GameService/Models"

	"github.com/buguang01/Logger"
	"github.com/buguang01/util"
)

//监听用户所有日志；设置后，可以把指定用户的日志都输出到一个独立的文件
func Nsqd_ListenUser(lg *Models.NsqEventBase) {
	lg.LogicRun(func() (result int) {
		result = ConstantCode.Success
		et := lg.Data
		listen := util.NewStringAny(et["Listen"]).ToBoolV()
		if listen {
			Logger.SetListenKeyID(lg.MemberID)
		} else {
			Logger.RemoveListenKeyID(lg.MemberID)
		}
		return result
	})
}
