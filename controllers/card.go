package controllers

import (
	"net/http"
	"github.com/Amniversary/wedding-logic-server/config"
	"encoding/json"
	"log"
	"github.com/Amniversary/wedding-logic-server/models"
	"math/rand"
	"fmt"
	"time"
	"io/ioutil"
	"net/url"
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
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardInfo json decode err: %v\n", err)
		return
	}
	card, err := models.GetCardData(req.CardId, req.UserId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = card
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取名片列表
func GetCardList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardList json decode err: %v\n", err)
		return
	}
	cardList, ok := models.GetCardList(req)
	if !ok {
		return
	}
	Response.Data = cardList
	Response.Code = config.RESPONSE_OK
}

//TODO: 名片点赞
func ClickLick(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
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

//TODO: 获取短信验证码
func GetValidateCode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.ValidateCode{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("ValidateCode json decode err: %v", err)
		return
	}
	rands := rand.New(rand.NewSource(time.Now().UnixNano()))
	num := fmt.Sprintf("%04d", rands.Int63n(10000))
	vCode := "#code#=" + num
	sms, res := models.CreateSMS(req, num)
	if !res {
		Response.Msg = config.ERROR_MSG
		return
	}
	netReturn, ok := SendJuHeSMS(req.Phone, config.SMS_CODE, vCode)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	if rst := models.UpdateSMS(netReturn, &sms); !rst {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取我的名片信息
func MyCardInfo(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetCardInfo{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getCardInfo json decode err: %v\n", err)
		return
	}
	card, err := models.GetUserCardInfo(req.UserId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = card
	Response.Code = config.RESPONSE_OK
}

//TODO: 创建动态
func NewDynamic(w http.ResponseWriter, r *http.Request) {

}

//TODO: 校验验证码
func CheckValidateCode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code:config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.CheckValidateCode{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("checkValidateCode json decode err : %v", err)
		return
	}
	res, err := models.GetUserCode(req.UserId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	t := time.Now().Unix()
	if (t - res.CreateAt) > 600 {
		Response.Msg = "验证码已过期, 请重新获取 !"
		return
	}
	if res.Code != req.Code {
		Response.Msg = "验证码错误 !"
		return
	}
	Response.Code = config.RESPONSE_OK
}

func SendJuHeSMS(phone string, tpId string, vCode string) (map[string]interface{}, bool) {
	key := "6962e47932431e9608350c1d5bfb523c"
	juheURL := "http://v.juhe.cn/sms/send"
	param := url.Values{}
	param.Set("mobile", phone)    //接收短信的手机号码
	param.Set("tpl_id", tpId)     //短信模板ID，请参考个人中心短信模板设置
	param.Set("tpl_value", vCode) //变量名和变量值对。如果你的变量名或者变量值中带有#&amp;=中的任意一个特殊符号，请先分别进行urlencode编码后再传递，&lt;a href=&quot;http://www.juhe.cn/news/index/id/50&quot; target=&quot;_blank&quot;&gt;详细说明&gt;&lt;/a&gt;
	param.Set("key", key)         //应用APPKEY(应用详细页查询)
	param.Set("dtype", "json")    //返回数据的格式,xml或json，默认json

	data, err := Get(juheURL, param)
	if err != nil {
		log.Printf("getJuhe Request err : %v", err)
		return nil, false
	}
	var netReturn map[string]interface{}
	json.Unmarshal(data, &netReturn)
	if netReturn["error_code"].(float64) == 0 {
		//log.Printf("接口返回result字段是:\r\n%v", netReturn)
		return netReturn, true
	}
	return nil, false
}



func Get(apiURL string, params url.Values) (rs []byte, err error) {
	var Url *url.URL
	Url, err = url.Parse(apiURL)
	if err != nil {
		log.Printf("parse url err: %v", err)
		return nil, err
	}
	//如果参数中有中文参数,这个方法会进行URLEncode
	Url.RawQuery = params.Encode()
	resp, err := http.Get(Url.String())
	if err != nil {
		log.Printf("get request err: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
