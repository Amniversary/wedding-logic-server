package mysql

type Card struct {
	ID           int64  `gorm:"primary_key" json:"id"`
	UserId       int64  `gorm:"not null;default 0;type:int;index" json:"userId"`
	Name         string `gorm:"not null;default:'';type:varchar(64)" json:"name"`
	Phone        string `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Pic          string `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Qrcode       string `gorm:"not null;default:'';type:varchar(256)" json:"qrcode"`
	BgPic        string `gorm:"not null;default:'';type:varchar(256)" json:"bg_pic"`
	Professional string `gorm:"not null;default:'';type:varchar(64)" json:"professional"`
	Company      string `gorm:"not null;default:'';type:varchar(128)" json:"company"`
	Site         string `gorm:"not null;default:'';type:varchar(256)" json:"site"`
	Explain      string `gorm:"not null;type:text" json:"explain"`
	Fame         int64  `gorm:"not null;default:0;type:int" json:"fame"`
	Like         int64  `gorm:"not null;default:0;type:int" json:"like"`
	Production   int64  `gorm:"not null;default:0;type:int" json:"production"`
	CreateAt     int64  `gorm:"not null;default:0;type:int" json:"create_at"`
	UpdateAt     int64  `gorm:"not null;default:0;type:int" json:"update_at"`
}

/**
	TODO: 名片详情表
 */
func (Card) TableName() string {
	return "Card"
}

type Collection struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	UserId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"userId"`
	CardId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"cardId"`
	Status   int64 `gorm:"not null;default:1;type:int" json:"status"`
	CreateAt int64 `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 用户浏览热度表
 */
func (Collection) TableName() string {
	return "Collection"
}

type Production struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	CardId   int64  `gorm:"not null;default:0;type:int;index" json:"cardId"`
	Content  string `gorm:"not null;type:text" json:"content"`
	Pic      string `gorm:"not null;type:text" json:"pic"`
	Like     int64  `gorm:"not null;default:0;type:int" json:"like"`
	Status   int64  `gorm:"not null;default:1;type:int;index" json:"-"`
	CreateAt int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 作品表
 */
func (Production) TableName() string {
	return "Production"
}

type ClickProduction struct {
	ID           uint64 `gorm:"primary_key" json:"id"`
	UserId       int64  `gorm:"not null;default:0;type:int;index" json:"userId"`
	ProductionId int64  `gorm:"not null;default:0;type:int;index" json:"productionId"`
	Status       int64  `gorm:"not null;default:1;type:int" json:"status"`
	CreateAt     int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 作品点赞表
 */
func (ClickProduction) TableName() string {
	return "ClickProduction"
}

type SmsMessage struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	UserId   int64  `gorm:"not null;default:0;type:int;index" json:"user_id"`
	Phone    string `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Count    int64  `gorm:"not null;default:0;type:int" json:"count"`
	Fee      int64  `gorm:"not null;default:0;type:int" json:"fee"`
	Type     int64  `gorm:"not null;default:0;type:int" json:"type"`
	Code     string `gorm:"not null;default:'';type:varchar(64)" json:"code"`
	Sid      string `gorm:"not null;default:'';type:varchar(256)" json:"sid"`
	Text     string `gorm:"not null;default:'';type:varchar(256)" json:"text"`
	Status   int64  `gorm:"not null;default:0;type:int;index" json:"status"`
	CreateAt int64  `gorm:"not null;default:0;type:int;" json:"create_at"`
}

/**
	TODO: 短信表
 */
func (SmsMessage) TableName() string {
	return "SmsMessage"
}

type Schedule struct {
	ID         int64   `gorm:"primary_key" json:"id"`
	UserId     int64   `gorm:"not null;default:0;type:int" json:"user_id"`
	Theme      string  `gorm:"not null;default:'';type:varchar(128)" json:"theme"`
	Phone      string  `gorm:"not null;default:'';type:varchar(128)" json:"phone"`
	Site       string  `gorm:"not null;default:'';type:varchar(128)" json:"site"`
	Time       string  `gorm:"not null;default:'';type:varchar(128)" json:"time"`
	Remind     string  `gorm:"not null;default:'';type:varchar(128)" json:"remind"`
	TimeFrame  string  `gorm:"not null;default:'';type:varchar(128)" json:"time_frame"`
	HavePay    float64 `gorm:"not null;default:0;type:decimal(12,2)" json:"have_pay"`
	TotalPrice float64 `gorm:"not null;default:0;type:decimal(12,2)" json:"total_money"`
	PayStatus  int64   `gorm:"not null;default:0;type:int" json:"pay_status"`
	Status     int64   `gorm:"not null;default:1;type:int;index" json:"status"`
	CreateAt   int64   `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 档期表
 */
func (Schedule) TableName() string {
	return "Schedule"
}

type Cooperation struct {
	ID           int64  `gorm:"primary_key" json:"id"`
	ScheduleId   int64  `gorm:"not null;default:0;type:int;index" json:"schedule_id"`
	Professional string `gorm:"not null;default:'';type:varchar(128)" json:"professional"`
	Name         string `gorm:"not null;default:'';type:varchar(128)" json:"name"`
	Phone        string `gorm:"not null;default:'';type:varchar(128)" json:"phone"`
	CreateAt     int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 合作人关联表
 */
func (Cooperation) TableName() string {
	return "Cooperation"
}

type Team struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	UserId   int64  `gorm:"not null;default:0;type:int;index" json:"user_id"`
	Pic      string `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Name     string `gorm:"not null;default:'';type:varchar(128);index" json:"name"`
	Cover    string `gorm:"not null;default:'';type:varchar(512)" json:"cover"`
	Explain  string `gorm:"not null;default:'';type:varchar(512)" json:"explain"`
	CreateAt int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 团队信息表
 */
func (Team) TableName() string {
	return "Team"
}

type TeamProduction struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	TeamId   int64  `gorm:"not null;default:0;type:int;index" json:"teamId"`
	Content  string `gorm:"not null;default:'';type:text" json:"content"`
	Pic      string `gorm:"not null;default:'';type:text" json:"pic"`
	Like     int64  `gorm:"not null;default:0;type:int" json:"like"`
	Status   int64  `gorm:"not null;default:0;type:int" json:"status"`
	CreateAt int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 团队作品表
 */
func (TeamProduction) TableName() string {
	return "TeamProduction"
}

type TeamClickProduction struct {
	ID           int64 `gorm:"primary_key" json:"id"`
	UserId       int64 `gorm:"not null;default:0;type:int;index" json:"userId"`
	ProductionId int64 `gorm:"not null;default:0;type:int;index" json:"productionId"`
	Status       int64 `gorm:"not null;default:1;type:int" json:"status"`
	CreateAt     int64 `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 团队作品点赞表
 */
func (TeamClickProduction) TableName() string {
	return "TeamClickProduction"
}

type ApplyList struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	TeamId   int64 `gorm:"not null;default:0;type:int;index" json:"team_id"`
	UserId   int64 `gorm:"not null;default:0;type:int;index" json:"user_id"`
	Type     int64 `gorm:"not null;default:0;type:int" json:"type"`
	Status   int64 `gorm:"not null;default:2;type:int" json:"status"`
	CreateAt int64 `gorm:"not null;default:0;type:int" json:"create_at"`
}

/**
	TODO: 团队申请列表
 */
func (ApplyList) TableName() string {
	return "ApplyList"
}

type TeamMembers struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	TeamId   int64 `gorm:"not null;default:0;type:int;index" json:"team_id"`
	UserId   int64 `gorm:"not null;default:0;type:int;index" json:"user_id"`
	CreateAt int64 `gorm:"not null;default:0;type:int" json:"create_at"`
}

func (TeamMembers) TableName() string {
	return "TeamMembers"
}
