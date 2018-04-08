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
func (s *Server) NewTeam(w http.ResponseWriter, r *http.Request) {
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

/**
	TODO: 上传团队作品
 */
func (s *Server) NewTeamProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &mysql.TeamProduction{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("newTeamProduction json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.TeamId == 0 {
		Response.Msg = "teamId can not be empty."
		log.Printf("%v : [%v]", Response.Msg, req)
		return
	}
	if ok := mysql.NewTeamProduction(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 删除团队作品
 */
func (s *Server) DelTeamProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.DelProduction{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("delTeamProduction json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if Ok := mysql.DelTeamProduction(req.ProductionId); !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取团队作品列表
 */
func (s *Server) GetTeamProductionList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetTeamProduction{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getTeamProductionList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetTeamProductionList(req)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 团队作品点赞
 */
func (s *Server) ClickLikeTeamProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.ClickTeamProduction{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("clickLikeTeamProduction json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.ClickLikeTeamProduction(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 搜索团队
 */
func (s *Server) SearchTeam(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.SearchTeam{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("searchTeam json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.SearchTeamModel(req.Name)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 申请加入团队
 */
func (s *Server) ApplyJoinTeam(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetApplyInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("applyJoinTeam json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.ApplyJoin(req.UserId, req.TeamId); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取申请列表
 */
func (s *Server) ApplyJoinList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetTeamInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("applyJoinList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetApplyJoinList(req.TeamID)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 修改申请状态
 */
func (s *Server) UpJoinStatus(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.UpJoinStatus{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("upJoinStatus json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.UpdateJoinStatus(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取成员列表
 */
func (s *Server) GetTeamList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetTeamInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getTeamList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.TeamID == 0 {
		log.Printf("params can not be empty: [%v]", req)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetTeamList(req.TeamID)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}
