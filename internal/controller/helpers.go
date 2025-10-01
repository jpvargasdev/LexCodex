package controller

import (
	"context"
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Controller) RecalculateAllAccountBalances(c *gin.Context) {
	// get user id
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Start a new context with a timeout for the operation
	ctx, cancel := context.WithTimeout(c.Request.Context(), 10*time.Second)
	defer cancel()

	models.RecalculateAllAccountBalances(ctx, uid)
	c.JSON(http.StatusOK, gin.H{"message": "Account balances recalculated successfully"})
}
