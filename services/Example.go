package services

import (
	"time"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/modules"
)

var (
	Sconf   *ServiceConf
	MysqlEx *model.MysqlAccess   //mysql管理器
	RedisEx *model.RedisAccess   //redis管理器
	GameEx  *modules.GameService //系统模块管理器

	DBEx        *modules.DataBaseModule  //DB操作模块
	LogicEx     *modules.LogicModule     //逻辑操作模块
	TaskEx      *modules.AutoTaskModule  //内存管理模块
	NsqdEx      *modules.NsqdModule      //nsq消息队列通信模块
	WebEx       *modules.WebModule       //HTTP通信模块
	WebSocketEx *modules.WebSocketModule //ws通信模块
)

type ServiceConf struct {
	ServiceID   int             //游戏服务器ID
	PStatusTime time.Duration   //打印状态的时间（秒）
	LogLv       Logger.LogLevel //写日志的等级
	LogPath     string          //日志写目录
	LogMode     Logger.LogMode  //日志模式
	LogServerID string          //log服务器ID

	// MysqlConf      MysqlConf //mysql配置
	DBConf         model.MysqlConfigModel //Mysql管理器
	RedisConf      RedisConf              //redis配置
	NsqdAddr       []string               //nsqd地址组
	NsqLookupdAddr []string               //lookup 地址组
	WebAddr        string                 //Web的地址
	WsAddr         string                 //Websocket 的地址

}

type MysqlConf struct {
	Addr       string
	User       string
	Pwd        string
	DBName     string
	MaxOpenNum int //最大连接数
	IdleNum    int //最大待机连接数

}

type RedisConf struct {
	Addr        string        //连接字符串
	Indexdb     int           //默认DB号
	Auth        string        //连接密码
	MaxIdle     int           //最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态
	MaxActive   int           //最大的激活连接数，表示同时最多有N个连接 ，为0事表示没有限制
	IdleTimeout time.Duration //最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭(秒)
}
