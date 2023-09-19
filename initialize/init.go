package initialize

import (
	"file_share_server/router"
	"github.com/go75/tcpx"
)

func Init() *tcpx.Engine {
	initConfig()
	initDB()
	initLog()
	initBroadcast()
	engine := initServer()

	router.Init(engine)

	return engine
}