package Flag

import (
	"flag"
	"log"
)

var usageStr = `
Usage: TigerS [options] 

Options:
	-c <url>            服务器配置文件
`
var (
	Flagc = flag.String("c", "server.json", "服务器配置文件")
)

func usage() {
	log.Fatalf(usageStr)
}

func init() {
	flag.Usage = usage
}
