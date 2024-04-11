package views_handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	dto "github.com/reonardoleis/views/internal/core/dto/views"
)

func getIP(c *gin.Context) string {
	ip := c.GetHeader("X-Forwarded-For")
	if ip == "" {
		ip = c.ClientIP()
	}

	return ip
}

func (h ViewsHandler) AddView(c *gin.Context) {
	req := new(dto.AddViewRequest)
	if err := c.ShouldBindJSON(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx, cancel := context.WithTimeout(c, time.Second*5)
	defer cancel()

	response, err := h.usecase.AddView(ctx, getIP(c), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
