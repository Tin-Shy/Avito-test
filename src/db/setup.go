package db

import (
    "log"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "tender_service/models" // Импортируем модели
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("tenders.db"), &gorm.Config{})
    if err != nil {
        log.Panic("failed to connect database")
    }

    // Автоматическая миграция (создание таблиц) с использованием моделей из пакета models
    DB.AutoMigrate(&models.Tender{}, &models.Bid{})
}
