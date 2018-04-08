package server

import (
	"log"
	"net/http"
	"encoding/json"

	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/Amniversary/wedding-logic-server/config/mysql"
)

/**
	TODO: 获取作品列表
 */
func (s *Server) GetProductionList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetProductionList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getDynamicList json decode err : %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetProductionList(req)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 删除作品
 */
func (s *Server) DelProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.DelProduction{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("delProduction json decode err : [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.DelProduction(req.ProductionId); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 作品点赞
 */
func (s *Server) ClickLikeProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.ProductionClickLike{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("clickLikeProduction json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.ProductionClickLike(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 创建作品
 */
func (s *Server) NewProduction(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &mysql.Production{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("NewDynamic json decode err : %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.CardId == 0 {
		log.Printf("params can not empty: [%v]", req)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.CreateProduction(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}
