package handler

import (
	"net/http"
	"tender_bid_system/model"
	"tender_bid_system/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// @title Tender Bid System API
// @version 1.0
// @description API server for Tender Bid System
// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Enter the token in the format `Bearer {token}`
// @host localhost:8080
// @BasePath /api/v1

// Register godoc
// @Summary Register a new user
// @Description Registers a new user and sends a verification code to their email
// @Tags User
// @Accept json
// @Produce json
// @Param user body model.User true "User registration details"
// @Success 200 {object} model.User
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /register [post]
func (h *UserHandler) Register(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	createUser, err := h.service.RegisterUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}
	ctx.JSON(http.StatusOK, createUser)
}

// Login godoc
// @Summary Login user
// @Description Authenticate user and returns JWT token
// @Tags User
// @Accept json
// @Produce json
// @Param credentials body model.LoginCredentials true "Login credentials"
// @Success 200 {object} model.LoginResponse
// @Failure 400 {object} string
// @Failure 401 {object} string
// @Router /login [post]
func (h *UserHandler) Login(ctx *gin.Context) {
	var req model.LoginCredentials
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	token, err := h.service.LoginUser(ctx, req.Email, req.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
