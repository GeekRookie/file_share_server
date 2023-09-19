package initialize

import (
	"encoding/json"
	"file_share_server/global"
	"os"
)

func initConfig() {
	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, global.Config)
	if err != nil {
		panic(err)
	}
}