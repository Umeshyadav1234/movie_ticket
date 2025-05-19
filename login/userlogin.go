package auth

import (
    "MTBS/db"
    "MTBS/jwt"
    "MTBS/models"
    "context"
    "net/http"

    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
)

func LoginUser(c *gin.Context) {
    var loginInput models.Login
    if err := c.ShouldBindJSON(&loginInput); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Define the filter to find the user by email
    filter := bson.M{"email": loginInput.Email}

    var user models.Customer
    err := db.CustomerCollection.FindOne(context.TODO(), filter).Decode(&user)
    if err != nil {
        if err == mongo.ErrNoDocuments {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
            return
        }
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
        return
    }

    // Verify role
    if user.Role != loginInput.Role {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid role"})
        return
    }


    // Generate JWT token
    token, err := jwt.GenerateToken(user.Email, user.Role)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}
