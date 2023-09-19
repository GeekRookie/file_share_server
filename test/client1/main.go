package main

import (
	"encoding/json"
	"file_share_server/common"
	"file_share_server/model"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"

	"github.com/go75/tcpx"
)

// test1是下载文件的
func main() {
	// 连接远程服务端
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:7890"))
	if err != nil {
		panic(err)
	}

	// 当前客户端的名称, 局域网在线的客户端名称唯一的
	username := "bob"

	// 告诉服务端, 客户端自己叫啥
	_, err = conn.Write([]byte(username))
	if err != nil {
		panic(err)
	}

	// 接受服务端的回应
	res := make([]byte, 1)
	_, err = io.ReadFull(conn, res)
	if err != nil {
		panic(err)
	}

	// 如果为0, 证明客户端上线失败, 可能是在线的用户已经有人叫这个名字了, 也有可能是网络问题等等
	if res[0] == 0 {
		// 失败
		panic("连接初始化失败")
	}

	// 询问服务端当前局域网有哪些分享的文件
	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Get, nil)))
	if err != nil {
		panic(err)
	}

	// 读取服务端返回的局域网中分享的文件列表
	fileListData := make([]byte, 4096)
	n, err := conn.Read(fileListData)
	if err != nil {
		panic(err)
	}

	// 文件列表反序列化
	var fileList = new([]*model.ShareFile)
	fmt.Println(string(fileListData[3:n]))
	err = json.Unmarshal(fileListData[3:n], fileList)
	if err != nil {
		panic(err)
	}

	// 拿到client0分享的文件信息
	fileInfo := (*fileList)[0]

	// 和client0建立连接
	downloadConn, err := net.Dial("tcp", fileInfo.ShareAddress)
	if err != nil {
		panic(err)
	}

	// 告诉client0我要获取哪个文件
	_, err = downloadConn.Write([]byte(fileInfo.Path))
	if err != nil {
		panic(err)
	}

	// 读取client0返回的该文件内容
	fileData := make([]byte, 4096)
	n, err = downloadConn.Read(fileData)
	if err != nil {
		panic(err)
	}

	// 这里把文件内容打印了, 应该是让客户端选择保存在哪个位置
	fmt.Printf("recv from %s:%s data:%s\n", fileInfo.ShareUser, fileInfo.ShareAddress, string(fileData))

	// 阻塞, 等待ctrl + c退出
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit
}