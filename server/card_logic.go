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
	"strconv"
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
	if card.UserId == 0 {
		Response.Msg = "params can not be empty."
		log.Printf("%v: [%v]", Response.Msg, card)
		return
	}
	cardId, err := mysql.CreateCard(card)
	if err != nil {
		log.Printf("create card err: [%v]", err)
		return
	}
	ok, err := components.SendGenCardQrcode(cardId)
	if !ok {
		log.Printf("sendGenCard request err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
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
	if req.CardId == 0 || req.UserId == 0 {
		Response.Msg = config.ERROR_MSG
		log.Printf("params can not be empty: [%v]", req)
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
	if req.UserId == 0 {
		Response.Msg = config.ERROR_MSG
		log.Printf("cardId can not be empty: [%v]", req)
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
	if req.Type == 0 {
		log.Printf("params [type] can not be empty : [%v]", req)
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
	time := t - res.CreateAt
	if time > 600 {
		Response.Msg = "验证码已过期, 请重新获取 !"
		log.Printf("%v : %v - %v = [%v]", Response.Msg, t, res.CreateAt, time)
		return
	}
	if res.Code != req.Code {
		Response.Msg = "验证码错误 !"
		return
	}
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 消息列表
 */
func (s *Server) GetMessageList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetMyCard{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getMessageList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, ok := mysql.GetMessageList(req.UserId)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 创建个性名片图
 */
func (s *Server) NewBusinessCard(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.NewBusinessCard{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("newBusinessCard json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if req.UserId == 0 || req.PicId == 0 {
		Response.Msg = "params can not be empty."
		log.Printf("%v: [%v]", Response.Msg, req)
		return
	}
	info, Ok := mysql.GetCardInfo(req.UserId)
	if !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	data := &config.NewBusinessCardReq{
		PicId:        req.PicId,
		Cover:        req.Cover,
		Text:         req.Text,
		Pic:          info.Pic,
		Qrcode:       info.Qrcode,
		Professional: info.Professional,
		Name:         info.Name,
	}
	res, Ok := components.GetNewBusinessCard(data)
	if !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	log.Printf("%v", res)
	Response.Data = res.Msg
	Response.Code = config.RESPONSE_OK
}

/**
	TODO: 获取个性名片文字列表
 */
func (s *Server) GetBusinessText(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	var list = []string{
		"将来的你，一定会感谢现在拼命的自己",
		"你可以迷茫，但请你不要虚度",
		"这一路走来，说不上多辛苦，庆幸自己很清楚",
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}
/**
	TODO: 获取名片背景图列表
 */
func (s *Server) GetBusinessBgList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetBusinessBgList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getBusinessBgList json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	list, Ok := mysql.GetBusinessBgList(req)
	if !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}
/**
	TODO: 收集用户推送码
 */
func (s *Server) GetToken(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetToken{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getToken json decode err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if len(req.Data) <= Empty {
		Response.Msg = "params can not be empty."
		log.Printf("%v", Response.Msg)
		return
	}
	headerUserId := r.Header.Get("userId")
	if headerUserId == "" {
		Response.Msg = "header userId params can not be empty."
		log.Printf("%v", Response.Msg)
		return
	}
	userId, err := strconv.ParseInt(headerUserId, 10, 64)
	if err != nil {
		log.Printf("header userId parseInt err: [%v]", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	if userId == Empty {
		Response.Msg = "header userId cannot be empty."
		log.Printf("%v: [%v]", Response.Msg, userId)
		return
	}
	if Ok := mysql.SaveToken(userId, req.Data); !Ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}