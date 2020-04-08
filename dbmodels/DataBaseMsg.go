package dbmodels

import (
	"runtime"

	"github.com/buguang01/bige/model"
)

type SaveDBfunc func(model.IConnDB) error

type DataBaseMsg struct {
	ThreadID int //所在DB协程
	DataKey  string
	RunFunc  SaveDBfunc //调用方法
}

//所在DB协程
func (this *DataBaseMsg) DBThreadID() int {
	//默认按CPU个数的十倍的协程数来分配对应的协程进行处理
	//分配时，按用户ID进行取余
	cpu := runtime.NumCPU() * 10
	return this.ThreadID % cpu

}

/*数据表,如果你的表放入时，不是马上保存的，那么后续可以用这个KEY来进行覆盖，
这样就可以实现多次修改一次保存的功能
所以这个字段建议是：用户ID+数据表名+数据主键
*/
func (this *DataBaseMsg) GetDataKey() string {
	return this.DataKey
}

func (this *DataBaseMsg) SaveDB(conn model.IConnDB) error {
	return this.RunFunc(conn)
}

func NewDbMsg(threadid int, datakey string, f SaveDBfunc) *DataBaseMsg {
	return &DataBaseMsg{
		ThreadID: threadid,
		DataKey:  datakey,
		RunFunc:  f,
	}
}
