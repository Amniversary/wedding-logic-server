package config

const (
	SET_CARD            = "setCard"           //TODO: 设置名片
	GET_VALIDATE_CODE   = "getValidateCode"   //TODO: 获取验证码
	GET_CARD_INFO       = "getCardInfo"       //TODO: 获取名片详情
	GET_CARD_LIST       = "getCardList"       //TODO: 获取名片列表
	NEW_DYNAMIC         = "newDynamic"        //TODO: 创建动态
	CLICK_LICK          = "clickLick"         //TODO: 名片点赞
	MY_CARD_INFO        = "myCardInfo"        //TODO: 我的名片
	CHECK_VALIDATE_CODE = "checkValidateCode" //TODO: 验证 手机验证码
	GET_DYNAMIC_LIST    = "getDynamicList"    //TODO: 获取动态列表
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

type CheckValidateCode struct {
	UserId int64  `json:"userId"`
	Code   string `json:"code"`
}

type UserCardList struct {
	Id           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Professional string `json:"professional"`
	Year         string `json:"year"`
	Fame         int64  `json:"fame"`
	Lick         int64  `json:"lick"`
	IsClick      int64  `json:"is_click"`
}

type NewDynamic struct {
	CardId  int64    `json:"cardId"`
	Content string   `json:"content"`
	Pic     []string `json:"pic"`
}

type GetDynamicList struct {
	CardId   int64 `json:"cardId"`
	UserId   int64 `json:"userId"`
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type DynamicList struct {
	ID       int64  `json:"id"`
	Content  string `json:"content"`
	Pic      string `json:"pic"`
	Lick     int64  `json:"lick"`
	IsClick  int64  `json:"is_click"`
	CreateAt int64  `json:"create_at"`
}
