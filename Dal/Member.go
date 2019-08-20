package Dal

/*
这是一个例子；
*/
import (
	"GameService/Service"
	"database/sql"
	"fmt"
	"time"

	"github.com/buguang01/bige/event"
)

/*
bige是框架中用来生成SQL的tag
select 是表示会做为查询的条件
bigekey 是表示在更新的时候，不做为更新的字段；但是在插入的时候，会做为数据写入；
- 是表示这个字段不参与保存
*/
type MemberMD struct {
	MemberID   int       `bige:"memberid,bigekey"` //用户ID
	UserName   string    `bige:"username"`         //
	Pwd        string    `bige:"pwd"`              //
	DriveID    string    `bige:"driveid"`          //
	OStype     string    `bige:"ostype"`           //
	CreateIP   string    `bige:"createip"`         //
	PlatFormID string    `bige:"platformid"`       //
	ChanID     string    `bige:"chanid"`           //渠道ID
	OpenID     string    `bige:"openid"`           //用户唯一标识
	UnionID    string    `bige:"unionid"`          //同一用户，对同一个微信开放平台下的不同应用，unionid是相同的
	CreateTime time.Time `bige:"createtime"`       //
	LoginTime  time.Time `bige:"logintime"`        //
	BanTime    time.Time `bige:"bantime"`          //
	ServerID   int       `bige:"serverid"`         //
}

//表名
func (this *MemberMD) GetTableName() string {
	return "member"
}

//保存参数列表
func (this *MemberMD) ParmArray() []interface{} {
	return []interface{}{this.MemberID, this.UserName, this.Pwd, this.DriveID, this.OStype, this.CreateIP,
		this.PlatFormID, this.ChanID, this.OpenID, this.UnionID, this.CreateTime, this.LoginTime, this.BanTime, this.ServerID}
}

//查询的参数列表
func (this *MemberMD) QueryArray() []interface{} {
	return []interface{}{}
}

func (md *MemberMD) LoadDB(rows *sql.Rows) {
	rows.Scan(
		&md.MemberID, &md.UserName, &md.Pwd, &md.DriveID,
		&md.OStype, &md.CreateIP, &md.PlatFormID, &md.ChanID,
		&md.OpenID, &md.UnionID, &md.CreateTime, &md.LoginTime,
		&md.BanTime, &md.ServerID)
}

func (md *MemberMD) ToJson() event.JsonMap {
	resultjs := make(event.JsonMap)
	resultjs["MemberID"] = md.MemberID
	resultjs["UserName"] = md.UserName
	// resultjs["ChanID"] = md.ChanID
	// resultjs["DriveID"] = md.DriveID
	// resultjs["OStype"] = md.OStype
	// resultjs["CreateIP"] = md.CreateIP
	// resultjs["PlatFormID"] = md.PlatFormID
	// resultjs["OpenID"] = md.OpenID
	resultjs["CreateTime"] = md.CreateTime.Unix()
	resultjs["BanTime"] = md.BanTime.Unix()
	resultjs["ServerID"] = md.ServerID
	return resultjs
}

func (this *MemberMD) Clone() *MemberMD {
	result := *this

	return &result
}

func MemberGetAll(db *sql.DB) *sql.Rows {
	return event.DataGetByID(db, &MemberMD{})
}

func MemberGetMaxMemberID(db *sql.DB) int {
	sqlstr := `
SELECT IFNULL(MAX(memberid),0) FROM member;
	`
	var result int = 0
	row := db.QueryRow(sqlstr)
	row.Scan(&result)
	return result
}

//UpDBMemberMD 写入用户信息
func UpDBMemberMD(md *MemberMD) {
	upmd := new(DalModel)
	upmd.KeyID = md.MemberID
	upmd.DataDBModel = md.Clone()
	upmd.UpTime = SAVE_LV0
	upmd.DataKey = fmt.Sprintf("member%d", md.MemberID)
	Service.DBEx.AddMsg(upmd)
}
