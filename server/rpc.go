package server

import (
	"log"
	"time"
	"fmt"
	"net/http"
	"github.com/Amniversary/wedding-logic-server/config"
	"encoding/json"
)

const (
	ServerName          = "FindWedding"
	SET_CARD            = "setCard"
	UP_CARD             = "upCard"
	GET_VALIDATE_CODE   = "getValidateCode"
	GET_CARD_INFO       = "getCardInfo"
	NEW_PRODUCTION      = "newProduction"
	CLICK_LIKE          = "clickLike"
	CHECK_VALIDATE_CODE = "checkValidateCode"
	GET_PRODUCTION_LIST = "getProductionList"
	CLICK_PRODUCTION    = "clickLikeProduction"
	DEL_PRODUCTION      = "delProduction"
	NEW_SCHEDULE 		= "newSchedule"
	UP_SCHEDULE			= "upSchedule"
)

func (s *Server) rpc(w http.ResponseWriter, r *http.Request) {
	res := &config.Response{Code: config.RESPONSE_OK}
	if r.Method != "POST" {
		log.Printf("Method not be Post Request [%s]\n", r.Method)
		EchoJson(w, http.StatusOK, res)
		return
	}
	serverName := r.Header.Get("ServerName")
	if serverName != ServerName {
		log.Printf("ServerName: [%s]  request -> ServerName: [%s] Method: [%s]\n", ServerName, serverName, r.Method)
		EchoJson(w, http.StatusOK, res)
		return
	}

	methodName := r.Header.Get("MethodName")
	start := time.Now()
	defer func() {
		log.Printf("Request MethodName: [%s], Rtime[%v]\n", methodName, time.Now().Sub(start))
	}()
	switch methodName {
	case SET_CARD:
		s.AddCard(w, r)
	case CLICK_PRODUCTION:
		s.ClickLikeProduction(w, r)
	case UP_CARD:
		s.UpCard(w, r)
	case GET_VALIDATE_CODE:
		s.GetValidateCode(w, r)
	case GET_CARD_INFO:
		s.GetCardInfo(w, r)
	case CHECK_VALIDATE_CODE:
		s.CheckValidateCode(w, r)
	case NEW_PRODUCTION:
		s.NewProduction(w, r)
	case CLICK_LIKE:
		s.ClickLikeProduction(w, r)
	case GET_PRODUCTION_LIST:
		s.GetProductionList(w, r)
	case DEL_PRODUCTION:
		s.DelProduction(w, r)
	case NEW_SCHEDULE:
		s.NewSchedule(w, r)
	case UP_SCHEDULE:
		s.UpSchedule(w, r)

	default:
		res.Code = 1
		res.Msg = fmt.Sprintf("Can't find the interface: [%s]", methodName)
		EchoJson(w, http.StatusOK, res)
	}
}

// TODO @ 输出Json数据
func EchoJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "x-requested-with,content-type,servername,methodname,userid,msgid")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
