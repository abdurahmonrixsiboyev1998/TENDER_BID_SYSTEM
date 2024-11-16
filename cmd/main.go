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
)

func main() {
	// Konfiguratsiyani yuklash
	cfg := config.LoadConfig()

	// Bazaga ulanish
	db, err := sql.Open("postgres", fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort,
	))
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Bazani tekshirish
	if err := db.Ping(); err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	log.Println("Connected to the database successfully!")

	// Repositorylarni yaratish
	userRepo := repository.NewUserRepository(db)
	tenderRepo := repository.NewTenderRepository(db)
	bidRepo := repository.NewBidRepository(db)

	// Servislarni yaratish
	userService := service.NewUserService(userRepo)
	tenderService := service.NewTenderService(tenderRepo)
	bidService := service.NewBidService(bidRepo)

	// Handlerlarni yaratish
	authHandler := handler.NewUserHandler(userService)
	tenderHandler := handler.NewTenderHandler(tenderService)
	bidHandler := handler.NewBidHandler(bidService)

	// Routerni konfiguratsiya qilish
	router := api.NewRouter(authHandler, tenderHandler, bidHandler)

	// Serverni ishga tushirish
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
