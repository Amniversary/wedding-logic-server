package models

import (
	"time"
	"log"
	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/jinzhu/gorm"
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

func CreateDynamic(req *CardDynamic) bool {
	req.CreateAt = time.Now().Unix()
	//pic, err := json.Marshal(req.Pic)
	//if err != nil {
	//	log.Printf("dynamic pic json encode err :%v", err)
	//	return false
	//}
	if err := db.Create(&req).Error; err != nil {
		log.Printf("create dynamic model err : %v", err)
		return false
	}
	return true
}

func DelDynamic(dynamic int64) bool {
	if err := db.Model(&CardDynamic{}).Where("id = ?", dynamic).Update("status", 0).Error; err != nil {
		log.Printf("delete card dynamic err: %v", err)
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

func SetDynamicClickLick(req *config.DynamicClick) (bool, error) {
	dynamic := ClickDynamic{}
	if err := db.Where("user_id = ? and dynamic_id = ?", req.UserId, req.DynamicId).First(&dynamic).Error; err != nil {
		tx := db.Begin()
		newDynamic := ClickDynamic{UserId: req.UserId, DynamicId: req.DynamicId, Status: 1}
		if err := db.Create(&newDynamic).Error; err != nil {
			log.Printf("create click dynamic err: %v", err)
			tx.Rollback()
			return false, err
		}
		err := db.Model(&CardDynamic{}).Where("id = ?", req.DynamicId).Update("lick", gorm.Expr("lick + ?", 1)).Error
		if err != nil {
			log.Printf("update card dynamic lick err : %v, [DynamicId: %d]", err, req.DynamicId)
			tx.Rollback()
			return false, err
		}
		err = db.Model(&Card{}).Where("id = ?", req.CardId).Update("lick", gorm.Expr("lick + ?", 1)).Error
		if err != nil {
			log.Printf("update card lick err: %v, [CardId:%d]", err, req.CardId)
			tx.Rollback()
			return false, err
		}
		tx.Commit()
		return true, nil
	}
	if req.Status == 2 {
		req.Status = 0
	}
	if dynamic.Status == req.Status {
		return true, nil
	}
	tx := db.Begin()
	if err := db.Model(&dynamic).Update("status", req.Status).Error; err != nil {
		log.Printf("update click dynamic err : %v, [dynamicId: %d, status:%d]", err, req.DynamicId, req.Status)
		tx.Rollback()
		return false, err
	}
	var err error
	switch req.Status {
	case 0:
		if err = db.Model(&Card{}).Where("id = ?", req.CardId).Update("lick", gorm.Expr("lick - ?", 1)).Error; err != nil {
			log.Printf("update card lick err : %v, [CardId:%d]", err, req.CardId)
			tx.Rollback()
			return false, err
		}
		err = db.Model(&CardDynamic{}).Where("id = ?", req.DynamicId).Update("lick", gorm.Expr("lick - ?", 1)).Error
	case 1:
		if err = db.Model(&Card{}).Where("id = ?", req.CardId).Update("lick", gorm.Expr("lick + ?", 1)).Error; err != nil {
			log.Printf("update card lick err : %v, [CardId:%d]", err, req.CardId)
			tx.Rollback()
			return false, err
		}
		err = db.Model(&CardDynamic{}).Where("id = ?", req.DynamicId).Update("lick", gorm.Expr("lick + ?", 1)).Error
	}
	if err != nil {
		log.Printf("update card dynamic lick err : %v, [dynamicId:%d]", err, req.DynamicId)
		tx.Rollback()
		return false, err
	}
	tx.Commit()
	return true, nil
}
