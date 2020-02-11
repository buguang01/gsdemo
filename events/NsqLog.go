package events

import (
	"github.com/buguang01/util"
)

func Nlog_Register(user *userobj.UserModel) {
	dt := util.GetCurrTime()
	logmd := NewLogInfoByNum("PassPort", "Register", "",
		user.MemberID(), Service.Sconf.GameConf.ServiceID, dt,
		0, user.Member.CreateTime.Unix())
	sendNsqdLogMD(logmd)
}