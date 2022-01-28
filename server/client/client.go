package client

import (
	"common/config"
	"common/entity"
	"common/util/mq"
	"fmt"
	"github.com/apache/rocketmq-client-go/v2"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var globalMq rocketmq.Producer
// TODO get from configuration
var addr string

func init() {
	conf := config.GetConfig()
	addr = fmt.Sprintf("%s:%d", conf.RocketMQ.Host, conf.RocketMQ.Port)
	rand.Seed(time.Now().UnixNano())
	// global mq to notice client's room info
	p, err := mq.NewProducer(addr, "globalMq")
	if err != nil {
		panic("failed to create global mq." + err.Error())
	}
	globalMq = p
}

func ShutdownMq(){
	if globalMq != nil {
		globalMq.Shutdown()
	}
}

func RegistryRoutes(r *gin.Engine) {
	r.POST("/matchPlayers", func(context *gin.Context) {
		var user entity.User
		if err := context.BindJSON(&user); err != nil {
			log.Println("Failed to get user from request: ", err)
		}else {
			user.Host = context.Request.Host
			user.ResponseContext = context
			log.Println("get connect from ", user.Host)
			AddUserToPool(&user)
		}
		//context.JSON(http.StatusOK, dto.MqInfo{
		//	Addr: addr,
		//})
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
		for i := 0; i < pool.roomType.PlayerNumber(); i++ {
			n := rand.Intn(len(pool.idPool))
			users = append(users, pool.playerPool[pool.idPool[n]])
			delete(pool.playerPool, pool.idPool[n])
			pool.idPool = append(pool.idPool[:n], pool.idPool[n+1:]...)
		}
		<-pool.IsW

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
		players[i] = entity.NewPlayer(u, i+1, roles[index], entity.NewZhangFei())
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

	for _, p := range room.Players {
		p.User.ResponseContext.JSON(http.StatusOK, map[string]int64{"room": room.Id})
		err := mq.SendMsg(globalMq, strconv.FormatInt(p.Id, 10), *room)
		if err != nil {
			fmt.Printf("failed to send msg to %d: %s", p.Id, err)
		}
	}

	//jsonBytes, err := json.Marshal(room)
	//if err != nil {
	//	log.Println("failed to json room: ", err)
	//	return
	//}
}
