package server

import (
	"time"
	"net/http"
	"log"
	"fmt"

	"github.com/Amniversary/wedding-logic-server/config"
)

func (s *Server) collect(w http.ResponseWriter, r *http.Request) {
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
	methodExc, ok := s.collectMap[methodName]
	if !ok {
		res.Code = config.RESPONSE_ERROR
		res.Msg = fmt.Sprintf("Can't find the interface: [%s]", methodName)
		EchoJson(w, http.StatusOK, res)
		return
	}
	methodExc(w, r)
}

func (s *Server) initCollectMap() {
	var MethodMap = map[string]MethodFunc{
		"getToken": s.GetToken,
	}
	s.collectMap = MethodMap
}
