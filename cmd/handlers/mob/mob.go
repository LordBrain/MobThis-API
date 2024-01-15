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
	router.DELETE("/v1/mob/:mobsession", func(c *gin.Context) {
		leaveMob(c)
	})
	router.POST("/v1/mob/:mobsession/start", func(c *gin.Context) {
		startMob(c)
	})
	router.POST("/v1/mob/:mobsession/rotate", func(c *gin.Context) {
		rotateMob(c)
	})
	router.POST("/v1/mob/:mobsession/end", func(c *gin.Context) {
		endMob(c)
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

func leaveMob(c *gin.Context) {
	mobSessionName := c.Param("mobsession")
	var leaveMob mob.JoinMob

	err := c.ShouldBindJSON(&leaveMob)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Bad JSON input, could not bind.", "Details": err.Error()})
		return
	}

	if mobSessionName == mob.MobSession.SessionName {
		mob.MobSession.SessionLeave(leaveMob.Mobber)
	}

	c.JSON(http.StatusOK, leaveMob)

}

func startMob(c *gin.Context) {
	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()
	mobSessionName := c.Param("mobsession")

	if mobSessionName == mob.MobSession.SessionName {
		mob.MobSession.SessionStart()

	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}

func rotateMob(c *gin.Context) {
	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()
	mobSessionName := c.Param("mobsession")

	if mobSessionName == mob.MobSession.SessionName {
		mob.MobSession.SessionRotate()

	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}

func endMob(c *gin.Context) {
	mob.Mutex.Lock()
	defer mob.Mutex.Unlock()
	mobSessionName := c.Param("mobsession")

	if mobSessionName == mob.MobSession.SessionName {
		mob.MobSession.SessionEnd()

	} else {
		c.JSON(http.StatusNotFound, "Session not found")
	}

}
