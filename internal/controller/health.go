package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheckController godoc
// @Summary      Health check
// @Description  Returns OK if the server is running
// @Tags         health
// @Produce      json
// @Success      200  {object}  map[string]string
// @Router       /health [get]
func (h *Controller) HealthCheckController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
