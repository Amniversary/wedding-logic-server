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
	Schedule     int64   `gorm:"not null;default:0;type:int" json:"schedule"`
	Longitude    float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"longitude"`
	Latitude     float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"latitude"`
	CreateAt     int64   `gorm:"not null;default:0;type:int" json:"-"`
	UpdateAt     int64   `gorm:"not null;default:0;type:int" json:"-"`
}

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

func (Collection) TableName() string {
	return "Collection"
}

type Production struct {
	ID       int64  `gorm:"primary_key" json:"id"`
	CardId   int64  `gorm:"not null;default:0;type:int;index" json:"card_id"`
	Content  string `gorm:"not null;type:text" json:"content"`
	Pic      string `gorm:"not null;type:text" json:"pic"`
	Like     int64  `gorm:"not null;default:0;type:int" json:"like"`
	Status   int64  `gorm:"not null;default:1;type:int;index" json:"-"`
	CreateAt int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

func (Production) TableName() string {
	return "Production"
}

type ClickProduction struct {
	ID           uint64 `gorm:"primary_key" json:"id"`
	UserId       int64  `gorm:"not null;default:0;type:int;index" json:"userId"`
	ProductionId int64  `gorm:"not null;default:0;type:int;index" json:"productionId"`
	Status       int64  `gorm:"not null;default:1;type:int" json:"status"`
}

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

func (SmsMessage) TableName() string {
	return "SmsMessage"
}
