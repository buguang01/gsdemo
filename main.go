package main

import (
	"GameService/flags"
	"GameService/routes"
	"GameService/services"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/buguang01/Logger"
	"github.com/buguang01/bige/model"
	"github.com/buguang01/bige/modules"
	"github.com/buguang01/util"
)

func main() {
	services.Sconf = new(services.ServiceConf)

	if !flag.Parsed() {
		flag.Parse()
	}

	f, err := os.Open(*flags.Flagc)
	if err != nil {
		panic(err)
	}
	b, _ := ioutil.ReadAll(f)
	f.Close()

	json.Unmarshal(b, services.Sconf)
	Logger.Init(services.Sconf.LogLv, services.Sconf.LogPath, services.Sconf.LogMode)
	defer Logger.LogClose()

	services.MysqlEx = model.NewMysqlAccess(&services.Sconf.DBConf)
	defer services.MysqlEx.Close()
	if err := services.MysqlEx.Ping(); err != nil {
		Logger.PError(err, "")
		return
	}
	services.RedisEx = model.NewRedisAccess(
		model.RedisSetAddr(services.Sconf.RedisConf.Addr),
		model.RedisSetAuth(services.Sconf.RedisConf.Auth),
		model.RedisSetIndexDB(services.Sconf.RedisConf.Indexdb),
		model.RedisSetMaxActive(services.Sconf.RedisConf.MaxActive),
		model.RedisSetMaxIdle(services.Sconf.RedisConf.MaxIdle),
	)
	defer services.RedisEx.Close()
	c := services.RedisEx.GetConn()
	if err := c.Err(); err != nil {
		Logger.PError(err, "")
		return
	}
	c.Close()

	services.DBEx = modules.NewDataBaseModule(services.MysqlEx.GetDB())
	services.LogicEx = modules.NewLogicModule()
	services.TaskEx = modules.NewAutoTaskModule()
	services.NsqdEx = modules.NewNsqdModule(
		modules.NsqdSetPorts(services.Sconf.NsqdAddr...),
		modules.NsqdSetLookup(services.Sconf.NsqLookupdAddr...),
		modules.NsqdSetMyTopic(util.ToString(services.Sconf.ServiceID)),
		modules.NsqdSetMyChannelName(fmt.Sprintf("chancel_%d", services.Sconf.ServiceID)),
		modules.NsqdSetRoute(routes.NsqdRoute),
	)
	services.WebEx = modules.NewWebModule(
		modules.WebSetIpPort(services.Sconf.WebAddr),
		modules.WebSetRoute(routes.WebRoute),
		modules.WebSetTimeoutFunc(routes.WebTimeout),
	)
	services.WebSocketEx = modules.NewWebSocketModule(
		modules.WebSocketSetIpPort(services.Sconf.WsAddr),
		modules.WebSocketSetRoute(routes.WebSocketRoute),
		modules.WebSocketSetOnlineFun(routes.WebScoketOnline),
	)
	services.GameEx = modules.NewGameService(
		modules.GameServiceSetSID(101),
	)
	services.GameEx.ServiceStopHander = ServiceStopHander
	services.GameEx.ServiceStartHander = ServiceStartHander
	services.GameEx.AddModule(
		services.DBEx,
		services.LogicEx,
		services.TaskEx,
		services.NsqdEx,
		services.WebEx,
		services.WebSocketEx,
	)
	services.GameEx.Run()

}

//当服务器被关掉时，先调用的方法
func ServiceStopHander() {

}

//当服务器所有服务都启动后，先调用的方法
func ServiceStartHander() {

}
