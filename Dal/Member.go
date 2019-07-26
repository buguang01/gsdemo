package Dal

/*
这是一个例子；
*/
import (
	"GameService/Service"
	"database/sql"
	"fmt"
	"time"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/event"
	"github.com/buguang01/bige/model"
)

type MemberMD struct {
	MemberID   int       //用户ID
	UserName   string    //
	Pwd        string    //
	DriveID    string    //
	OStype     string    //
	CreateIP   string    //
	PlatFormID string    //
	ChanID     string    //渠道ID
	OpenID     string    //用户唯一标识
	UnionID    string    //同一用户，对同一个微信开放平台下的不同应用，unionid是相同的
	CreateTime time.Time //
	LoginTime  time.Time
	BanTime    time.Time //
	ServerID   int       //
}

func (md *MemberMD) LoadDB(rows *sql.Rows) {
	rows.Scan(
		&md.MemberID, &md.UserName, &md.Pwd, &md.DriveID,
		&md.OStype, &md.CreateIP, &md.PlatFormID, &md.ChanID,
		&md.OpenID, &md.UnionID, &md.CreateTime, &md.LoginTime,
		&md.BanTime, &md.ServerID)
}

func (md *MemberMD) ToJson() map[string]interface{} {
	resultjs := make(map[string]interface{})
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

func MemberGetAll(db *sql.DB) *sql.Rows {
	sqlstr := `
SELECT 
 memberid
,username
,pwd
,driveid
,ostype
,createip
,platformid
,chanid
,openid
,unionid
,createtime
,logintime
,bantime
,serverid
FROM member;
`
	read, err := db.Query(sqlstr)
	if err != nil {
		Logger.PFatal(err)
		panic(err)
	}
	return read
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

func MemberUpData(conndb model.IConnDB, datamd event.DataDBModel) error {
	md := datamd.(*MemberMD)
	sqlstr := `
	INSERT INTO member(
		memberid,username,pwd,driveid
		,ostype,createip,platformid,chanid
		,openid,unionid,createtime,logintime
		,bantime,serverid
	) VALUES(
		?,?,?,? ,?,?,?,? ,?,?,?,?	,?,?
	)
	ON DUPLICATE KEY UPDATE
		username=values(username),
		pwd=values(pwd),
		driveid=values(driveid),
		ostype=values(ostype),
		createip=values(createip),
		platformid=values(platformid),
		chanid=values(chanid),
		openid=values(openid),
		unionid=values(unionid),
		createtime=values(createtime),
		logintime=values(logintime),
		bantime=values(bantime),
		serverid=values(serverid)

	;
	`
	_, err := conndb.Exec(
		sqlstr,
		md.MemberID, md.UserName, md.Pwd, md.DriveID,
		md.OStype, md.CreateIP, md.PlatFormID, md.ChanID,
		md.OpenID, md.UnionID, md.CreateTime, md.LoginTime,
		md.BanTime, md.ServerID)
	if err != nil {
		return err
	}
	return nil
}

//UpDBMemberMD 写入用户信息
func UpDBMemberMD(md *MemberMD) {
	upmd := new(DalModel)
	upmd.KeyID = md.MemberID
	tmp := *md
	upmd.DataDBModel = &tmp
	upmd.SaveFun = MemberUpData
	upmd.UpTime = SAVE_LV0
	upmd.DataKey = fmt.Sprintf("member%d", md.MemberID)
	Service.DBEx.AddMsg(upmd)
}
