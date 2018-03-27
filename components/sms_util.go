package components

import (
	"log"
	"net/url"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

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