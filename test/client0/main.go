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
	conn, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:7890"))
	if err != nil {
		panic(err)
	}

	username := "jack"

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

	go func() {
		lis, err := net.Listen("tcp", "127.0.0.1:9999")
		if err != nil {
			panic(err)
		}
		for {
			uploadConn, err := lis.Accept()
			if err != nil {
				panic(err)
			}
			filePath := make([]byte, 1024)
			n, err := uploadConn.Read(filePath)
			if err != nil {
				panic(err)
			}

			fmt.Println("要写入文件的路径:", string(filePath))

			filedata, err := os.ReadFile(string(filePath[:n]))
			if err != nil {
				panic(err)
			}
			
			_, err = uploadConn.Write(filedata)
			if err != nil {
				panic(err)
			}
		}
	}()

	data, err := json.Marshal(model.ShareFile{
		ShareUser: username,
		Path: "/home/v/cv_debug.log",
		Size: 100,
		ShareTime: 10,
		ShareAddress: "127.0.0.1:9999",
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

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt)
	<-exit

	_, err = conn.Write(tcpx.Pack(tcpx.NewMessage(common.Del, data)))
	if err != nil {
		panic(err)
	}
}
