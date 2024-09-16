// В файле api/bid.go

package api

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "tender_service/db"
    "tender_service/models"
     "github.com/google/uuid"
)

func GetBidsByTender(c *gin.Context) {
    tenderID := c.Param("tenderId")
    var bids []models.Bid
    if err := db.DB.Where("tender_id = ?", tenderID).Find(&bids).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bids)
}

func EditBid(c *gin.Context) {
    bidID := c.Param("bidId")
    var bid models.Bid
    if err := db.DB.First(&bid, "id = ?", bidID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
        return
    }

    if err := c.BindJSON(&bid); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Save(&bid).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bid)
}

func RollbackBid(c *gin.Context) {
    bidID := c.Param("bidId")
    // Логика отката версии
    // Здесь нужно добавить логику для отката версии предложения

    var bid models.Bid
    if err := db.DB.First(&bid, "id = ?", bidID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Bid not found"})
        return
    }

    c.JSON(http.StatusOK, bid)
}

func GetReviewsByTender(c *gin.Context) {
    tenderID := c.Param("tenderId")
    authorUsername := c.Query("authorUsername")
    organizationID := c.Query("organizationId")
    var reviews []models.Review
    if err := db.DB.Where("tender_id = ? AND author_username = ? AND organization_id = ?", tenderID, authorUsername, organizationID).Find(&reviews).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, reviews)
}

func CreateBid(c *gin.Context) {
    var bid models.Bid
    if err := c.BindJSON(&bid); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    bid.ID = uuid.New()
    // Дополнительные проверки и логика, если нужно

    if err := db.DB.Create(&bid).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, bid)
}

func GetUserBids(c *gin.Context) {
    username := c.Query("username")
    var bids []models.Bid
    if err := db.DB.Where("creator_username = ?", username).Find(&bids).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, bids)
}
