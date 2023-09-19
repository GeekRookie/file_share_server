package handler

import (
	"encoding/json"
	"file_share_server/common"
	"file_share_server/dao"
	"file_share_server/pkg/e"
	"file_share_server/global"
	"file_share_server/model"

	"github.com/go75/gol"
	"github.com/go75/tcpx"
)

func SetFileInfo(req *tcpx.Request) {
	var file = new(model.ShareFile)
	err := json.Unmarshal(req.Data(), file)
	if err != nil {
		err = req.Conn().Send(1, []byte("文件信息反序列化失败:"+err.Error()))
		if err != nil {
			gol.Errorln(e.ConnWriteErr(err))
			return
		}
	}

	gol.Infof("文件序列化成功:%v", file)

	db := dao.CreateShareFile(file)
	if db.RowsAffected != 1 {
		gol.Errorln("共享文件创建失败:", err)
		return
	}

	_, err = global.Broadcaster.Write(append([]byte{common.Get}, req.Data()...))
	if err != nil {
		gol.Errorln("广播失败:", err)
	}
}