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
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:7890"))
	if err != nil {
		panic(err)
	}

	username := "bob"

	_, err = conn.Write([]byte(username))
	if err != nil {
		panic(err)
	}

	res := make([]byte, 1)

	_, err = io.ReadFull(conn, res)
	if err != nil {
		panic(err)
	}

	if res[0] == 0 {
		// 失败
		panic("连接初始化失败")
	}

	data, err := json.Marshal(model.ShareFile{
		ShareUser: username,
		Path: "/",
		Size: 100,
		ShareTime: 10,
		ShareAddress: "127.0.0.1:8888",
	})

	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))
	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Set, data)))
	if err != nil {
		panic(err)
	}

	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Get, nil)))
	if err != nil {
		panic(err)
	}


	fileListData := make([]byte, 4096)

	n, err := conn.Read(fileListData)
	if err != nil {
		panic(err)
	}

	var fileList = new([]*model.ShareFile)
	fmt.Println(string(fileListData[3:n]))
	err = json.Unmarshal(fileListData[3:n], fileList)
	if err != nil {
		panic(err)
	}

	fileInfo := (*fileList)[0]

	downloadConn, err := net.Dial("tcp", fileInfo.ShareAddress)
	if err != nil {
		panic(err)
	}

	_, err = downloadConn.Write([]byte(fileInfo.Path))
	if err != nil {
		panic(err)
	}

	fmt.Println("写入成功")

	fileData := make([]byte, 4096)
	n, err = downloadConn.Read(fileData)
	if err != nil {
		panic(err)
	}

	fmt.Printf("recv from %s:%s data:%s\n", fileInfo.ShareUser, fileInfo.ShareAddress, string(fileData))

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Del, data)))
	if err != nil {
		panic(err)
	}
}