package initialize

import (
	"file_share_server/global"

	"github.com/go75/gol"
)

func initLog() {
	err := gol.SetLogFile(global.Config.Log.Location)
	if err != nil {
		panic(err)
	}

	gol.SetLogLevel(gol.DebugLevel)
}