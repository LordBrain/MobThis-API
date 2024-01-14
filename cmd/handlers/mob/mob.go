package mob

import (
	"net/http"

	"github.com/LordBrain/MobThis-API/cmd/services/mob"
	"github.com/gin-gonic/gin"
)

func AddMobsV1(router *gin.Engine) {
	router.POST("/v1/mob", func(c *gin.Context) {
		newMob(c)
	})
	router.POST("/v1/mob/:mobsession", func(c *gin.Context) {
		joinMob(c)
	})
	router.GET("/v1/mob/:mobsession", func(c *gin.Context) {
		getMob(c)
	})
	// router.DELETE("/v1/mob/:mobsession", func(c *gin.Context) {
	// 	leaveMob(c)
	// })
	router.POST("/v1/mob/:mobsession/start", func(c *gin.Context) {
		startMob(c)
	})
}

func newMob(c *gin.Context) {

	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()

	var newMob mob.NewMobSession
	err := c.ShouldBindJSON(&newMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}

	mob.MobSession.CreateSession(newMob)

	c.JSON(http.StatusCreated, mob.MobSession)
}

func getMob(c *gin.Context) {
	mobSessionName := c.Param("mobsession")

	if mobSessionName == mob.MobSession.SessionName {
		c.JSON(http.StatusOK, mob.MobSession)
	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}

func joinMob(c *gin.Context) {
	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()
	mobSessionName := c.Param("mobsession")
	var joinMob mob.JoinMob

	err := c.ShouldBindJSON(&joinMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}

	if mobSessionName == mob.MobSession.SessionName {
		err := mob.MobSession.SessionJoin(joinMob)
		if err != nil {
			c.JSON(err.StatusCode(), err.ErrorMessage())
		} else {
			c.JSON(http.StatusOK, mob.MobSession)
		}

	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}

// func leaveMob(c *gin.Context) {
// 	mobSessionName := c.Param("mobsession")
// 	var leaveMob mob.MobSession

// 	err := c.ShouldBindJSON(&leaveMob)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
// 		return
// 	}

// 	leaveMob.SessionName = mobSessionName
// 	// error := leaveMob.SessionLeave(redisClient)

// 	// if error != nil {
// 	// 	c.JSON(error.StatusCode(), error.GetErrorJSON())
// 	// 	return
// 	// }

// 	c.JSON(http.StatusOK, leaveMob)

// }

func startMob(c *gin.Context) {
	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()
	mobSessionName := c.Param("mobsession")
	// var joinMob mob.JoinMob

	// err := c.ShouldBindJSON(&joinMob)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
	// 	return
	// }

	if mobSessionName == mob.MobSession.SessionName {
		mob.MobSession.SessionStart()

		// err := mob.MobSession.SessionJoin(joinMob)
		// if err != nil {
		// 	c.JSON(err.StatusCode(), err.ErrorMessage())
		// } else {
		// 	c.JSON(http.StatusOK, mob.MobSession)
		// }

	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}
