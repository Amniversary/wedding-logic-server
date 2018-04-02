package mysql

import (
	"time"
	"log"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/Amniversary/wedding-logic-server/config"
)

const (
	CANCEL_LIKE = 0
	CLICK_LIKE  = 1
)

func CreateCard(card *Card) (error) {
	card.UpdatedAt = time.Now().Unix()
	card.CreatedAt = time.Now().Unix()
	if err := db.Create(&card).Error; err != nil {
		log.Printf("Create Card Model error: %v", err)
		return err
	}
	if card.ID == 0 {
		return fmt.Errorf("save card info res Id empty.")
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

func GetUserCardInfo(userId int64, cardId int64) (*Card, error) {
	card := &Card{}
	if err := db.Where("card_id = ? and user_id = ?",cardId ,userId).First(&card).Error; err != nil {
		log.Printf("select [MyCardInfo] err: %v", err)
		return card, err
	}
	if card.UserId != userId {
		if _, err := CreateCollect(cardId, userId); err != nil {
			return card, err
		}
	}
	return card, nil
}

func GetMyCardInfo(userId int64) (*Card, bool) {
	card := &Card{}
	if err := db.Where("user_id = ?", userId).First(&card).Error; err != nil {
		log.Printf("getMyCardInfo query err: [%v]", err)
		return nil, false
	}
	return card, true
}

func CreateCollect(cardId int64, userId int64) (Collection, error) {
	collect := Collection{CardId: cardId, UserId: userId, Status: 1}
	if err := db.Where("user_id = ? and card_id = ?", userId, cardId).First(&collect).Error; err != nil {
		log.Printf("select collect err : %v", err)
	}
	if collect.ID == 0 {
		collect.CreatedAt = time.Now().Unix()
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

func CreateProduction(production *Production) bool {
	production.CreatedAt = time.Now().Unix()
	tx := db.Begin()
	if err := tx.Create(&production).Error; err != nil {
		log.Printf("create dynamic model err : %v", err)
		tx.Rollback()
		return false
	}
	err := tx.Table("Card").Where("id = ?", production.CardId).Update("production", gorm.Expr("production + 1")).Error
	if err != nil {
		log.Printf("update card production err : [%v]", err)
		tx.Rollback()
		return false
	}
	return true
}

func ProductionClickLike(req *config.ProductionClickLike) bool {
	click := &ClickProduction{}
	if err := db.Where("user_id = ? and production_id = ?", req.UserId, req.ProductionId).First(&click).Error; err != nil {
		if click.ID == 0 {
			tx := db.Begin()
			clickProduction := &ClickProduction{UserId: req.UserId, ProductionId: req.ProductionId, Status: CLICK_LIKE}
			if err := tx.Create(&clickProduction).Error; err != nil {
				tx.Rollback()
				log.Printf("create clickProduction err: [%v] ", err)
				return false
			}
			err := tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("like + ?"), 1).Error
			if err != nil {
				tx.Rollback()
				log.Printf("uddate Production like err : [%v] ", err)
				return false
			}
			err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("like + ?"), 1).Error
			if err != nil {
				tx.Rollback()
				log.Printf("update Card like err : [%v]", err)
				return false
			}
			tx.Commit()
			return true
		}
	}
	if req.Status == 2 {
		req.Status = CANCEL_LIKE
	}
	if click.Status == req.Status {
		return true
	}
	tx := db.Begin()
	if err := tx.Model(&ClickProduction{}).Update("status", req.Status).Error; err != nil {
		log.Printf("update click Production err : [%v], [ProductionId: %d, status: %d]", err, req.ProductionId, req.Status)
		tx.Rollback()
		return false
	}
	var err error
	switch req.Status {
	case CANCEL_LIKE:
		if err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("like - ?", 1)).Error; err != nil {
			log.Printf("update card like err : [%v], [CardId: %d]", err, req.CardId)
			tx.Rollback()
			return false
		}
		err = tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("like - ?", 1)).Error
	case CLICK_LIKE:
		if err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("like + ?", 1)).Error; err != nil {
			log.Printf("update card like err : [%v], [CardId: %d]", err, req.CardId)
			tx.Rollback()
			return false
		}
		err = tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("lick + ?", 1)).Error
	}
	if err != nil {
		log.Printf("update card Production like err : [%v], [ProductionId: %d]", err, req.ProductionId)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DelProduction(productionId int64) bool {
	if err := db.Model(&Production{}).Where("id = ?", productionId).Update("status = 0").Error; err != nil {
		log.Printf("update Production status err : [%v]", err)
		return false
	}
	return true
}

func CreateSMS(req *config.ValidateCode, vCode string) (*SmsMessage, bool) {
	sms := &SmsMessage{UserId: req.UserId, Phone: req.Phone, Type: req.Type, Code: vCode}
	sms.CreatedAt = time.Now().Unix()
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

func GetProductionList(req *config.GetProductionList) ([]config.ProductionList, bool) {
	var list []config.ProductionList
	err := db.Table("Production pd").
		Select("pd.id, content, pic, like, create_at, ifnull(cp.status, 0) as is_click").
		Joins("left join ClickProduction cp on pd.id=cp.production_id and user_id = ?", req.UserId).
		Where("cd.card_id = ? and cd.status = 1", req.CardId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("cd.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetProductionList] err: %v", err)
		return nil, false
	}
	return list, true
}

func GetUserCode(userId int64) (SmsMessage, error) {
	sms := SmsMessage{}
	err := db.Where("user_id = ? and status = 1", userId).Select("id, user_id, code, created_at").Order("id desc").Limit(1).Find(&sms).Error
	if err != nil {
		log.Printf("select user code err : %v", err)
		return sms, err
	}
	return sms, err
}