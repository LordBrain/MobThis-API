package mob

import (
	"fmt"
	"net/http"

	"github.com/LordBrain/MobThis-API/cmd/services/mob"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func AddMobsV1(router *gin.Engine, redisClient *redis.Client) {
	router.POST("/v1/mob", func(c *gin.Context) {
		newMob(c, redisClient)
	})
	router.POST("/v1/mob/:mobsession", func(c *gin.Context) {
		joinMob(c, redisClient)
	})
	router.GET("/v1/mob/:mobsession", func(c *gin.Context) {
		getMob(c, redisClient)
	})
	router.DELETE("/v1/mob/:mobsession", func(c *gin.Context) {
		leaveMob(c, redisClient)
	})

}

func newMob(c *gin.Context, redisClient *redis.Client) {

	var newMob mob.MobSession
	err := c.ShouldBindJSON(&newMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}
	fmt.Println(newMob)

	error := newMob.CreateSession(redisClient)
	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusCreated, newMob)
}

func getMob(c *gin.Context, redisClient *redis.Client) {
	mobSessionName := c.Param("mobsession")

	var mobSession mob.MobSession
	mobSession.SessionName = mobSessionName
	error := mobSession.SessionState(redisClient)

	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusOK, mobSession)

}

func joinMob(c *gin.Context, redisClient *redis.Client) {
	mobSessionName := c.Param("mobsession")
	var joinMob mob.MobSession

	err := c.ShouldBindJSON(&joinMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}

	joinMob.SessionName = mobSessionName
	error := joinMob.SessionJoin(redisClient)

	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusOK, joinMob)

}

func leaveMob(c *gin.Context, redisClient *redis.Client) {
	mobSessionName := c.Param("mobsession")
	var leaveMob mob.MobSession

	err := c.ShouldBindJSON(&leaveMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}

	leaveMob.SessionName = mobSessionName
	error := leaveMob.SessionLeave(redisClient)

	if error != nil {
		c.JSON(error.StatusCode(), error.GetErrorJSON())
		return
	}

	c.JSON(http.StatusOK, leaveMob)

}
