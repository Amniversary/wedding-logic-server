package config

import "log"

const (
	RESPONSE_OK = 0
	RESPONSE_ERROR = 1
	ERROR_MSG = "系统错误"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type DBInfo struct {
	User   string
	Pass   string
	Host   string
	DBName string
}

type Config struct {
	Debug bool
	DBDebug bool

	Version string
	Port string
	DBInfo
}

func NewConfig() *Config {
	c := new(Config)
	c.DBInfo.Host = "sh-cdb-c7gk8cwq.sql.tencentcdb.com:63769"
	c.DBInfo.User = "root"
	c.DBInfo.Pass = "tkC42cwy2U3SQwHw"
	c.DBInfo.DBName = "wedding_card"

	c.Debug = false
	c.DBDebug = true
	c.Version = "1.0.1"
	c.Port = ":5609"

	log.Printf("Debug:[%v], DBDebug:[%v], Version:[%s], DBInfo:[User:[%s], Pass:[%s], Host:[%s], DBName:[%s]].",
		c.Debug, c.DBDebug, c.Version, c.User, c.Pass, c.Host, c.DBName, )

	return c
}
