package router

import (
	"file_share_server/common"
	"file_share_server/handler"

	"github.com/go75/tcpx"
)


func Init(engine *tcpx.Engine) {
	engine.Regist(common.Get, handler.GetAllFileInfo)
	engine.Regist(common.Set, handler.SetFileInfo)
	engine.Regist(common.Del, handler.DelFileInfo)
}