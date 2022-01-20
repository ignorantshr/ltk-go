package main

import (
	"common/entity"
	"github.com/gin-gonic/gin"
	"server/client"
)

func main() {
	r := gin.Default()
	client.RegistryRoutes(r)

	// initial
	client.GeneratePools(2, entity.FiveRoom)
	client.MatchPlayers()

	err := r.Run("127.0.0.1:9001")
	if err != nil {
		return
	}
}
