package msgmodels

import "github.com/buguang01/bige/messages"

type NsqResult struct {
	messages.NsqdMessage
	//数据
	Data interface{}
}

func (msg *NsqResult) GetAction() uint32{
	return msg.ActionID
}

//自定义结构的返回
func NewNsqResult(msg messages.INsqMessageHandle,v interface{}) (result messages.INsqdResultMessage){
	md:=new(NsqResult)
md.ActionID=msg.GetAction()
md.SendSID=msg.GetTopic()
md.SendUserID=msg.GetSendUserID()
	md.Data=v
	result=md

}
//json数据的返回
func NewNsqResultJSON(msg messages.INsqMessageHandle,v H)(messages.INsqdResultMessage){
	md:=new(NsqResult)
	md.ActionID=msg.GetAction()
	md.SendSID=msg.GetTopic()
	md.SendUserID=msg.GetSendUserID()
		md.Data=v
		return md
}