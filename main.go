package main

import (
	"github.com/Amniversary/wedding-logic-server/server"
	"github.com/Amniversary/wedding-logic-server/config"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	server.NewServer(config.NewConfig()).Run()
}



