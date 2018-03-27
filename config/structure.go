package config

//TODO: 短信模板 Code
const (
	SMS_CODE = "62560"
)

type GetCardInfo struct {
	CardId int64 `json:"cardId"`
	UserId int64 `json:"userId"`
}

type ProductionClickLike struct {
	CardId       int64 `json:"cardId"`
	UserId       int64 `json:"userId"`
	Status       int64 `json:"status"`
	ProductionId int64 `json:"productionId"`
}

type DelProduction struct {
	ProductionId int64 `json:"productionId"`
}

type ValidateCode struct {
	UserId int64  `json:"userId"`
	Phone  string `json:"phone"`
	Type   int64  `json:"type"`
}

type GetProductionList struct {
	CardId   int64 `json:"cardId"`
	UserId   int64 `json:"userId"`
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type ProductionList struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	Pic      string `json:"pic"`
	Like     int64  `json:"like"`
	IsClick  int64  `json:"is_click"`
	CreateAt int64  `json:"create_at"`
}

type CheckValidateCode struct {
	UserId int64  `json:"userId"`
	Code   string `json:"code"`
}