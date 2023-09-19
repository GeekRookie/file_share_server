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

// test0是上传文件的
func main() {
	// 连接远程服务端
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:7890"))
	if err != nil {
		panic(err)
	}

	// 当前客户端的名称, 局域网在线的客户端名称唯一的
	username := "jack"

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

	// 我们要分享的文件信息
	data, err := json.Marshal(model.ShareFile{
		ShareUser: username, // 分享人, 当前客户端
		Path: "/home/v/cv_debug.log", // 分享的文件在我们本机的路径
		Size: 100, // 文件大小, 我瞎填的, 这个就是一个提示作用, 不影响正常的业务逻辑
		ShareTime: 10,  // 该文件分享的时间戳, 我瞎填的, 这个就是一个提示作用, 不影响正常的业务逻辑
		ShareAddress: "127.0.0.1:9999", // 我们分享文件的网络地址, 等待别的客户端连接
	})

	if err != nil {
		panic(err)
	}

	// 发送服务器, 表明我们分享这个文件
	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Set, data)))
	if err != nil {
		panic(err)
	}

	// 启动监听服务, 等待别的客户端来获取当前客户端分享的文件 (p2p传输文件)
	go func() {
		// 监听
		lis, err := net.Listen("tcp", "127.0.0.1:9999")
		if err != nil {
			panic(err)
		}
		for {
			// 等待别的客户端连接
			uploadConn, err := lis.Accept()
			if err != nil {
				panic(err)
			}

			// 连接建立成功后, 读取对方请求的文件路径
			filePath := make([]byte, 1024)
			n, err := uploadConn.Read(filePath)
			if err != nil {
				panic(err)
			}

			// 读取这个文件的内容, 这个步骤要判断一下对方请求的文件在不在当前客户端的文件分享列表中, 为了省事, 我暂时没弄
			filedata, err := os.ReadFile(string(filePath[:n]))
			if err != nil {
				panic(err)
			}

			// 将读取到文件内容发送给对方客户端, 这里直接一次传输了, 之后可升级多次传输
			_, err = uploadConn.Write(filedata)
			if err != nil {
				panic(err)
			}
		}
	}()

	// 阻塞, 等待ctrl + c退出
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	// 取消共享的文件, 或者服务端连接断开就把属于这个用户分享的文件全部删除(这个还没写, 所以现在暂时客户端手动删除)
	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Del, data)))
	if err != nil {
		panic(err)
	}
}
