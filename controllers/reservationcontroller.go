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

// Create a new reservation
func CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation.ID = primitive.NewObjectID()
	result, err := db.ReservationCollection.InsertOne(context.Background(), reservation)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reservation"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get all reservations
func GetReservations(c *gin.Context) {
	cursor, err := db.ReservationCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reservations"})
		return
	}
	var reservations []models.Reservation
	for cursor.Next(context.Background()) {
		var res models.Reservation
		if err := cursor.Decode(&res); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode reservation"})
			return
		}
		reservations = append(reservations, res)
	}
	c.JSON(http.StatusOK, reservations)
}

// Get reservation by ID
func GetReservationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}
	var res models.Reservation
	err = db.ReservationCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&res)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Reservation not found"})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Update reservation by ID
func UpdateReservationByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}
	var updated models.Reservation
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": updated}
	_, err = db.ReservationCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update reservation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation updated successfully"})
}

// Delete reservation by ID
func DeleteReservation(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid reservation ID"})
		return
	}
	_, err = db.ReservationCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete reservation"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Reservation deleted successfully"})
}
