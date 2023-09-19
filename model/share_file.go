package model

type ShareFile struct {
	// 文件分享者的名字
	ShareUser string
	// 文件的路径
	Path string
	// 文件大小
	Size  uint32
	// 分享时间
	ShareTime int64
	// 分享者的网络地址
	ShareAddress string
}