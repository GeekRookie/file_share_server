package initialize

import (
	"file_share_server/global"
	"file_share_server/pkg/e"
	"fmt"
	"time"

	"github.com/go75/gol"
	"github.com/go75/tcpx"
)

func initServer() *tcpx.Engine {
	engine, err := tcpx.New("tcp4", fmt.Sprintf("%s:%d", global.Config.Server.IP, global.Config.Server.Port), 16, 1024, time.Microsecond << 4, prehook, posthook, nil)
	if err != nil {
		panic(err)
	}
	return engine
}

func prehook(c *tcpx.Connection) {
	name := make([]byte, 1024)
	n, err := c.Conn.Read(name)
	if err != nil {
		gol.Error("用户名字读取失败:", err)
		return
	}

	c.Property.Set("name", string(name[:n]))

	global.OnlineUsersLock.Lock()
	defer global.OnlineUsersLock.Unlock()

	if _, ok := global.OnlineUsers[string(name[:n])]; ok {
		_, err = c.Conn.Write([]byte{0})
		if err != nil {
			gol.Errorln(e.ConnWriteErr(err))
			return
		}
	}

	_, err = c.Conn.Write([]byte{1})
	if err != nil {
		gol.Errorln(e.ConnWriteErr(err))
		return
	}

	global.OnlineUsers[string(name[:n])] = struct{}{}
}

func posthook(c *tcpx.Connection) {
	name, ok := c.Property.Get("name")
	if !ok {
		return
	}

	global.OnlineUsersLock.Lock()
	defer global.OnlineUsersLock.Unlock()

	delete(global.OnlineUsers, name.(string))
}