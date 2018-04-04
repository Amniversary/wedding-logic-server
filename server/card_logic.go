package server

import (
	"log"
	"time"
	"fmt"
	"net/http"
	"encoding/json"
	"math/rand"

	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/Amniversary/wedding-logic-server/config/mysql"
	"github.com/Amniversary/wedding-logic-server/components"
)

/**
	TODO: 创建婚礼人信息
 */
func (s *Server) AddCard(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	card := &mysql.Card{}
	if err := json.NewDecoder(r.Body).Decode(card); err != nil {
		log.Printf("setCard json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if err := mysql.CreateCard(card); err != nil {
		log.Printf("create card err: [%v]", err)
		return
	}
	//ok, err := SendGenCardQrcode(cardId)
	//if !ok {
	//	log.Printf("sendGenCard request err: %v", err)
	//	Response.Msg = config.ERROR_MSG
	//	return
	//}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 更新婚礼人信息
 */
func (s *Server) UpCard(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	card := &mysql.Card{}
	if err := json.NewDecoder(r.Body).Decode(card); err != nil {
		log.Printf("upCard json decode err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if ok := mysql.UpdateCardModel(card); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取名片详情
 */
func (s *Server) GetCardInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardInfo json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	card, err := mysql.GetUserCardInfo(req.UserId, req.CardId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = card
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取我的名片详情
 */
func (s *Server) GetMyCardInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetMyCard{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardInfo json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	info, ok := mysql.GetMyCardInfo(req.UserId)
	if !ok {
		Response.Data = ""
		Response.Code = config.RESPONSE_OK
		return
	}
	Response.Data = info
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取短信验证码
 */
func (s *Server) GetValidateCode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.ValidateCode{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("ValidateCode json decode err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := fmt.Sprintf("%04d", rands.Int63n(9999))
	vCode := "#code#=" + num
	sms, res := mysql.CreateSMS(req, num)
	if !res {
		Response.Msg = config.ERROR_MSG
		return
	}
	netReturn, ok := components.SendJuHeSMS(req.Phone, config.SMS_CODE, vCode)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	if rst := mysql.UpdateSMS(netReturn, sms); !rst {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 校验验证码
 */
func (s *Server) CheckValidateCode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.CheckValidateCode{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("checkValidateCode json decode err : %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	res, err := mysql.GetUserCode(req.UserId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	t := time.Now().Unix()
	time := t - res.CreatedAt
	if time > 600 {
		Response.Msg = "验证码已过期, 请重新获取 !"
		log.Printf("%v : %v - %v = [%v]", Response.Msg, t, res.CreatedAt, time)
		return
	}
	if res.Code != req.Code {
		Response.Msg = "验证码错误 !"
		return
	}
	Response.Code = config.RESPONSE_OK
}
