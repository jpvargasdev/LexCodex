package controller

import (
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetBudgetSummaryController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	startDay := c.Query("start_day")
	endDay := c.Query("end_day")

	budgetSummary, err := models.GetBudgetSummary(startDay, endDay, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, budgetSummary)
}
