package controller

import (
	"context"
	"guilliman/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Controller) RecalculateAllAccountBalances(c *gin.Context) {
	// Start a new context with a timeout for the operation
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	models.RecalculateAllAccountBalances(ctx)
	c.JSON(http.StatusOK, gin.H{"message": "Account balances recalculated successfully"})
}
