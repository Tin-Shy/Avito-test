package main

import (
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "tender_service/db"
    "tender_service/api"
)

func main() {
    db.InitDB()

    router := gin.Default()
    api.SetupRoutes(router)

    serverAddress := os.Getenv("SERVER_ADDRESS")
    if serverAddress == "" {
        serverAddress = "0.0.0.0:8080"
    }

    log.Fatal(router.Run(serverAddress))
}
