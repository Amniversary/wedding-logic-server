package server

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/Amniversary/wedding-logic-server/config/mysql"
	"github.com/Amniversary/wedding-logic-server/config"
)

const (
	Authorize       = 1
	CancelAuthorize = 2
)

/**
	TODO: 创建档期
 */
func (s *Server) NewSchedule(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.NewSchedule{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("setSchedule json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	log.Printf("req: [%v]", req)
	if ok := mysql.NewSchedule(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 修改档期信息
 */
func (s *Server) UpSchedule(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.UpSchedule{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("upSchedule json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.ID == 0 {
		Response.Msg = "scheduleId can not be empty."
		log.Printf("%s", Response.Msg)
		return
	}
	log.Printf("%v", req)
	if ok := mysql.UpdateSchedule(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取用户档期列表
 */
func (s *Server) GetUserScheduleList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetUserScheduleList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getScheduleList json decode err : [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetUserScheduleList(req)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
	Response.Data = list
}

/**
	TODO: 获取档期详情
 */
func (s *Server) GetScheduleInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetScheduleInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getScheduleInfo json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}

	info, err := mysql.GetScheduleInfo(req.ScheduleId)
	if !err {
		Response.Msg = config.ERROR_MSG
		return
	}
	mysql.GetScheduleInfo(req.ScheduleId)
	Response.Code = config.RESPONSE_OK
	Response.Data = info
}

/**
	TODO: 删除档期
 */
func (s *Server) DelSchedule(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetScheduleInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("delSchedule json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.DelSchedule(req.ScheduleId); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 邀请档期成员
 */
func (s *Server) InvitationSchedule(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.InvitationSchedule{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("invitationSchedule json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if Ok := mysql.InvitationSchedule(req); !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 授权婚礼
 */
func (s *Server) AuthorizeWedding(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.AuthWedding{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("authorizeWedding json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.UserId == 0 || req.ScheduleId == 0 || req.WeddingId == 0 || req.Type == 0 {
		Response.Msg = "params can not be empty."
		log.Printf("%v: [%v]", Response.Msg, Response)
		return
	}
	switch req.Type {
	case Authorize:
		if Ok, err := mysql.AuthWedding(req); !Ok {
			if err != nil {
				Response.Msg = "已授权其他婚礼, 授权失败"
				log.Printf("%v", err)
				return 
			}
			Response.Msg = config.ERROR_MSG
			return
		}
	case CancelAuthorize:
		if Ok := mysql.CancelAuthWedding(req); !Ok {
			Response.Msg = config.ERROR_MSG
			return
		}
	}
	Response.Code = config.RESPONSE_OK
}

