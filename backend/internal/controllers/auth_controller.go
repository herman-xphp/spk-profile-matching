package controllers

import (
	"net/http"
	"os"
	"time"

	"backend/internal/dto"
	"backend/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type AuthController struct {
	svc *services.AuthService
}

func NewAuthController(s *services.AuthService) *AuthController {
	return &AuthController{svc: s}
}

func (ac *AuthController) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ac.svc.Authenticate(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// 🔐 Buat JWT token
	secret := os.Getenv("SECRET_KEY")
	if secret == "" {
		secret = "defaultsecret" // fallback dev mode
	}

	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // token berlaku 24 jam
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// ✅ Return token + user info (using DTO)
	response := dto.LoginResponse{
		Token: tokenString,
		User:  dto.MapUserToResponse(user),
	}

	c.JSON(http.StatusOK, response)
}
