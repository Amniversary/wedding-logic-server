package server

import (
	"net/http"
	"encoding/json"
	"log"

	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/Amniversary/wedding-logic-server/config/mysql"
)
/**
	TODO: 创建团队
 */
func (s *Server) NewTeam(w http.ResponseWriter,r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.NewTeam{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("newTeam json decode err : [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.UserId == 0 {
		Response.Msg = "用户Id不能为空"
		return
	}
	if req.Name == "" {
		Response.Msg = "团队名称不能为空"
		return
	}
	if ok := mysql.NewTeam(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
/**
	TODO: 获取团队信息
 */
func (s *Server) GetTeamInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetTeamInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getTeamInfo json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	info, err := mysql.GetTeamInfo(req.TeamID)
	if !err {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
	Response.Data = info
}
/**
	TODO: 更新团队信息
 */
func (s *Server) UpTeam(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &mysql.Team{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("upTeam json decode err : [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.ID == 0 {
		Response.Msg = "参数异常"
		log.Printf("[upTeam] %v", req)
		return
	}
	if ok := mysql.UpTeamInfo(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
