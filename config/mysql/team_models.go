package mysql

import (
	"log"
	"time"

	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/jinzhu/gorm"
)

const (
	JoinSuccess = 1
	SearchCity  = 1
	SearchTeam  = 2
)

func NewTeam(req *config.NewTeam) bool {
	team := &Team{
		UserId:   req.UserId,
		Name:     req.Name,
		Pic:      req.Pic,
		Province: req.Province,
		City:     req.City,
		CreateAt: time.Now().Unix(),
	}
	tx := db.Begin()
	if err := tx.Create(&team).Error; err != nil {
		log.Printf("create team model err : [%v]", err)
		return false
	}
	teamMember := &TeamMembers{TeamId: team.ID, UserId: req.UserId, CreateAt: time.Now().Unix(), Type: 1}
	if err := tx.Create(&teamMember).Error; err != nil {
		log.Printf("create team members err : [%v]", err)
		return false
	}
	tx.Commit()
	return true
}

func GetTeamInfo(teamId int64) (*Team, bool) {
	team := &Team{}
	if err := db.Where("id = ?", teamId).First(&team).Error; err != nil {
		log.Printf("getTeamInfo query err: [%v]", err)
		return nil, false
	}
	return team, true
}

func UpTeamInfo(req *Team) bool {
	//req.CreateAt = time.Now().Unix()
	if err := db.Table("Team").Where("id = ?", req.ID).Update(&req).Error; err != nil {
		log.Printf("upTeam query err : [%v]", err)
		return false
	}
	return true
}

func NewTeamProduction(req *TeamProduction) bool {
	req.CreateAt = time.Now().Unix()
	req.Status = 1
	if err := db.Create(&req).Error; err != nil {
		log.Printf("newTeamProduction create err : [%v]", err)
		return false
	}
	return true
}

func DelTeamProduction(productionId int64) bool {
	if err := db.Model(&TeamProduction{}).Where("id = ?", productionId).Update("status", 0).Error; err != nil {
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
			clickProduction := &TeamClickProduction{UserId: req.UserId, ProductionId: req.ProductionId, Status: CLICK_LIKE, CreateAt: time.Now().Unix()}
			if err := tx.Create(&clickProduction).Error; err != nil {
				log.Printf("create click production err : [%v]", err)
				tx.Rollback()
				return false
			}
			err := tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` + 1")).Error
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
	if err := tx.Model(&TeamClickProduction{}).Where("user_id = ? and production_id = ?", req.UserId, req.ProductionId).Update("status", req.Status).Error; err != nil {
		log.Printf("update click TeamProduction err : [%v], [ProductionId: %d, status: %d]", err, req.ProductionId, req.Status)
		tx.Rollback()
		return false
	}
	switch req.Status {
	case CANCEL_LIKE:
		err = tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` - 1")).Error
	case CLICK_LIKE:
		err = tx.Model(&TeamProduction{}).Where("id = ?", req.ProductionId).Update("like", gorm.Expr("`like` + 1")).Error
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
		Select("tp.id, content, pic, `like`, tp.create_at, ifnull(cp.status, 0) as is_click").
		Joins("left join TeamClickProduction cp on tp.id=cp.production_id and user_id = ?", req.UserId).
		Where("tp.team_id = ? and tp.status = 1", req.TeamId).
		Offset((req.PageNo - 1) * req.PageSize).
		Limit(req.PageSize).
		Order("tp.create_at desc").Find(&list).Error
	if err != nil {
		log.Printf("select [GetProductionList] err: %v", err)
		return nil, false
	}
	return list, true
}

func SearchTeamModel(req *config.SearchTeam) ([]config.SearchTeamList, bool) {
	var list []config.SearchTeamList
	var err error
	switch req.Type {
	case SearchCity:
		err = db.Table("Team").
			Select("id, name, pic, city, province, create_at").
			Where("city = ?", req.Name).Find(&list).Error
	case SearchTeam:
		err = db.Table("Team").
			Select("id, name, pic, city, province,create_at").
			Where("name like ?", req.Name+"%").Find(&list).Error
	}
	if err != nil {
		log.Printf("searchTeamModel query err: [%v]", err)
		return nil, false
	}
	return list, true
}

func ApplyJoin(userId int64, teamId int64) (int64) {
	members := &TeamMembers{}
	if err := db.Where("user_id = ?", userId).First(&members).Error; err == nil {
		if members.ID != 0 {
			return 1 //fmt.Errorf("已加入团队, 无法申请")
		}
	}
	apply := &ApplyList{}
	if err := db.Where("team_id = ? and user_id = ? and status = 2", teamId, userId).First(&apply).Error; err != nil {
		if apply.ID == 0 {
			applyInfo := &ApplyList{TeamId: teamId, UserId: userId, Status: 2, CreateAt: time.Now().Unix(), Type: 1}
			if err := db.Create(&applyInfo).Error; err != nil {
				log.Printf("create applyJoinList err: [%v]", err)
				return 2
			}
		}
	}
	return 0
}

func GetApplyJoinList(teamId int64) ([]config.ApplyJoinList, bool) {
	var list []config.ApplyJoinList
	err := db.Table("ApplyList al").
		Joins("inner join Card c on al.user_id = c.user_id").
		Select("al.id, al.user_id, c.name, al.create_at").
		Where("al.team_id = ? and status = 2", teamId).Find(&list).Error
	if err != nil {
		log.Printf("getApplyJoinList query err: [%v]", err)
		return nil, false
	}
	return list, true
}

func UpdateJoinStatus(req *config.UpJoinStatus) (bool) {
	apply := &ApplyList{}
	if err := db.Where("id = ?", req.ID).First(&apply).Error; err != nil {
		log.Printf("updateJoinStatus select query err: [%v]", err)
		return false
	}
	if apply.Status == req.Status {
		return true
	}
	tx := db.Begin()
	err := tx.Where("id = ?", req.ID).Table("ApplyList").Update("status", req.Status).Error
	if err != nil {
		log.Printf("updateJoinStatus up query err: [%v]", err)
		return false
	}
	if req.Status == JoinSuccess {
		teamMember := &TeamMembers{TeamId: apply.TeamId, UserId: apply.UserId, Type: 2, CreateAt: time.Now().Unix()}
		if err := tx.Create(&teamMember).Error; err != nil {
			log.Printf("create teamMembers err : [%v]", err)
			return false
		}
	}
	tx.Commit()
	return true
}

func InvitationJoinTeam(req *config.GetApplyInfo) bool {
	teamMember := &TeamMembers{}
	if err := db.Where("team_id = ? and user_id = ?", req.TeamId, req.UserId).First(&teamMember).Error; err != nil {
		log.Printf("invitation Query err: [%v]", err)
	}
	if teamMember.ID == 0 {
		member := &TeamMembers{TeamId: req.TeamId, UserId: req.UserId, Type: 2, CreateAt: time.Now().Unix()}
		if err := db.Create(&member).Error; err != nil {
			log.Printf("create join Team err: [%v]", err)
			return false
		}
	}
	return true
}

func GetTeamList(teamId int64) ([]config.GetTeamList, bool) {
	var list []config.GetTeamList
	err := db.Select("ap.id, c.id as card_id, ap.user_id, name, pic ,professional").
		Table("ApplyList ap").
		Joins("inner join Card c on ap.user_id = c.user_id").
		Where("team_id = ? and status = 1", teamId).Find(&list).Error
	if err != nil {
		log.Printf("getTeamList query err : [%v]", err)
		return nil, false
	}
	log.Printf("%v", list)
	return list, true
}

func DelTeamMember(id int64) (bool) {
	if err := db.Where("id = ?", id).Delete(&TeamMembers{}).Error; err != nil {
		log.Printf("delTeamMember query err: [%v]", err)
		return false
	}
	return true
}

func GetTeamScheduleList(req *config.GetTeamScheduleList) ([]config.GetTeamScheduleRes, bool) {
	var list []config.GetTeamScheduleRes
	err := db.Select("s.id, c.user_id, c.name, c.pic, s.time_frame").
		Table("TeamMembers tm").
		Joins("left join `Schedule` s on tm.user_id = s.user_id").
		Joins("left join Card c on tm.user_id = c.user_id").
		Where("team_id = ? and `time` = ? and time_frame = ? and pay_status = 1", req.TeamId, req.Time, req.TimeFrame).Find(&list).Error
	if err != nil {
		log.Printf("getTeamScheduleList query err: [%v]", err)
		return nil, false
	}
	return list, true

}
