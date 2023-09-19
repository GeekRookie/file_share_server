package initialize

import (
	"file_share_server/global"
	"fmt"
	"net"
)

func initBroadcast() {
	var err error
	global.Broadcaster, err = net.Dial("udp", fmt.Sprintf("%s:%d", global.Config.Broadcast.IP, global.Config.Broadcast.Port))
	if err != nil {
		panic(err)
	}
}
