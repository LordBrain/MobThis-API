package router

import "github.com/gin-gonic/gin"

func New() (newRouter *gin.Engine) {
	newRouter = gin.Default()
	// r.Use(logging.RequestResponseLogger())

	// gin.SetMode(gin.TestMode)
	newRouter.Use(gin.Recovery())
	newRouter.RedirectTrailingSlash = false

	return
}
