package controller

import (
	"lexcodex/internal/models"
	"lexcodex/internal/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Controller) GetCategoriesController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	categories, err := models.GetCategories(uid)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

func (h *Controller) CreateCategoryController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var newCategory models.Category
	if err := c.ShouldBindJSON(&newCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate main category
	if !utils.IsValidMainCategory(newCategory.MainCategory) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid main category. Supported: Needs, Wants, Savings, Income, Transfer"})
		return
	}

	newCategory.UserID = uid

	category, err := models.AddCategory(newCategory)
	if err != nil {
		log.Printf("Error adding category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add category"})
		return
	}
	c.JSON(http.StatusCreated, category)
}

func (h *Controller) UpdateCategoryController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var updatedCategory models.Category
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedCategory.UserID = uid

	category, err := models.UpdateCategory(updatedCategory)
	if err != nil {
		log.Printf("Error updating category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed updating category"})
		return
	}
	c.JSON(http.StatusOK, category)
}

func (h *Controller) DeleteCategoryController(c *gin.Context) {
	uid, err := utils.GetUserUID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	var deletedCategory models.Category
	if err := c.ShouldBindJSON(&deletedCategory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	deletedCategory.UserID = uid

	if err := models.DeleteCategory(deletedCategory); err != nil {
		log.Printf("Error deleting category: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed deleting category"})
		return
	}

	c.JSON(http.StatusOK, "Category deleted successfully")
}
