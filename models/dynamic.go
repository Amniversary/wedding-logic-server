package models

import (
	"time"
	"log"
	"encoding/json"
	"github.com/Amniversary/wedding-logic-server/config"
)

type CardDynamic struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	CardId   int64  `gorm:"not null;default:0;type:int;index" json:"cardId"`
	Content  string `gorm:"not null;type:text" json:"content"`
	Pic      string `gorm:"not null;type:text" json:"pic"`
	Lick     int64  `gorm:"not null;default:0;type:int" json:"lick"`
	Status   int64  `gorm:"not null;default:1;type:int;index" json:"-"`
	CreateAt int64  `gorm:"not null;default:0;type:int" json:"create_at"`
}

type ClickDynamic struct {
	ID        uint64 `gorm:"primary_key" json:"id"`
	UserId    int64  `gorm:"not null;default:0;type:int;index" json:"userId"`
	DynamicId int64  `gorm:"not null;default:0;type:int;index" json:"dynamicId"`
	Status    int64  `gorm:"not null;default:1;type:int" json:"status"`
}

func (CardDynamic) TableName() string {
	return "cCardDynamic"
}

func (ClickDynamic) TableName() string {
	return "cClickDynamic"
}

func CreateDynamic(req *config.NewDynamic) bool {
	dynamic := CardDynamic{CardId:req.CardId, Content:req.Content}
	dynamic.CreateAt = time.Now().Unix()
	pic, err := json.Marshal(req.Pic)
	if err != nil {
		log.Printf("dynamic pic json encode err :%v", err)
		return false
	}
	dynamic.Pic = string(pic)
	if string(pic) == "null" {
		dynamic.Pic = ""
	}
	if err := db.Create(&dynamic).Error; err != nil {
		log.Printf("create dynamic model err : %v", err)
		return false
	}
	return true
}

func GetDynamicList(req *config.GetDynamicList) ([]config.DynamicList, bool) {
	var list []config.DynamicList
	err := db.Table("cCardDynamic cd").
		Select("cd.id, content, pic, lick, create_at, ifnull(cc.status, 0) as is_click").
		Joins("left join cClickDynamic cc on cd.id=cc.dynamic_id and user_id = ?", req.UserId).
		Where("cd.card_id = ? and cd.status = 1", req.CardId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("cd.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetDynamicList] err: %v", err)
		return nil, false
	}
	return list, true
}