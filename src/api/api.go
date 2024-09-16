package api

import (
    "github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/ping", Ping)
    router.GET("/tenders", GetTenders)
    router.POST("/tenders/new", CreateTender)
    router.GET("/tenders/my", GetUserTenders)
    router.PATCH("/tenders/:tenderId/edit", EditTender)
    router.PUT("/tenders/:tenderId/rollback/:version", RollbackTender)
    router.POST("/bids/new", CreateBid)
    router.GET("/bids/my", GetUserBids)
    router.GET("/bids/:tenderId/list", GetBidsByTender)
    router.PATCH("/bids/:bidId/edit", EditBid)
    router.PUT("/bids/:bidId/rollback/:version", RollbackBid)
    router.GET("/bids/:tenderId/reviews", GetReviewsByTender)
}
