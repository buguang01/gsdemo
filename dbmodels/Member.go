package dbmodels

import (
	"fmt"

	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/messages"
	"github.com/buguang01/bige/model"
)

type Member struct {
	MemberID int    `bige:"memberid,bigekey,select"` //用户ID
	UserName string `bige:"username"`                //用户名字
	Other    string `bige:"other"`                   //其他字段
}

//这种是简单的全字段保存
func (md *Member) save(conn model.IConnDB) error {
	sqlstr := event.MarshalUpSql(md, "member")
	_, err := conn.Exec(sqlstr, md.MemberID, md.UserName, md.Other)
	return err
}

//生成保存的db消息
func (md *Member) InitSaveMsg() messages.IDataBaseMessage {
	msg := NewDbMsg(md.MemberID, fmt.Sprintf("%d_member_%d", md.MemberID, md.MemberID), md.save)
	// services.DBEx.AddMsg(msg)
	return msg
}
