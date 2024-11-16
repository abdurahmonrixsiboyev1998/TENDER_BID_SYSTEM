package api

import (
	"tender_bid_system/api/handler"

	"github.com/gin-gonic/gin"
)

func NewRouter(
	authHandler *handler.UserHandler,
	tenderHandler *handler.TenderHandler,
	bidHandler *handler.BidHandler,
) *gin.Engine {
	router := gin.Default()

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	router.POST("/tender", tenderHandler.CreateTender)
	router.GET("/tenders", tenderHandler.ListTenders)
	router.PUT("/tender/:id", tenderHandler.UpdateTender)
	router.DELETE("/tender/:id", tenderHandler.DeleteTender)

	router.POST("/tenders/:id/bids", bidHandler.SubmitBid)
	router.GET("/tenders/:id/bids", bidHandler.ViewBidsByTenderID)
	router.GET("/bid", bidHandler.GetBidsByPrice)

	return router
}
