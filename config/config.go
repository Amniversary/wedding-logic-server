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
	UserId int64 `json:"userId"`
}

const (
	DBName = "wedding_card"
	USER   = "hebihan"
	PASS   = "tkC42cwy2U3SQwHw"
	HOST   = "172.17.0.5"
)
