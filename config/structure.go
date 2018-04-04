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
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	Pic       string `json:"pic"`
	Like      int64  `json:"like"`
	IsClick   int64  `json:"is_click"`
	CreatedAt int64  `json:"created_at"`
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
	Cooperation []NewCooperation `json:"cooperation"`
}

type NewCooperation struct {
	Professional string `json:"professional"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
}

type GetUserScheduleList struct {
	UserId int64 `json:"userId"`
}

type GetUserScheduleListRes struct {
	ID        int64  `json:"id"`
	Theme     string `json:"theme"`
	TimeFrame string `json:"time_frame"`
	CreatedAt int64  `json:"created_at"`
}

type GetScheduleInfo struct {
	ScheduleId int64 `json:"scheduleId"`
}

type GetScheduleInfoRes struct {
	ID          int64                `json:"id"`
	Theme       string               `json:"theme"`
	TimeFrame   string               `json:"time_frame"`
	Site        string               `json:"site"`
	Time        string               `json:"time"`
	Remind      string               `json:"remind"`
	HavePay     float64              `json:"have_pay"`
	TotalPrice  float64              `json:"total_price"`
	Status      int64                `json:"status"`
	Phone       string               `json:"phone"`
	Cooperation []NewCooperationInfo `json:"cooperation"`
}

type NewCooperationInfo struct {
	ID           int64  `json:"id"`
	Professional string `json:"professional"`
	Name         string `json:"name"`
	Phone        string `json:"phone"`
	CreatedAt    int64  `json:"created_at"`
}

type NewTeam struct {
	UserId int64  `json:"userId"`
	Pic    string `json:"pic"`
	Name   string `json:"name"`
}

type GetTeamInfo struct {
	TeamID int64 `json:"teamId"`
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
}

type SearchTeamList struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Pic       string `json:"pic"`
	CreatedAt int64  `json:"created_at"`
}
