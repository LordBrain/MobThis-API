package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LordBrain/MobThis-API/cmd/handlers/mob"
	"github.com/LordBrain/MobThis-API/cmd/handlers/up"
	"github.com/LordBrain/MobThis-API/cmd/services/router"
)

var PORT = ""

// var RedisAddr = ""

func init() {
	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
}

func main() {

	router := router.New()

	up.AddUpV1(router)
	mob.AddMobsV1(router)

	log.Printf("Running mobthis-backend on :%s...", PORT)
	err := router.Run(fmt.Sprintf(":%s", PORT))
	if err != nil {
		log.Fatal(err.Error())
	}
}
