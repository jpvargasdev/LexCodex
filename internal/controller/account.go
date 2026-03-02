package controller

import (
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccountsController godoc
// @Summary      Get all accounts
// @Description  Get all accounts for the authenticated user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}   models.Account
// @Failure      401  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /accounts [get]
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

// AddAccountController godoc
// @Summary      Create a new account
// @Description  Create a new financial account for the authenticated user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        account  body      models.Account  true  "Account object"
// @Success      201      {object}  models.Account
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /accounts [post]
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

	// Validate currency
	if !utils.IsValidCurrency(newAccount.Currency) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid currency. Supported: SEK, USD, EUR, GBP, COP, MXN, etc."})
		return
	}

	// Validate account type
	if !utils.IsValidAccountType(newAccount.Type) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid account type. Supported: Bank, Credit Card, Cash, Savings, Investment, Checking, Digital Wallet"})
		return
	}

	newAccount.UserID = uid

	account, err := models.AddAccount(newAccount)
	if err != nil {
		log.Printf("Error adding account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add account"})
		return
	}
	c.JSON(http.StatusCreated, account)
}

// UpdateAccountController godoc
// @Summary      Update an account
// @Description  Update an existing account for the authenticated user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id       path      string          true  "Account ID"
// @Param        account  body      models.Account  true  "Account object"
// @Success      200      {object}  models.Account
// @Failure      400      {object}  map[string]string
// @Failure      401      {object}  map[string]string
// @Failure      500      {object}  map[string]string
// @Router       /accounts/{id} [put]
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

	account, err := models.UpdateAccount(newAccount)
	if err != nil {
		log.Printf("Error updating account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account"})
		return
	}
	c.JSON(http.StatusOK, account)
}

// DeleteAccountController godoc
// @Summary      Delete an account
// @Description  Delete an account for the authenticated user
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id  path      string  true  "Account ID"
// @Success      200 {string}  string  "OK"
// @Failure      401 {object}  map[string]string
// @Failure      500 {object}  map[string]string
// @Router       /accounts/{id} [delete]
func (h *Controller) DeleteAccountController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	idParam := c.Param("id")

	err = models.DeleteAccount(idParam, uid)
	if err != nil {
		log.Printf("Error deleting account: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete account"})
		return
	}
	c.JSON(http.StatusOK, "OK")
}
