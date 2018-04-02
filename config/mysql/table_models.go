package mysql

type Card struct {
	ID           int64   `gorm:"primary_key" json:"id"`
	UserId       int64   `gorm:"not null;default 0;type:int;index" json:"userId"`
	Name         string  `gorm:"not null;default:'';type:varchar(64)" json:"name"`
	Phone        string  `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Pic          string  `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Qrcode       string  `gorm:"not null;default:'';type:varchar(256)" json:"qrcode"`
	BgPic        string  `gorm:"not null;default:'';type:varchar(256)" json:"bg_pic"`
	Professional string  `gorm:"not null;default:'';type:varchar(64)" json:"professional"`
	Company      string  `gorm:"not null;default:'';type:varchar(128)" json:"company"`
	Site         string  `gorm:"not null;default:'';type:varchar(256)" json:"site"`
	Explain      string  `gorm:"not null;type:text" json:"explain"`
	Fame         int64   `gorm:"not null;default:0;type:int" json:"fame"`
	Like         int64   `gorm:"not null;default:0;type:int" json:"like"`
	Production   int64   `gorm:"not null;default:0;type:int" json:"production"`
	Longitude    float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"longitude"`
	Latitude     float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"latitude"`
	CreatedAt    int64   `gorm:"not null;default:0;type:int" json:"-"`
	UpdatedAt    int64   `gorm:"not null;default:0;type:int" json:"-"`
}

/**
	TODO: 名片详情表
 */
func (Card) TableName() string {
	return "Card"
}

type Collection struct {
	ID        int64 `gorm:"primary_key" json:"id"`
	UserId    int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"userId"`
	CardId    int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"cardId"`
	Status    int64 `gorm:"not null;default:1;type:int" json:"status"`
	CreatedAt int64 `gorm:"not null;default:0;type:int" json:"created_at"`
}

/**
	TODO: 用户浏览热度表
 */
func (Collection) TableName() string {
	return "Collection"
}

type Production struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	CardId    int64  `gorm:"not null;default:0;type:int;index" json:"card_id"`
	Content   string `gorm:"not null;type:text" json:"content"`
	Pic       string `gorm:"not null;type:text" json:"pic"`
	Like      int64  `gorm:"not null;default:0;type:int" json:"like"`
	Status    int64  `gorm:"not null;default:1;type:int;index" json:"-"`
	CreatedAt int64  `gorm:"not null;default:0;type:int" json:"created_at"`
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
	CreatedAt    int64  `gorm:"not null;default:0;type:int" json:"created_at"`
}

/**
	TODO: 作品点赞表
 */
func (ClickProduction) TableName() string {
	return "ClickProduction"
}

type SmsMessage struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	UserId    int64  `gorm:"not null;default:0;type:int;index" json:"user_id"`
	Phone     string `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Count     int64  `gorm:"not null;default:0;type:int" json:"count"`
	Fee       int64  `gorm:"not null;default:0;type:int" json:"fee"`
	Type      int64  `gorm:"not null;default:0;type:int" json:"type"`
	Code      string `gorm:"not null;default:'';type:varchar(64)" json:"code"`
	Sid       string `gorm:"not null;default:'';type:varchar(256)" json:"sid"`
	Text      string `gorm:"not null;default:'';type:varchar(256)" json:"text"`
	Status    int64  `gorm:"not null;default:0;type:int;index" json:"status"`
	CreatedAt int64  `gorm:"not null;default:0;type:int;" json:"created_at"`
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
	CreatedAt  int64   `gorm:"not null;default:0;type:int" json:"created_at"`
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
	CreatedAt    int64  `gorm:"not null;default:0;type:int" json:"created_at"`
}

/**
	TODO: 合作人关联表
 */
func (Cooperation) TableName() string {
	return "Cooperation"
}

type Team struct {
	ID        int64  `gorm:"primary_key" json:"id"`
	UserId    int64  `gorm:"not null;default:0;type:int;index" json:"user_id"`
	Pic       string `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Name      string `gorm:"not null;default:'';type:varchar(128)" json:"name"`
	Cover     string `gorm:"not null;default:'';type:varchar(512)" json:"cover"`
	Explain   string `gorm:"not null;default:'';type:varchar(512)" json:"explain"`
	CreatedAt int64  `gorm:"not null;default:0;type:int" json:"created_at"`
}

/**
	TODO: 团队信息表
 */
func (Team) TableName() string {
	return "Team"
}
