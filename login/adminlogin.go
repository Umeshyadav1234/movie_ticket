package auth

import (
	"MTBS/db"
	"MTBS/jwt"
	"MTBS/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func LoginAdmin(c *gin.Context) {
	var loginInput models.Login
	if err := c.ShouldBindJSON(&loginInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	filter := bson.M{"email": loginInput.Email}
	var admin models.Admin
	err := db.AdminCollection.FindOne(context.TODO(), filter).Decode(&admin)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not find document"})
		return
	}

	if admin.Role != loginInput.Role {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
		return
	}
	if admin.Email != loginInput.Email {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email"})
		return
	}

	// Generate JWT token
	token, err := jwt.GenerateToken(admin.Email, admin.Role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
