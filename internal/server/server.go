package server

import (
	"github.com/gin-gonic/gin"
	"github.com/reonardoleis/views/internal/router"
)

func NewServer() *gin.Engine {
	r := gin.Default()
	router.SetHandlers(r)

	return r
}
