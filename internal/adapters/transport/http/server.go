package http

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	views_handlers "github.com/reonardoleis/views/internal/adapters/transport/http/handlers/views"
	views_domain "github.com/reonardoleis/views/internal/core/domain/views"
)

type Server struct {
	r *gin.Engine
}

func NewServer(viewsUsecase views_domain.ViewUsecase) *Server {
	r := gin.New()

	r.Use(gin.Recovery())

	if os.Getenv("GIN_MODE") != "release" {
		r.Use(gin.Logger())
	}

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	viewsHandler := views_handlers.New(viewsUsecase)

	r.POST("/views", viewsHandler.AddView)
	r.GET("/views", viewsHandler.GetViewCount)

	return &Server{r: r}
}

func (s Server) Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("TOOLBOX_PORT"))
	if port == ":" {
		port = ":3000"
	}

	return s.r.Run(port)
}
