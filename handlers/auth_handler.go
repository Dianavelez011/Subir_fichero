package handlers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	auth "github.com/gin/authentication"
	"github.com/gin/services" // Reemplaza con tu nombre de módulo
	"github.com/gin/utils"    // Reemplaza con tu nombre de módulo
)

type AuthHandler struct {
    authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
    return &AuthHandler{
        authService: authService,
    }
}

func (h *AuthHandler) SignUp(c *gin.Context) {
    var req auth.SignUpRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if req.Password != req.ConfirmPassword {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Passwords do not match"})
        return
    }

    if err := h.authService.CreateUser(req.Name, req.Email, req.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *AuthHandler) SignIn(c *gin.Context) {
    var req auth.SignInRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := h.authService.ValidateUser(req.Email, req.Password); err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciales inválidas"})
        return
    }

    token, err := utils.GenerateToken(req.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error generando el token"})
        return
    }

    c.JSON(http.StatusOK, auth.JwtResponse{Token: token})
}


func (h *AuthHandler) Logout(c *gin.Context) {
    token := c.GetHeader("Authorization")
    if token == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No token provided"})
        return
    }

    // Remover "Bearer " del token
    token = strings.TrimPrefix(token, "Bearer ")
    
    // Invalidar el token
    utils.InvalidateToken(token)
    
    c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}