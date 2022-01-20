package client

import (
	"common/entity"
	"math/rand"
	"time"
)

type Pool struct {
	idPool []int64
	roomType entity.RoomType
	playerPool map[int64]*entity.User
	IsW chan struct{}	// is writing or matching
}

var pools []*Pool

func GeneratePools(num int, roomType entity.RoomType) {
	for i:=0; i<num; i++ {
		pools = append(pools, &Pool{
			idPool: make([]int64, 0),
			roomType: roomType,
			playerPool: make(map[int64]*entity.User),
			IsW: make(chan struct{}, 1),
		})
	}
}

func GetPools() []*Pool {
	return pools
}

func AddUserToPool(user *entity.User) {
	rand.Seed(time.Now().UnixNano())
	toBeAddPool := pools[rand.Intn(len(pools))]
	toBeAddPool.IsW <- struct {}{}
	toBeAddPool.playerPool[user.Id] = user
	toBeAddPool.idPool = append(toBeAddPool.idPool, user.Id)
	<-toBeAddPool.IsW
}
