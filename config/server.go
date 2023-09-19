package config

type ConfigModel struct {
	Server	   Server
	Broadcast  Broadcast
	Database   Database
	Log        Log
}

type Server struct {
	IP  string `json:IP`
	Port uint `json:Port`
}

type Broadcast struct {
	IP  string `json:IP`
	Port uint `json:Port`
}

type Database struct {
	DSN           string `json:"DSN"`
	SlowThreshold int    `json:"SlowThreshold"`
	LogLevel      int    `json:"LogLevel"`
	Colorful      bool   `json:"Colorful"`
}

type Log struct {
	Location string `json:"Location"`
}