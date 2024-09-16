package api

import (
    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
    "net/http"
    "tender_service/db"
    "tender_service/models"
)

func Ping(c *gin.Context) {
    c.String(http.StatusOK, "ok")
}

func GetTenders(c *gin.Context) {
    var tenders []models.Tender
    if err := db.DB.Find(&tenders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tenders)
}

func CreateTender(c *gin.Context) {
    var tender models.Tender
    if err := c.BindJSON(&tender); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    tender.ID = uuid.New()
    tender.Status = models.TenderCreated // Обновлено
    tender.Version = 1

    if err := db.DB.Create(&tender).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, tender)
}

func GetUserTenders(c *gin.Context) {
    username := c.Query("username")
    var tenders []models.Tender
    if err := db.DB.Where("creator_username = ?", username).Find(&tenders).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tenders)
}

func GetTenderStatus(c *gin.Context) {
    tenderID := c.Param("tenderId")
    var tender models.Tender
    if err := db.DB.First(&tender, "id = ?", tenderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"status": tender.Status})
}

func UpdateTenderStatus(c *gin.Context) {
    tenderID := c.Param("tenderId")
    var tender models.Tender
    if err := db.DB.First(&tender, "id = ?", tenderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
        return
    }

    var updateStatus struct {
        Status models.TenderStatus `json:"status"`
    }
    if err := c.BindJSON(&updateStatus); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tender.Status = updateStatus.Status
    if err := db.DB.Save(&tender).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tender)
}

func EditTender(c *gin.Context) {
    tenderID := c.Param("tenderId")
    var tender models.Tender
    if err := db.DB.First(&tender, "id = ?", tenderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
        return
    }

    if err := c.BindJSON(&tender); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    tender.Version++
    if err := db.DB.Save(&tender).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, tender)
}

func RollbackTender(c *gin.Context) {
    tenderID := c.Param("tenderId")
    var tender models.Tender
    if err := db.DB.First(&tender, "id = ?", tenderID).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Tender not found"})
        return
    }

    // Логика отката версии
    // Здесь нужно добавить логику для отката версии тендера

    c.JSON(http.StatusOK, tender)
}
