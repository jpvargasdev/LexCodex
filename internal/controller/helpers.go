package controller

import (
	"context"
	"guilliman/internal/models"
	"guilliman/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Controller) RecalculateAllAccountBalances(c *gin.Context) {
	_, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Start a new context with a timeout for the operation
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	models.RecalculateAllAccountBalances(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "Account balances recalculated successfully"})
}
