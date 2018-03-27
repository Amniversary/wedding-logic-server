package main

import (
	"github.com/Amniversary/wedding-logic-server/server"
	"github.com/Amniversary/wedding-logic-server/config"
)

func main() {
	server.NewServer(config.NewConfig()).Run()
}



