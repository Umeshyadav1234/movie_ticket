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

// Create a new transaction
func CreateTransaction(c *gin.Context) {
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	transaction.ID = primitive.NewObjectID()
	result, err := db.TransactionCollection.InsertOne(context.Background(), transaction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create transaction"})
		return
	}
	c.JSON(http.StatusOK, result)
}

// Get all transactions
func GetTransactions(c *gin.Context) {
	cursor, err := db.TransactionCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}
	var transactions []models.Transaction
	for cursor.Next(context.Background()) {
		var trans models.Transaction
		if err := cursor.Decode(&trans); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode transaction"})
			return
		}
		transactions = append(transactions, trans)
	}
	c.JSON(http.StatusOK, transactions)
}

// Get transaction by ID
func GetTransactionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	var trans models.Transaction
	err = db.TransactionCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&trans)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}
	c.JSON(http.StatusOK, trans)
}

// Update transaction by ID
func UpdateTransactionByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	var updated models.Transaction
	if err := c.ShouldBindJSON(&updated); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	update := bson.M{"$set": updated}
	_, err = db.TransactionCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update transaction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}

// Delete transaction by ID
func DeleteTransaction(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid transaction ID"})
		return
	}
	_, err = db.TransactionCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete transaction"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Transaction deleted successfully"})
}