package client

import (
	"common/entity"
	"common/util"
	"common/warrior"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var globalMq rocketmq.Producer
// TODO get from configuration
var addr string = "127.0.0.1:9876"

func init() {
	rand.Seed(time.Now().UnixNano())
	// global mq to notice client's room info
	globalMq = util.NewProducer(addr, "globalMq")
}

func RegistryRoutes(r *gin.Engine) {
	r.POST("/matchPlayers", func(context *gin.Context) {
		var user entity.User
		if err := context.BindJSON(&user); err != nil {
			log.Println("Failed to get user from request: ", err)
		}else {
			user.Host = context.Request.Host
			log.Println("get connect from ", user.Host)
			AddUserToPool(&user)
		}
		context.JSON(http.StatusOK, addr)
	})
}

// MatchPlayers start match on each pool
func MatchPlayers() {
	for _, p := range pools {
		go matchPlayers(p)
	}
}

func matchPlayers(pool *Pool) {
	for {
		rand.Seed(time.Now().Unix())
		for {
			log.Println("Matching... ", pool.idPool)
			if len(pool.idPool) < pool.roomType.PlayerNumber() {
				time.Sleep(time.Second * 10)
			} else {
				break
			}
		}

		pool.IsW <- struct{}{}
		var users []*entity.User
		for i := 0; i < 5; i++ {
			n := rand.Intn(len(pool.idPool))
			users = append(users, pool.playerPool[pool.idPool[n]])
			delete(pool.playerPool, pool.idPool[n])
			pool.idPool = append(pool.idPool[:n], pool.idPool[n+1:]...)
		}
		<-pool.IsW

		for _, p := range users {
			fmt.Println(p)
		}

		startRoom(users, pool.roomType)
	}
}

func startRoom(users []*entity.User, roomType entity.RoomType) {
	// init a room
	room := &entity.Room{
		Id:      rand.Int63n(1000),
	}
	// init pile
	room.Pile = entity.InitPile(roomType)
	// load players
	players := make(map[int]*entity.Player)

	counts := roomType.PlayerNumber()
	roles := roomType.Roles()
	for i, u := range users {
		index := rand.Intn(counts)
		// TODO fill in the warrior
		i++
		players[i] = entity.NewPlayer(u, i, roles[index], warrior.NewZhangFei())
		// dispatch cards to user
		players[i].DispatchCards(room.Pile.DispatchCards(4, true))
		if index == counts-1 {
			roles = roles[:index]
		}else {
			roles = append(roles[:index], roles[index+1:]...)
		}
		counts--
	}
	room.Players = players

	//jsonBytes, err := json.Marshal(room)
	//if err != nil {
	//	log.Println("failed to json room: ", err)
	//	return
	//}
}
