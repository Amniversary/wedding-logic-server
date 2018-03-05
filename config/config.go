package config

const (
	ServerName = "CardLogic"
)

type Response struct {
	Code int64       `json:"code"`
	Msg  string      `json:"msg,omitempty"`
	Data interface{} `json:"data,omitempty"`
}

type GetCardInfo struct {
	CardId int64 `json:"cardId"`
	UserId int64 `json:"userId"`
}

type GetCardList struct {
	UserId   int64 `json:"userId"`
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

const (
	DBName = "wedding-card" 
	USER   = "root"         
	PASS   = "root"         
	HOST   = "127.0.0.1"    
	DEBUG  = "dev"          // dev or prod
)
