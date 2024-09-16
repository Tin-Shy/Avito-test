package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tender_service/db"
	"tender_service/models"
)

// GetBidReviews - Обработчик для получения отзывов на предложения
func GetBidReviews(c *gin.Context) {
	tenderId := c.Param("tenderId")
	authorUsername := c.Query("authorUsername")
	organizationId := c.Query("organizationId")

	var reviews []models.Review
	result := db.DB.Where("tender_id = ? AND author_username = ? AND organization_id = ?", tenderId, authorUsername, organizationId).Find(&reviews)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, reviews)
}

// SubmitBidFeedback - Обработчик для добавления отзыва на предложение
func SubmitBidFeedback(c *gin.Context) {
	var feedback models.Review
	if err := c.ShouldBindJSON(&feedback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := db.DB.Create(&feedback)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, feedback)
}
