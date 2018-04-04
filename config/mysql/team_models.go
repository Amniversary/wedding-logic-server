package mysql

import (
	"log"
	"time"

	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/jinzhu/gorm"
)

func NewTeam(req *config.NewTeam) bool {
	team := &Team{UserId: req.UserId, Name: req.Name, Pic: req.Pic, CreatedAt: time.Now().Unix()}
	if err := db.Create(&team).Error; err != nil {
		log.Printf("create team model err : [%v]", err)
		return false
	}
	return true
}

func GetTeamInfo(teamId int64) (*Team, bool) {
	team := &Team{}
	if err := db.Where("id = ?", teamId).
		Select("`id`, `name`, `pic`, `cover`, `explain`, `created_at`").
		First(&team).Error; err != nil {
		log.Printf("getTeamInfo query err: [%v]", err)
		return nil, false
	}
	return team, true
}

func UpTeamInfo(req *Team) bool {
	req.CreatedAt = time.Now().Unix()
	if err := db.Table("Team").Where("id = ?", req.ID).Update(&req).Error; err != nil {
		log.Printf("upTeam query err : [%v]", err)
		return false
	}
	return true
}

func NewTeamProduction(req *TeamProduction) bool {
	req.CreatedAt = time.Now().Unix()
	req.Status = 1
	if err := db.Create(&req).Error; err != nil {
		log.Printf("newTeamProduction create err : [%v]", err)
		return false
	}
	return true
}

func DelTeamProduction(productionId int64) bool {
	if err := db.Model(&TeamProduction{}).Where("id = ?", productionId).Update("status = 0").Error; err != nil {
		log.Printf("delTeamProduction query err: [%v]", err)
		return false
	}
	return true
}

func ClickLikeTeamProduction(req *config.ClickTeamProduction) bool {
	production := &TeamClickProduction{}
	err := db.Where("user_id = ? and production_id = ?", req.UserId, req.ProductionId).First(&production).Error
	if err != nil {
		if production.ID == 0 {
			tx := db.Begin()
			clickProduction := &TeamClickProduction{UserId: req.UserId, ProductionId: req.ProductionId, Status: CLICK_LIKE, CreatedAt: time.Now().Unix()}
			if err := tx.Create(&clickProduction).Error; err != nil {
				log.Printf("create click production err : [%v]", err)
				tx.Rollback()
				return false
			}
			err := tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("like + 1"))
			if err != nil {
				log.Printf("update teamProduction like err: [%v]", err)
				tx.Rollback()
				return false
			}
			tx.Commit()
			return true
		}
	}
	if req.Status == 2 {
		req.Status = CANCEL_LIKE
	}
	if req.Status == production.Status {
		return true
	}
	tx := db.Begin()
	if err := tx.Model(&TeamClickProduction{}).Update("status", req.Status).Error; err != nil {
		log.Printf("update click TeamProduction err : [%v], [ProductionId: %d, status: %d]", err, req.ProductionId, req.Status)
		tx.Rollback()
		return false
	}
	switch req.Status {
	case CANCEL_LIKE:
		err = tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("like - ?", 1)).Error
	case CLICK_LIKE:
		err = tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("lick + ?", 1)).Error
	}
	if err != nil {
		log.Printf("update card TeamProduction like err : [%v], [ProductionId: %d]", err, req.ProductionId)
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func GetTeamProductionList(req *config.GetTeamProduction) ([]config.ProductionList, bool) {
	var list []config.ProductionList
	err := db.Table("TeamProduction tp").
		Select("tp.id, content, pic, like, created_at, ifnull(cp.status, 0) as is_click").
		Joins("left join TeamClickProduction cp on tp.id=cp.production_id and user_id = ?", req.UserId).
		Where("tp.team_id = ? and cd.status = 1", req.TeamId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("cd.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetProductionList] err: %v", err)
		return nil, false
	}
	return list, true
}

func SearchTeamModel(name string) ([]config.SearchTeamList, bool) {
	var list []config.SearchTeamList
	err := db.Table("Team").
		Select("id, name, pic, created_at").
		Where("name like ?", name + "%").Find(&list).Error
	if err != nil {
		log.Printf("searchTeamModel query err: [%v]", err)
		return nil, false
	}
	return list, true
}
