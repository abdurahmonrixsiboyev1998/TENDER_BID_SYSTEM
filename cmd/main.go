package main

import (
	"database/sql"
	"fmt"
	"log"
	"tender_bid_system/api"
	"tender_bid_system/api/handler"
	"tender_bid_system/config"
	"tender_bid_system/repository"
	"tender_bid_system/service"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.LoadConfig()

	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to the database successfully!")

	userRepo := repository.NewUserRepository(db)
	tenderRepo := repository.NewTenderRepository(db)
	bidRepo := repository.NewBidRepository(db)
	notiRepo := repository.NewNotificationRepository(db)

	userService := service.NewUserService(userRepo)
	tenderService := service.NewTenderService(tenderRepo)
	bidService := service.NewBidService(bidRepo)
	notiService := service.NewNotificationService(notiRepo)

	authHandler := handler.NewUserHandler(userService)
	tenderHandler := handler.NewTenderHandler(tenderService)
	bidHandler := handler.NewBidHandler(bidService, userService)
	notiHandler := handler.NewNotificationHandler(notiService)

	router := api.NewRouter(authHandler, tenderHandler, bidHandler, notiHandler)

	if err := router.Run("localhost:8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
