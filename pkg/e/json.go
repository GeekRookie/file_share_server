package e

func MarshalErr(err error) string {
	return "数据序列化失败:" + err.Error()
}

func Unmarshal(err error) string {
	return "数据反序列化失败:" + err.Error()
}