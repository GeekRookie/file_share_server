package e

func ConnWriteErr(err error) string {
	return "发送数据给客户端失败:" + err.Error()
}