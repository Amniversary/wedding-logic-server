package controllers

import (
	"net/http"
	"github.com/Amniversary/wedding-logic-server/config"
	"encoding/json"
	"log"
	"github.com/Amniversary/wedding-logic-server/models"
)

//TODO: 创建名片
func SetCard(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	card := &models.Card{}
	err := json.NewDecoder(r.Body).Decode(card)
	if err != nil {
		log.Printf("setCard json decode err: %v", err)
		return
	}
	if err := models.CreateCardModel(card); err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取名片信息
func GetCardInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardInfo json decode err: %v\n", err)
		return
	}
	card, err := models.GetCardData(req.CardId, req.UserId)
	if  err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = card
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取名片列表
func GetCardList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardList json decode err: %v\n", err)
		return
	}

}

//TODO: 名片点赞
func ClickLick(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.ClickLick{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("ClickLick json decode err: %v", err)
		return
	}
	if ok, err := models.SetClickLick(req); !ok {
		log.Printf("setClickLick error: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}