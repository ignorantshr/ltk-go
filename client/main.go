package main

import (
	"bytes"
	"common/entity"
	"encoding/json"
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
	// TODO get id from db
	id := rand.Int()
	user := entity.User{
		Id: int64(id),
		Name: "user" + strconv.Itoa(id),
	}
	byteBody, _ := json.Marshal(user)
	resp, err := http.Post("http://localhost:9001/matchPlayers", "text/json", bytes.NewReader(byteBody))
	if err != nil {
		log.Fatalln("failed to match player: ", err)
		return
	}
	defer resp.Body.Close()

	byteContent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("read resp err: ", err)
	}
	log.Println("get globalmq info", string(byteContent))
}
