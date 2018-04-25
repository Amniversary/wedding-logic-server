package config

//TODO: 短信模板 Code
const (
	SMS_CODE = "62560"
)

type GetCardInfo struct {
	CardId int64 `json:"cardId"`
	UserId int64 `json:"userId"`
}

type GetMyCard struct {
	UserId int64 `json:"userId"`
}

type GetCardInfoRes struct {
	ID           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	TeamId       int64  `json:"team_id"`
	Identity     int64  `json:"identity"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	Pic          string `json:"pic"`
	Qrcode       string `json:"qrcode"`
	BgPic        string `json:"bg_pic"`
	Professional string `json:"professional"`
	Company      string `json:"company"`
	Site         string `json:"site"`
	Explain      string `json:"explain"`
	Fame         int64  `json:"fame"`
	Like         int64  `json:"like"`
	Production   int64  `json:"production"`
	Schedule     int64  `json:"schedule"`
}

type ProductionClickLike struct {
	CardId       int64 `json:"cardId"`
	UserId       int64 `json:"userId"`
	Status       int64 `json:"status"`
	ProductionId int64 `json:"productionId"`
}

type DelProduction struct {
	CardId       int64 `json:"cardId"`
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
	ID           int64  `json:"id"`
	CardId       int64  `json:"card_id"`
	Picture      string `json:"picture"`
	Name         string `json:"name"`
	Professional string `json:"professional"`
	Content      string `json:"content"`
	Pic          string `json:"pic"`
	Like         int64  `json:"like"`
	IsClick      int64  `json:"is_click"`
	CreateAt     int64  `json:"create_at"`
}

type CheckValidateCode struct {
	UserId int64  `json:"userId"`
	Code   string `json:"code"`
}

type NewSchedule struct {
	UserId      int64            `json:"userId"`
	Theme       string           `json:"theme"`
	TimeFrame   string           `json:"time_frame"`
	Site        string           `json:"site"`
	Time        string           `json:"time"`
	Remind      string           `json:"remind"`
	HavePay     float64          `json:"have_pay"`
	TotalPrice  float64          `json:"total_price"`
	PayStatus   int64            `json:"pay_status"`
	Phone       string           `json:"phone"`
	Longitude   float64          `json:"longitude"`
	Latitude    float64          `json:"latitude"`
	Cooperation []NewCooperation `json:"cooperation"`
}

type UpSchedule struct {
	ID          int64            `json:"id"`
	Theme       string           `json:"theme"`
	TimeFrame   string           `json:"time_frame"`
	Site        string           `json:"site"`
	Time        string           `json:"time"`
	Remind      string           `json:"remind"`
	HavePay     float64          `json:"have_pay"`
	TotalPrice  float64          `json:"total_price"`
	PayStatus   int64            `json:"pay_status"`
	Status      int64            `json:"status"`
	Phone       string           `json:"phone"`
	Longitude   float64          `json:"longitude"`
	Latitude    float64          `json:"latitude"`
	Cooperation []NewCooperation `json:"cooperation"`
}

type NewCooperation struct {
	UserId       int64  `json:"user_id"`
	Professional string `json:"professional"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
}

type GetUserScheduleList struct {
	UserId int64  `json:"userId"`
	Time   string `json:"time"`
}

type GetUserScheduleListRes struct {
	ID        int64  `json:"id"`
	WeddingId int64  `json:"wedding_id"`
	Theme     string `json:"theme"`
	TimeFrame string `json:"time_frame"`
	Time      string `json:"time"`
}

type GetScheduleInfo struct {
	ScheduleId int64 `json:"scheduleId"`
}

type GetScheduleInfoRes struct {
	ID          int64                `json:"id"`
	WeddingId   int64                `json:"wedding_id"`
	Theme       string               `json:"theme"`
	TimeFrame   string               `json:"time_frame"`
	Site        string               `json:"site"`
	Time        string               `json:"time"`
	Remind      string               `json:"remind"`
	HavePay     float64              `json:"have_pay"`
	TotalPrice  float64              `json:"total_price"`
	Status      int64                `json:"status"`
	Phone       string               `json:"phone"`
	Longitude   float64              `json:"longitude"`
	Latitude    float64              `json:"latitude"`
	Cooperation []NewCooperationInfo `json:"cooperation"`
}

type NewCooperationInfo struct {
	ID           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	Professional string `json:"professional"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	CreateAt     int64  `json:"create_at"`
}

type NewTeam struct {
	UserId   int64  `json:"userId"`
	Pic      string `json:"pic"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Province string `json:"province"`
	City     string `json:"city"`
}

type GetTeamInfo struct {
	TeamID int64 `json:"teamId"`
}

type GetApplyInfo struct {
	TeamId int64 `json:"teamId"`
	UserId int64 `json:"userId"`
}

type ClickTeamProduction struct {
	ProductionId int64 `json:"productionId"`
	UserId       int64 `json:"userId"`
	Status       int64 `json:"status"`
}

type GetTeamProduction struct {
	TeamId   int64 `json:"teamId"`
	UserId   int64 `json:"userId"`
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type SearchTeam struct {
	Name string `json:"name"`
	Type int64  `json:"type"`
}

type SearchTeamList struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Pic      string `json:"pic"`
	Province string `json:"province"`
	City     string `json:"city"`
	CreateAt int64  `json:"create_at"`
}

type ApplyJoinList struct {
	ID       int64  `json:"id"`
	UserId   int64  `json:"user_id"`
	Name     string `json:"name"`
	CreateAt int64  `json:"create_at"`
}

type UpJoinStatus struct {
	ID     int64 `json:"id"`
	Status int64 `json:"status"`
}

type GetTeamList struct {
	ID           int64  `json:"id"`
	UserId       int64  `json:"user_id"`
	CardId       int64  `json:"card_id"`
	Identity     int64  `json:"identity"`
	Name         string `json:"name"`
	Pic          string `json:"pic"`
	Professional string `json:"professional"`
}

type GetMessageList struct {
	ID     int64  `json:"id"`
	TeamId int64  `json:"team_id"`
	Name   string `json:"name"`
	Status int64  `json:"status"`
	Type   int64  `json:"type"`
}

type DelTeamMember struct {
	ID int64 `json:"id"`
}

type GetQrcode struct {
	CardId int64 `json:"cardId"`
}

type RequestJson struct {
	ActionName string      `json:"action_name"`
	Data       interface{} `json:"data"`
}

type GetTeamScheduleList struct {
	TeamId int64  `json:"teamId"`
	Time   string `json:"time"`
}

type GetTeamScheduleRes struct {
	CardId    int64  `json:"card_id"`
	UserId    int64  `json:"user_id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	TimeFrame string `json:"time_frame"`
}

type GetTeamSchedule struct {
	Id        int64    `json:"id"`
	CardId    int64    `json:"card_id"`
	UserId    int64    `json:"user_id"`
	Name      string   `json:"name"`
	Pic       string   `json:"pic"`
	TimeFrame []string `json:"time_frame"`
}
type NewBusinessCard struct {
	UserId int64  `json:"userId"`
	PicId  int64  `json:"picId"`
	Cover  string `json:"cover"`
	Text   string `json:"text"`
}

type NewBusinessCardReq struct {
	PicId        int64  `json:"picId"`
	Cover        string `json:"cover"`
	Pic          string `json:"pic"`
	Qrcode       string `json:"qrcode"`
	Text         string `json:"text"`
	Professional string `json:"professional"`
	Name         string `json:"name"`
}

type GetBusinessBgList struct {
	PageNo   int64 `json:"pageNo"`
	PageSize int64 `json:"pageSize"`
}

type InvitationSchedule struct {
	ScheduleId int64 `json:"scheduleId"`
	UserId     int64 `json:"userId"`
}

type AuthWedding struct {
	ScheduleId int64 `json:"scheduleId"`
	UserId     int64 `json:"userId"`
	WeddingId  int64 `json:"weddingId"`
	Type       int64 `json:"type"`
}

type CardTeamMemberInfo struct {
	ID     int64  `json:"id"`
	UserId int64  `json:"user_id"`
	Name   string `json:"name"`
	TeamId int64  `json:"team_id"`
}

type DelTeamRequest struct {
	TeamId int64 `json:"teamId"`
	UserId int64 `json:"userId"`
}

type GetToken struct {
	OpenId string     `json:"openId"`
	Data   []FormData `json:"data"`
}

type FormData struct {
	FormId string `json:"formId"`
	Expire int64  `json:"expire"`
}

type TransferAdmin struct {
	TeamId  int64 `json:"teamId"`
	AdminId int64 `json:"adminId"`
	UserId  int64 `json:"userId"`
}
