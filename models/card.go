package models

import (
	"time"
	"log"
	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/jinzhu/gorm"
)

type Card struct {
	ID           uint64  `gorm:"primary_key" json:"id"`
	UserId       int64   `gorm:"not null;default 0;type:int;index" json:"userId"`
	Name         string  `gorm:"not null;default:'';type:varchar(64)" json:"name"`
	Phone        string  `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Pic          string  `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Professional string  `gorm:"not null;default:'';type:varchar(64)" json:"professional"`
	Year         string  `gorm:"not null;default:'';type:varchar(64)" json:"year"`
	Company      string  `gorm:"not null;default:'';type:varchar(128)" json:"company"`
	Site         string  `gorm:"not null;default:'';type:varchar(256)" json:"site"`
	Accessment   string  `gorm:"not null;type:text" json:"accessment"`
	Fame         int64   `gorm:"not null;default:0;type:int" json:"fame"`
	Lick         int64   `gorm:"not null;default:0;type:int" json:"lick"`
	Collect      int64   `gorm:"not null;default:0;type:int" json:"collect"`
	Longitude    float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"longitude"`
	Latitude     float64 `gorm:"not null;default:0;type:decimal(10,7)" json:"latitude"`
	CreateAt     int64   `gorm:"not null;default:0;type:int" json:"-"`
	UpdateAt     int64   `gorm:"not null;default:0;type:int" json:"-"`
}

type Collection struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	UserId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"userId"`
	CardId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"cardId"`
	Status   int64 `gorm:"not null;default:1;type:int" json:"status"`
	IsLick   int64 `gorm:"not null;default:0;type:int" json:"is_lick"`
	IsFame   int64 `gorm:"not null;default:1;type:int" json:"is_fame"`
	CreateAt int64 `gorm:"not null;default:0;type:int" json:"create_at"`
	UpdateAt int64 `gorm:"not null;default:0;type:int" json:"update_at"`
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
	return "cSmsMessage"
}

func (Card) TableName() string {
	return "cCard"
}

func (Collection) TableName() string {
	return "cCollection"
}

func CreateCardModel(card *Card) error {
	card.UpdateAt = time.Now().Unix()
	card.CreateAt = time.Now().Unix()
	if err := db.Create(&card).Error; err != nil {
		log.Printf("Create Card Model error: %v", err)
		return err
	}
	return nil
}

func UpdateCardModel(card *Card) bool {
	if err := db.Table("cCard").Where("id = ? and user_id = ?", card.ID, card.UserId).Update(&card).Error; err != nil {
		log.Printf("update card err : %v, [%v]", err, card)
		return false
	}
	return true
}

func GetCardData(cardId int64, userId int64) (config.UserCardInfo, error) {
	info, err := GetCardInfo(cardId, userId)
	if err != nil {
		return info, err
	}
	if info.UserId != userId {
		if _, err := CreateCollect(cardId, userId); err != nil {
			return info, err
		}
	}
	return info, nil
}

func GetCardList(req *config.GetCardList) ([]config.UserCardList, bool) {
	var list []config.UserCardList
	err := db.Table("cCollection cc").
		Select("cr.id, cr.user_id, name, pic, professional,year,fame, lick, is_lick").
		Joins("inner join cCard cr on cc.card_id = cr.id").
		Where("cc.user_id = ?", req.UserId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("cc.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetCardList] err: %v", err)
		return nil, false
	}
	return list, true
}

func GetUserCardInfo(userId int64) ([]Card, error) {
	var card []Card
	if err := db.Where("user_id = ?", userId).First(&card).Error; err != nil {
		log.Printf("select [MyCardInfo] err: %v", err)
		return nil, err
	}
	return card, nil
}

func GetCardInfo(cardId int64, userId int64) (config.UserCardInfo, error) {
	card := config.UserCardInfo{}
	if err := db.Table("cCard cc").
		Select("cc.id, cc.user_id, name, phone, pic, professional, year, company, site, accessment, fame, lick, ifnull(cn.is_lick,0) as is_lick").
		Joins("left join cCollection cn on cc.id=cn.card_id and cn.user_id = ?", userId).
		Where("cc.id = ?", cardId).Limit(1).
		Find(&card).Error; err != nil {
		log.Printf("select [GetCardInfo] err: %v", err)
		return card, err
	}
	return card, nil
}

func CreateCollect(cardId int64, userId int64) (Collection, error) {
	collect := Collection{CardId: cardId, UserId: userId, Status: 1}
	if err := db.Where("user_id = ? and card_id = ?", userId, cardId).First(&collect).Error; err != nil {
		collect.CreateAt = time.Now().Unix()
		collect.UpdateAt = time.Now().Unix()
		if err := db.Create(&collect).Error; err != nil {
			log.Printf("create collection error: %v", err)
			return collect, err
		}
		if err := db.Model(&Card{}).Update("fame", gorm.Expr("fame + 1")).Where("card_id = ?", cardId).Error; err != nil {
			log.Printf("update card fame error: %v ,  cardId:[%d]", err, cardId)
		}
	}
	return collect, nil
}

func SetClickLick(req *config.ClickLick) (bool, error) {
	collection := Collection{}
	if err := db.Where("user_id = ? and card_id = ?", req.UserId, req.CardId).First(&collection).Error; err != nil {
		return false, err
	}
	if req.Status == 2 {
		req.Status = 0
	}
	if collection.IsLick == req.Status {
		return true, nil
	}
	tx := db.Begin()
	err := tx.Model(&Collection{}).Where("user_id = ? and card_id = ?", req.UserId, req.CardId).Update("is_lick", req.Status).Error;
	if err != nil {
		tx.Rollback()
		return false, err
	}
	switch req.Status {
	case 0:
		err = db.Model(&Card{}).Where("id = ?", req.CardId).Update("lick", gorm.Expr("lick - ?", 1)).Error
	case 1:
		err = db.Model(&Card{}).Where("id = ?", req.CardId).Update("lick", gorm.Expr("lick + ?", 1)).Error
	}
	if err != nil {
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}

func CreateSMS(req *config.ValidateCode, vCode string) (SmsMessage, bool) {
	sms := SmsMessage{UserId: req.UserId, Phone: req.Phone, Type: req.Type, Code: vCode}
	sms.CreateAt = time.Now().Unix()
	if err := db.Create(&sms).Error; err != nil {
		log.Printf("create sms message err: %s", err)
		return sms, false
	}
	return sms, true
}

func UpdateSMS(netReturn map[string]interface{}, sms *SmsMessage) bool {
	net := netReturn["result"].(map[string]interface{})
	err := db.Model(&sms).Where("id = ?", sms.ID).Updates(map[string]interface{}{"count": net["count"], "fee": net["fee"], "sid": net["sid"], "text": netReturn["reason"], "status": 1}).Error
	if err != nil {
		log.Printf("update sms message err : %v", err)
		return false
	}
	return true
}

func GetUserCode(userId int64) (SmsMessage, error) {
	sms := SmsMessage{}
	err := db.Where("user_id = ? and status = 1", userId).Select("id, user_id, code, create_at").Order("id desc").Limit(1).Find(&sms).Error
	if err != nil {
		log.Printf("select user code err : %v", err)
		return sms, err
	}
	return sms, err
}
