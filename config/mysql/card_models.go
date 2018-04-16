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

func CreateCard(card *Card) (int64, error) {
	card.UpdateAt = time.Now().Unix()
	card.CreateAt = time.Now().Unix()
	if err := db.Create(&card).Error; err != nil {
		log.Printf("Create Card Model error: %v", err)
		return 0, err
	}
	if card.ID == 0 {
		return 0, fmt.Errorf("save card info res Id empty.")
	}
	return card.ID, nil
}

func UpdateCardModel(card *Card) bool {
	card.UpdateAt = time.Now().Unix()
	if err := db.Table("Card").Where("id = ? and user_id = ?", card.ID, card.UserId).Update(&card).Error; err != nil {
		log.Printf("update card err : %v, [%v]", err, card)
		return false
	}
	return true
}

func GetUserCardInfo(userId int64, cardId int64) (*config.GetCardInfoRes, error) {
	card := &config.GetCardInfoRes{}
	if err := db.Table("Card c").
		Joins("left join TeamMembers tm on c.user_id = tm.user_id").
		Select("c.id, c.user_id, name, phone, pic, qrcode, bg_pic, professional, company, site, `explain`, fame, `like`, production, ifnull(tm.team_id, 0) as team_id, ifnull(tm.type, 0) as identity").
		Where("c.id = ?", cardId).Find(&card).Error; err != nil {
		log.Printf("select [MyCardInfo] err: %v", err)
		return nil, err
	}
	var count int64
	if err := db.Model(&Schedule{}).Where("user_id = ? and pay_status = 1", userId).Count(&count).Error; err != nil {
		log.Printf("getCountSchedule query err: [%v]", err)
		return nil, err
	}
	card.Schedule = count
	if card.UserId != userId {
		if _, err := CreateCollect(cardId, userId); err != nil {
			return nil, err
		}
	}
	return card, nil
}

func GetMyCardInfo(userId int64) (*config.GetCardInfoRes, bool) {
	card := &config.GetCardInfoRes{}
	if err := db.Table("Card c").
		Joins("left join TeamMembers tm on c.user_id = tm.user_id").
		Select("c.id, c.user_id, name, phone, pic, qrcode, bg_pic, professional, company, site, `explain`, fame, `like`, production, ifnull(tm.team_id, 0) as team_id, ifnull(tm.type, 0) as identity").
		Where("c.user_id = ?", userId).Find(&card).Error; err != nil {
		log.Printf("getMyCardInfo query err: [%v]", err)
		return nil, false
	}
	var count int64
	if err := db.Model(&Schedule{}).Where("user_id = ? and pay_status = 1", userId).Count(&count).Error; err != nil {
		log.Printf("getCountSchedule query err: [%v]", err)
		return nil, false
	}
	card.Schedule = count
	return card, true
}

func CreateCollect(cardId int64, userId int64) (Collection, error) {
	collect := Collection{CardId: cardId, UserId: userId, Status: 1}
	if err := db.Where("user_id = ? and card_id = ?", userId, cardId).First(&collect).Error; err != nil {
		log.Printf("select collect err : %v", err)
	}
	if collect.ID == 0 {
		collect.CreateAt = time.Now().Unix()
		if err := db.Create(&collect).Error; err != nil {
			log.Printf("create collection error: %v", err)
			return collect, err
		}
		if err := db.Model(&Card{}).Where("id = ?", cardId).Updates(map[string]interface{}{"fame": gorm.Expr("fame + 1"), "update_at": time.Now().Unix()}).Error; err != nil {
			log.Printf("update card fame error: %v ,  cardId:[%d]", err, cardId)
		}
	}
	return collect, nil
}

func CreateProduction(production *Production) bool {
	production.CreateAt = time.Now().Unix()
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
	tx.Commit()
	return true
}

func ProductionClickLike(req *config.ProductionClickLike) bool {
	click := &ClickProduction{}
	if err := db.Where("user_id = ? and production_id = ?", req.UserId, req.ProductionId).First(&click).Error; err != nil {
		if click.ID == 0 {
			tx := db.Begin()
			clickProduction := &ClickProduction{UserId: req.UserId, ProductionId: req.ProductionId, Status: CLICK_LIKE, CreateAt: time.Now().Unix()}
			if err := tx.Create(&clickProduction).Error; err != nil {
				tx.Rollback()
				log.Printf("create clickProduction err: [%v] ", err)
				return false
			}
			err := tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` + 1")).Error
			if err != nil {
				tx.Rollback()
				log.Printf("udpate Production like err : [%v] ", err)
				return false
			}
			err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("`like` + 1")).Error
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
	if err := tx.Model(&ClickProduction{}).Where("user_id = ? and production_id = ?", req.UserId, req.ProductionId).Update("status", req.Status).Error; err != nil {
		log.Printf("update click Production err : [%v], [ProductionId: %d, status: %d]", err, req.ProductionId, req.Status)
		tx.Rollback()
		return false
	}
	var err error
	switch req.Status {
	case CANCEL_LIKE:
		if err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("`like` - 1")).Error; err != nil {
			log.Printf("update card like err : [%v], [CardId: %d]", err, req.CardId)
			tx.Rollback()
			return false
		}
		err = tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` - 1")).Error
	case CLICK_LIKE:
		if err = tx.Model(&Card{}).Where("id = ?", req.CardId).Update("like", gorm.Expr("`like` + 1")).Error; err != nil {
			log.Printf("update card like err : [%v], [CardId: %d]", err, req.CardId)
			tx.Rollback()
			return false
		}
		err = tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` + 1")).Error
	}
	if err != nil {
		log.Printf("update card Production like err : [%v], [ProductionId: %d]", err, req.ProductionId)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DelProduction(req *config.DelProduction) bool {
	log.Printf("productionId: [%v], cardId: [%v]", req.ProductionId, req.CardId)
	tx := db.Begin()
	if err := tx.Model(&Production{}).Where("id = ?", req.ProductionId).Update("status", 0).Error; err != nil {
		log.Printf("update Production status err : [%v]", err)
		return false
	}
	if err := tx.Model(&Card{}).Where("id = ?", req.CardId).Update("production", gorm.Expr("production - 1")).Error; err != nil {
		log.Printf("update Card production err : [%v]", err)
		return false
	}
	tx.Commit()
	return true
}

func CreateSMS(req *config.ValidateCode, vCode string) (*SmsMessage, bool) {
	sms := &SmsMessage{UserId: req.UserId, Phone: req.Phone, Type: req.Type, Code: vCode}
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

func GetProductionList(req *config.GetProductionList) ([]config.ProductionList, bool) {
	var list []config.ProductionList
	err := db.Select("pd.id, content, pic, `like`, pd.create_at, ifnull(cp.status, 0) as is_click").
		Table("Production pd").
		Joins("left join ClickProduction cp on pd.id=cp.production_id and user_id = ?", req.UserId).
		Where("pd.card_id = ? and pd.status = 1", req.CardId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("pd.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetProductionList] err: %v", err)
		return nil, false
	}
	return list, true
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

func GetMessageList(userId int64) ([]config.GetMessageList, bool) {
	var list []config.GetMessageList
	err := db.Table("ApplyList al").
		Joins("inner join Team t on al.team_id = t.id").
		Select("al.id, team_id, name, status, type").
		Where("al.user_id = ? and status in (0,1)", userId).Find(&list).Error
	if err != nil {
		log.Printf("getMessageList query err: [%v]", err)
		return nil, false
	}
	return list, true
}

func GetCardInfo (userId int64) (*Card, bool) {
	card := &Card{}
	if err := db.Where("user_id = ?", userId).First(&card).Error; err != nil {
		log.Printf("getCardInfo query first err: [%v]", err)
		return nil, false
	}
	return card, true
}