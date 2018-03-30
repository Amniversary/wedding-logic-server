package server

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/Amniversary/wedding-logic-server/config/mysql"
	"github.com/Amniversary/wedding-logic-server/config"
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
		return
	}
	if req.ID == 0 {
		Response.Msg = "scheduleId can not be empty."
		log.Printf("%s", Response.Msg)
		return
	}
	if ok := mysql.UpdateSchedule(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
