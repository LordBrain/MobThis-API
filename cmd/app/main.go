package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LordBrain/MobThis-API/cmd/handlers/mob"
	"github.com/LordBrain/MobThis-API/cmd/handlers/up"
	"github.com/LordBrain/MobThis-API/cmd/services/redis"
	"github.com/LordBrain/MobThis-API/cmd/services/router"
)

var PORT = ""
var RedisAddr = ""

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	RedisAddr = os.Getenv("REDIS_ADDRESS")
	if RedisAddr == "" {
		RedisAddr = "localhost:6379"
	}

}

func main() {
	redisClient := redis.RedisConnection(RedisAddr)
	router := router.New()

	up.AddUpV1(router)
	mob.AddMobsV1(router, redisClient)

	log.Printf("Running mobthis-backend on :%s...", PORT)
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
}
