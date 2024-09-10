package handler

import (
	"net/http"

	"github.com/Sergio-Na/argus/server/internal/auth"
	"github.com/Sergio-Na/argus/server/internal/database"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	auth *auth.Service
	db   *database.Service
}

func NewHandler(auth *auth.Service, db *database.Service) *Handler {
	return &Handler{
		auth: auth,
		db:   db,
	}
}

func (h *Handler) Home(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Welcome to Argus API!",
	})
}

func (h *Handler) SignUp(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
		return
	}
	
	err := h.auth.SignUp(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.auth.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Sign up successful, but auto-login failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"token":   token.AccessToken,
	})
}

func (h *Handler) SignIn(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.auth.SignIn(req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token.AccessToken})
}

func (h *Handler) GetUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, user)
}
