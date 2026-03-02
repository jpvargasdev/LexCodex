package controller

import (
	"net/http"

	"lexcodex/internal/models"
	"lexcodex/internal/utils"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetIncomesController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountParam := c.Query("account_id")
	limitParam := c.Query("limit")
	offsetParam := c.Query("offset")

	incomes, err := models.GetTransactions(models.TransactionTypeIncome, accountParam, limitParam, offsetParam, uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, incomes)
}
