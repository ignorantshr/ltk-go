package main

import (
	"common/config"
	"common/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"server/client"
)

func main() {
	conf := config.GetConfig()
	r := gin.Default()
	client.RegistryRoutes(r)

	// initial
	client.GeneratePools(2, entity.TwoRoom)
	client.MatchPlayers()

	defer client.ShutdownMq()

	err := r.Run(fmt.Sprintf("%s:%d", conf.Server.Host, conf.Server.Port))
	if err != nil {
		return
	}
}
