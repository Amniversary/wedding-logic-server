package mysql

import (
	"log"
	"time"

	"github.com/Amniversary/wedding-logic-server/config"
)

func NewSchedule(req *config.NewSchedule) bool {
	schedule := &Schedule{
		UserId:     req.UserId,
		Theme:      req.Theme,
		Phone:      req.Phone,
		Site:       req.Site,
		Time:       req.Time,
		Remind:     req.Remind,
		TimeFrame:  req.TimeFrame,
		HavePay:    req.HavePay,
		TotalPrice: req.TotalPrice,
		PayStatus:  req.PayStatus,
		Status:     1,
		CreateAt:  time.Now().Unix(),
	}
	tx := db.Begin()
	if err := tx.Create(&schedule).Error; err != nil {
		tx.Rollback()
		log.Printf("schedule create db err : [%v]", err)
		return false
	}
	if len(req.Cooperation) > 0 {
		for _, v := range req.Cooperation {
			err := tx.Create(&Cooperation{
				ScheduleId:   schedule.ID,
				Professional: v.Professional,
				Name:         v.Name,
				Phone:        v.Phone,
				CreateAt:    time.Now().Unix(),
			}).Error
			if err != nil {
				tx.Rollback()
				log.Printf("create cooperation err : [%v]", err)
				return false
			}
		}
	}
	tx.Commit()
	return true
}

func UpdateSchedule(req *config.UpSchedule) bool {
	schedule := &Schedule{
		ID:         req.ID,
		Theme:      req.Theme,
		Time:       req.Time,
		TimeFrame:  req.TimeFrame,
		Site:       req.Site,
		Remind:     req.Remind,
		HavePay:    req.HavePay,
		TotalPrice: req.TotalPrice,
		PayStatus:  req.PayStatus,
		Status:     req.Status,
		Phone:      req.Phone,
	}
	tx := db.Begin()
	if err := tx.Table("Schedule").Where("id = ?", req.ID).Updates(&schedule).Error; err != nil {
		tx.Rollback()
		log.Printf("update Schedule err: [%v]", err)
		return false
	}
	if err := tx.Where("schedule_id = ?", req.ID).Delete(&Cooperation{}).Error; err != nil {
		tx.Rollback()
		log.Printf("delete cooperation err: [%v]", err)
		return false
	}
	if len(req.Cooperation) > 0 {
		for _, v := range req.Cooperation {
			err := tx.Create(&Cooperation{
				ScheduleId:   schedule.ID,
				Professional: v.Professional,
				Name:         v.Name,
				Phone:        v.Phone,
				CreateAt:    time.Now().Unix(),
			}).Error
			if err != nil {
				tx.Rollback()
				log.Printf("create cooperation err : [%v]", err)
				return false
			}
		}
	}
	tx.Commit()
	return true
}


func GetUserScheduleList(req *config.GetUserScheduleList) ([]config.GetUserScheduleListRes, bool) {
	var list []config.GetUserScheduleListRes
	err := db.Table("Schedule").
		Select("id, theme, time_frame, create_at").
		Where("user_id = ? and `time` like ? and status = 1", req.UserId, req.Time + "%").Find(&list).Error
	if err != nil {
		log.Printf("getUserScheduleList err : [%v]", err)
		return nil, false
	}
	return list, true
}

func GetScheduleInfo(scheduleId int64) (*config.GetScheduleInfoRes, bool) {
	schedule := &config.GetScheduleInfoRes{}
	var newSchedule []config.NewCooperationInfo
	err := db.Table("Schedule").
		Select("id, theme, phone, site, `time`, time_frame, have_pay, total_price, pay_status, remind").
		Where("id = ? and status = 1", scheduleId).First(&schedule).Error
	if err != nil {
		log.Printf("getScheduleInfo query err : [%v]", err)
		return nil, false
	}
	err = db.Table("Cooperation").
		Select("id, professional, name, phone, create_at").
		Where("schedule_id = ?", scheduleId).Find(&newSchedule).Error
	if err != nil {
		log.Printf("select Cooperation query err: [%v]", err)
		return nil, false
	}
	schedule.Cooperation = newSchedule
	return schedule, true
}

func DelSchedule(scheduleId int64) bool {
	err := db.Table("Schedule").Where("id = ?", scheduleId).Update("status = 0").Error
	if err != nil {
		log.Printf("delSchedule query err : [%v]", err)
		return false
	}
	return true
}