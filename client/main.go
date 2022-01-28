package main

import (
	"bytes"
	"common/config"
	"common/entity"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	conf := config.GetConfig()
	// TODO get id from db
	id := rand.Int()
	user := entity.User{
		Id: int64(id),
		Name: "user" + strconv.Itoa(id),
	}
	byteBody, _ := json.Marshal(user)
	resp, err := http.Post(fmt.Sprintf("http://%s:%d/matchPlayers", conf.Server.Host, conf.Server.Port),
		"text/json", bytes.NewReader(byteBody))
	if err != nil {
		log.Fatalln("failed to match player: ", err)
		return
	}
	defer resp.Body.Close()

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read resp err: ", err)
	}
	fmt.Printf("get response: %s\n", byteContent)

	//var mqInfo dto.MqInfo
	//json.Unmarshal(byteContent, &mqInfo)
	//log.Println("get globalmq info:", mqInfo.Addr)
	//
	//c, err := mq.NewPushConsumer(mqInfo.Addr, "globalMq")
	//if err != nil {
	//	log.Fatalf("failed to create consumer for %s: %s", mqInfo.Addr, err)
	//}
	//rec := make(chan struct{})
	//go func() {
	//	time.Sleep(time.Second*20)
	//	err = mq.Subscribe(c, strconv.FormatInt(user.Id, 10), func(ctx context.Context, ext ...*primitive.MessageExt) (consumer.ConsumeResult, error) {
	//		for _, v := range ext {
	//			fmt.Println("recieve:", v)
	//		}
	//		return consumer.ConsumeSuccess, nil
	//	})
	//	if err != nil {
	//		log.Printf("failed to subscribe: %s\n", err)
	//	}
	//	rec <- struct{}{}
	//}()
	//<- rec
	//defer c.Shutdown()
}
