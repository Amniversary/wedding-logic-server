package server

import (
	"github.com/Amniversary/wedding-logic-server/config"
	"github.com/Amniversary/wedding-logic-server/config/mysql"
	"net/http"
	"log"
)

type ServerBase interface {
	Run()
}

type Server struct {
	cfg *config.Config

	methodMap map[string]MethodFunc
	collectMap map[string]MethodFunc
}

func NewServer(cfg *config.Config) ServerBase {
	return &Server{cfg: cfg}
}

func (s *Server) init() {
	mysql.NewMysql(s.cfg)
	s.initMap()
	s.initCollectMap()
}

func (s *Server) runServer() {
	http.HandleFunc("/rpc", s.rpc)
	http.HandleFunc("/token", s.collect)
	log.Printf("ListenServer Port: [%s]", s.cfg.Port)
	http.ListenAndServe(s.cfg.Port, nil)
}

func (s *Server) Run() {
	s.init()
	s.runServer()
}