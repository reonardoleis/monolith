package router

import (
	"github.com/gin-gonic/gin"
	"github.com/reonardoleis/views/internal/controllers"
)

func SetHandlers(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
}
