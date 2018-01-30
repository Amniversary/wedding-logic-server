package models

import (
	"time"
	"log"
)

type Card struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	UserId       int64  `gorm:"not null;default 0;type int" json:"userId"`
	Name         string `gorm:"not null;default:'';type:varchar(64)" json:"name"`
	Phone        string `gorm:"not null;default:'';type:varchar(64)" json:"phone"`
	Pic          string `gorm:"not null;default:'';type:varchar(256)" json:"pic"`
	Professional string `gorm:"not null;default:'';type:varchar(64)" json:"professional"`
	Year         string `gorm:"not null;default:'';type:varchar(64)" json:"year"`
	Code         string `gorm:"not null;default:'';type:varchar(64)" json:"code"`
	Company      string `gorm:"not null;default:'';type:varchar(128)" json:"company"`
	Site         string `gorm:"not null;default:'';type:varchar(256)" json:"site"`
	Accessment   string `gorm:"not null;type:text" json:"accessment"`
	Fame         int64  `gorm:"not null;default:0;type:int" json:"fame"`
	Lick         int64  `gorm:"not null;default:0;type:int" json:"lick"`
	Collect      int64  `gorm:"not null;default:0;type:int" json:"collect"`
	CreateAt     int64  `gorm:"not null;default:0;type:int" json:"create_at"`
	UpdateAt     int64  `gorm:"not null;default:0;type:int" json:"update_at"`
}

type Collection struct {
	ID       int64 `gorm:"primary_key" json:"id"`
	UserId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"userId"`
	CardId   int64 `gorm:"not null;default:0;type:int;unique_index:UserId_CardId" json:"cardId"`
	Status   int64 `gorm:"not null;default:1;type:int" json:"status"`
	CreateAt int64 `gorm:"not null;default:0;type:int" json:"create_at"`
	UpdateAt int64 `gorm:"not null;default:0;type:int" json:"update_at"`
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

func GetCardData(cardId int64, userId int64) (Card, error) {
	info, err := GetCardInfo(cardId)
	if err != nil {
		return info, err
	}
	if _, err := CreateCollect(cardId, userId); err != nil {
		return info, err
	}
	return info, nil
}

func GetCardList() {

}

func GetCardInfo(cardId int64) (Card, error) {
	card := Card{}
	if err := db.Where("id = ?", cardId).First(&card).Error; err != nil {
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
	}
	return collect, nil
}
