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
	DBName = "wedding_card"     //wedding_card
	USER   = "root"             //root
	PASS   = "tkC42cwy2U3SQwHw" //tkC42cwy2U3SQwHw
	HOST   = "172.17.0.5"       //172.17.0.5
	DEBUG  = "prod"             //prod
)
