package main

import "file_share_server/initialize"

func main() {
	engine := initialize.Init()
	err := engine.Run()
	if err != nil {
		panic(err)
	}
}
