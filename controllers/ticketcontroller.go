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

// Create a new ticket
func CreateTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ticket.ID = primitive.NewObjectID()
	result, err := db.TicketCollection.InsertOne(context.Background(), ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create ticket"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get all tickets
func GetTickets(c *gin.Context) {
	cursor, err := db.TicketCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tickets"})
		return
	}
	var tickets []models.Ticket
	for cursor.Next(context.Background()) {
		var ticket models.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode ticket"})
			return
		}
		tickets = append(tickets, ticket)
	}
	c.JSON(http.StatusOK, tickets)
}

// Get ticket by ID
func GetTicketByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}
	var ticket models.Ticket
	err = db.TicketCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&ticket)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Ticket not found"})
		return
	}
	c.JSON(http.StatusOK, ticket)
}

// Update ticket by ID
func UpdateTicketByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}
	var updatedTicket models.Ticket
	if err := c.ShouldBindJSON(&updatedTicket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": updatedTicket}
	_, err = db.TicketCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update ticket"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket updated successfully"})
}

// Delete ticket by ID
func DeleteTicket(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ticket ID"})
		return
	}
	_, err = db.TicketCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete ticket"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Ticket deleted successfully"})
}
