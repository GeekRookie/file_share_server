package global

import (
	"file_share_server/model"
	"sync"
)

// k:分享者的名字+" "+文件路径, v:分享的文件对象
var FileInfos = make(map[string]*model.ShareFile)
var FileLock = new(sync.RWMutex)