package controllers

import (
	"MTBS/db"
	"MTBS/models"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Create a new admin
func RegisterAdmin(c *gin.Context) {
	var admin models.Admin
	if err := c.ShouldBindJSON(&admin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	admin.ID = primitive.NewObjectID()
	result, err := db.AdminCollection.InsertOne(context.Background(), admin)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create admin"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get all admins
func GetAdmins(c *gin.Context) {
	cursor, err := db.AdminCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch admins"})
		return
	}
	var admins []models.Admin
	for cursor.Next(context.Background()) {
		var admin models.Admin
		if err := cursor.Decode(&admin); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode admin"})
			return
		}
		admins = append(admins, admin)
	}
	c.JSON(http.StatusOK, admins)
}

// Get admin by ID
func GetAdminByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var admin models.Admin
	err = db.AdminCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&admin)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Admin not found"})
		return
	}
	c.JSON(http.StatusOK, admin)
}

// Update admin by ID
func UpdateAdminByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	var updatedAdmin models.Admin
	if err := c.ShouldBindJSON(&updatedAdmin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": updatedAdmin}
	_, err = db.AdminCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update admin"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin updated successfully"})
}

// Delete admin by ID
func DeleteAdmin(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid admin ID"})
		return
	}
	_, err = db.AdminCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete admin"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Admin deleted successfully"})
}
