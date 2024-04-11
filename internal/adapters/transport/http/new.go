package http

import (
	"fmt"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	views_handlers "github.com/reonardoleis/views/internal/adapters/transport/http/handlers/views"
	views_domain "github.com/reonardoleis/views/internal/core/domain/views"
)

type Server struct {
	r *gin.Engine
}

func NewServer(viewsUsecase views_domain.ViewUsecase) *Server {
	r := gin.Default()

	router := r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST"},
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin"},
	}))

	viewsHandler := views_handlers.New(viewsUsecase)

	router.POST("/views", viewsHandler.AddView)
	router.GET("/views", viewsHandler.GetViewCount)

	return &Server{r: r}
}

func (s Server) Run() error {
	port := fmt.Sprintf(":%s", os.Getenv("TOOLBOX_PORT"))
	if port == ":" {
		port = ":3000"
	}

	return s.r.Run(port)
}
