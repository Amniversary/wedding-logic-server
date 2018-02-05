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
	"bytes"
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
	cardId, err := models.CreateCardModel(card)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	ok, err := SendGenCardQrcode(cardId)
	if !ok {
		log.Printf("sendGenCard request err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

//TODO: 更新名片
func UpCard(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	card := &models.Card{}
	if err := json.NewDecoder(r.Body).Decode(card); err != nil {
		log.Printf("upCard json decode err: %v", err)
		return
	}
	if ok := models.UpdateCardModel(card); !ok {
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

func DelDynamic(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.DelDynamic{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("DelDynamic json decode err: %v", err)
		return
	}
	if ok := models.DelDynamic(req.DynamicId); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取微信分享二维码
func GetQrcode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetQrcode{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("GetQrcode json decode err: %v", err)
		return
	}
	card, err := models.GenCardQrcode(req.CardId)
	if err != nil {
		Response.Msg = config.ERROR_MSG
		return
	}
	Url, err := SendGenQrcode(&card)
	if err != nil {
		log.Printf("send genQrcode request err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	rsp := &config.RepCode{}
	rsp.Url = Url
	Response.Data = rsp
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
	num := fmt.Sprintf("%04d", rands.Int63n(9999))
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
		Response.Code = config.RESPONSE_OK
		Response.Data = card
		return
	}
	Response.Data = card
	Response.Code = config.RESPONSE_OK
}

//TODO: 获取动态列表
func GetDynamicList(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.GetDynamicList{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("getDynamicList json decode err : %v", err)
		return
	}
	list, ok := models.GetDynamicList(req)
	if !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Data = list
	Response.Code = config.RESPONSE_OK
}

//TODO: 创建动态
func NewDynamic(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &models.CardDynamic{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("NewDynamic json decode err : %v", err)
		return
	}
	if ok := models.CreateDynamic(req); !ok {
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

//TODO: 校验验证码
func CheckValidateCode(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
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

//TODO: 动态点赞
func ClickLickDynamic(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := &config.DynamicClick{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("ClickLickDynamic json docode err: %v", err)
		return
	}
	if ok, err := models.SetDynamicClickLick(req); !ok {
		log.Printf("setClickLick error: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	Response.Code = config.RESPONSE_OK
}

func GetSystemParams(w http.ResponseWriter, r *http.Request) {
	Response := &config.Response{Code: config.RESPONSE_ERROR}
	defer func() {
		EchoJson(w, http.StatusOK, Response)
	}()
	req := config.SystemParams{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		log.Printf("GetSystemParams json decode err: %v", err)
		Response.Msg = config.ERROR_MSG
		return
	}
	req.CreateQrCode = 1
	Response.Data = req
	Response.Code = config.RESPONSE_OK
}

func SendGenCardQrcode(cardId int64) (bool, error) {
	Url := "http://172.17.16.11:5607/api/response.do"
	client := http.Client{}
	data := &config.GetQrcode{CardId: cardId}
	request := &config.GenWeddingCardReq{
		ActionName: "save_qrcode",
		Data:       data,
	}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		log.Printf("GenCardQrcode json encode err: %v", err)
		return false, err
	}
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Printf("http new request err: %v", err)
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http do request err : %v", err)
		return false, err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil realAll err: %v", err)
		return false, err
	}
	response := &config.Response{}
	if err := json.Unmarshal(rspBody, response); err != nil {
		log.Printf("json decode err: %v", err)
		return false, err
	}
	if response.Code != config.RESPONSE_OK {
		log.Printf("wedding card server is err: [%v], response-Code: [%d], errMsg:[%s]", request, response.Code, response.Msg)
		return false, fmt.Errorf("wedding card service genCard error.")
	}
	return true, nil
}

//TODO: 发送聚合短信
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

func Post(apiUrl string, params url.Values) (rs []byte, err error) {
	resp, err := http.PostForm(apiUrl, params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func SendGenQrcode(card *models.Card) (string, error) {
	Url := "http://172.17.16.11:5607/api/response.do"
	client := http.Client{}
	data := card
	request := &config.GenWeddingCardReq{
		ActionName: "gen_card_qrcode",
		Data:       data,
	}
	reqBytes, err := json.Marshal(request)
	if err != nil {
		log.Printf("SendCardQrcode json encode err: %v", err)
		return "", err
	}
	req, err := http.NewRequest("POST", Url, bytes.NewBuffer(reqBytes))
	if err != nil {
		log.Printf("http new request err: %v", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("http do request err : %v", err)
		return "", err
	}
	defer resp.Body.Close()
	rspBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("ioutil realAll err: %v", err)
		return "", err
	}
	response := &config.Response{}
	if err := json.Unmarshal(rspBody, response); err != nil {
		log.Printf("json decode err: %v", err)
		return "", err
	}
	if response.Code != config.RESPONSE_OK {
		log.Printf("wedding card server is err: [%v], response-Code: [%d], errMsg:[%s]", request, response.Code, response.Msg)
		return "", fmt.Errorf("wedding card service genCard error.")
	}
	return response.Msg, nil
}