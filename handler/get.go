package handler

import (
	"encoding/json"
	"file_share_server/dao"

	"github.com/go75/gol"
	"github.com/go75/tcpx"
)

func GetAllFileInfo(req *tcpx.Request) {
	files := dao.GetAllShareFile()
	if len(files) == 0 {
		gol.Warn("未读取到共享文件")
		return
	}

	data, err := json.Marshal(files)
	if err != nil {
		gol.Errorln("数据序列化失败:", err)
		return
	}

	err = req.Conn().Send(req.MsgID(), data)
	if err != nil {
		gol.Errorln("发送数据给客户端失败:", err)
		return
	}
}
