package controller

import (
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccounts godoc
// @Summary      Get accounts
// @Description  get accounts
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  model.Account
// @Failure      400  {object}  httputil.HTTPError
// @Failure      404  {object}  httputil.HTTPError
// @Failure      500  {object}  httputil.HTTPError
// @Router       /accounts/{id} [get]

func (h *Controller) GetAccountsController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	accountId := c.Param("id")

	accounts, err := models.GetAccounts(accountId, uid) // Fetch accounts from storage
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, accounts)
}

func (h *Controller) AddAccountController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newAccount models.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccount.UserID = uid

	account, err := models.AddAccount(newAccount) // Add account to storage
	if err != nil {
		// You can log the error or return it, depending on your application's needs
		log.Printf("Error adding account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add acccount"})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (h *Controller) UpdateAccountController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var newAccount models.Account
	if err := c.ShouldBindJSON(&newAccount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newAccount.UserID = uid

	account, err := models.UpdateAccount(newAccount) // Add account to storage
	if err != nil {
		// You can log the error or return it, depending on your application's needs
		log.Printf("Error adding account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add acccount"})
		return
	}
	c.JSON(http.StatusCreated, account)
}

func (h *Controller) DeleteAccountController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")

	err = models.DeleteAccount(idParam, uid) // delete account
	if err != nil {
		// You can log the error or return it, depending on your application's needs
		log.Printf("Error adding account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete acccount"})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
