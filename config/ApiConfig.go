package config

const (
	SET_CARD          = "setCard"         //TODO: 设置名片
	GET_VALIDATE_CODE = "getValidateCode" //TODO: 获取验证码
	GET_CARD_INFO     = "getCardInfo"     //TODO: 获取名片详情
	GET_CARD_LIST     = "getCardList"     //TODO: 获取名片列表
	NEW_DYNAMIC       = "newDynamic"      //TODO: 创建动态
	CLICK_LICK        = "clickLick"       //TODO: 名片点赞
	MY_CARD_INFO      = "myCardInfo"      //TODO: 我的名片
)

const (
	RESPONSE_OK    = 0
	RESPONSE_ERROR = 1
)

//TODO: 短信模板 Code
const (
	SMS_CODE = "62560"
)

const (
	ERROR_MSG = "系统错误"
)

type ClickLick struct {
	CardId int64 `json:"cardId"`
	UserId int64 `json:"userId"`
	Status int64 `json:"status"`
}

type ValidateCode struct {
	UserId int64  `json:"userId"`
	Phone  string `json:"phone"`
	Type   int64  `json:"type"`
}

type SmsCallBack struct {
	ErrorCode int64  `json:"error_code"`
	Reason    string `json:"reason"`
}
