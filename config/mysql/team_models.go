package mysql

import (
	"github.com/Amniversary/wedding-logic-server/config"
	"log"
	"time"
)

func NewTeam(req *config.NewTeam)  bool {
	team := &Team{UserId: req.UserId, Name: req.Name, Pic: req.Pic, CreatedAt:time.Now().Unix()}
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
