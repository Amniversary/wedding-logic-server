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
}

func NewServer(cfg *config.Config) ServerBase {
	return &Server{cfg: cfg}
}

func (s *Server) init() {
	mysql.NewMysql(s.cfg)
}

func (s *Server) runServer() {
	http.HandleFunc("/rpc", s.rpc)
	log.Printf("ListenServer Port: [%s]", s.cfg.Port)
	http.ListenAndServe(s.cfg.Port, nil)
}

func (s *Server) Run() {
	s.init()
	s.runServer()
}