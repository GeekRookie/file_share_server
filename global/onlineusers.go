package global

import "sync"

var OnlineUsers = make(map[string]struct{}, 0)

var OnlineUsersLock = new(sync.RWMutex)