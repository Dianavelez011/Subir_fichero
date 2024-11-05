package middleware

import (
    "github.com/gin-gonic/gin"
    "strings"
    "github.com/gin/utils"  // Ajusta según tu nombre de módulo
    "net/http"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        bearerToken := strings.Split(authHeader, " ")
        if len(bearerToken) != 2 {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
            c.Abort()
            return
        }

        tokenString := bearerToken[1]

        // Verificar si el token ha sido invalidado
        if utils.IsTokenInvalidated(tokenString) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Token has been invalidated"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        // Guardar el email en el contexto
        c.Set("email", claims["email"])
        c.Next()
    }
}