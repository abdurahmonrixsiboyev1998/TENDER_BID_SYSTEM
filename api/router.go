package api

import (
	"tender_bid_system/api/handler"
	"tender_bid_system/api/middleware"
	"time"

	_ "tender_bid_system/api/docs"

	"github.com/go-redis/redis/v8"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
)

var rdb *redis.Client

func initRedis() {
	rdb = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}

func NewRouter(
	authHandler *handler.UserHandler,
	tenderHandler *handler.TenderHandler,
	bidHandler *handler.BidHandler,
	notiHandler *handler.NotificationHandler,
) *gin.Engine {
	initRedis()
	router := gin.Default()

	// @title Tender Bid System API
	// @version 1.0
	// @description API server for Tender Bid System
	// @host localhost:8080
	// @BasePath /api/v1
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Use(middleware.RateLimitMiddleware(rdb, 5, time.Minute))

	router.GET("/admin", middleware.RoleMiddleware("admin"), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome Admin!",
		})
	})

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)

	router.POST("/tender", middleware.RoleMiddleware("contractor"), tenderHandler.CreateTender)
	router.GET("/tenders", middleware.RoleMiddleware("contractor"), tenderHandler.ListTenders)
	router.PUT("/tender/:id", middleware.RoleMiddleware("contractor"), tenderHandler.UpdateTender)
	router.DELETE("/tender/:id", middleware.RoleMiddleware("contractor"), tenderHandler.DeleteTender)

	router.POST("/tenders/bids", middleware.RoleMiddleware("client"), bidHandler.SubmitBid)
	router.GET("/tenders/bids/:id", middleware.RoleMiddleware("client"), bidHandler.ViewBidsByTenderID)
	router.GET("/tenders/bids/contractor/:id", bidHandler.ViewBidsByContractorID)
	router.GET("/bid", bidHandler.GetBidsByPrice)

	router.POST("/notifications", notiHandler.CreateNotification)

	return router
}
