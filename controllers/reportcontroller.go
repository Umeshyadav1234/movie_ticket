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

// CreateReport creates a new report document
func CreateReport(c *gin.Context) {
	var report models.Report
	if err := c.ShouldBindJSON(&report); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	report.ID = primitive.NewObjectID()
	result, err := db.ReportCollection.InsertOne(context.Background(), report)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create report"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// GetReports retrieves all report documents
func GetReports(c *gin.Context) {
	cursor, err := db.ReportCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch reports"})
		return
	}

	var reports []models.Report
	for cursor.Next(context.Background()) {
		var rpt models.Report
		if err := cursor.Decode(&rpt); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode report"})
			return
		}
		reports = append(reports, rpt)
	}

	c.JSON(http.StatusOK, reports)
}

// GetReportByID fetches a report by its ID
func GetReportByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}

	var rpt models.Report
	err = db.ReportCollection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&rpt)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Report not found"})
		return
	}

	c.JSON(http.StatusOK, rpt)
}

// UpdateReportByID updates a report by its ID
func UpdateReportByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}

	var updatedReport models.Report
	if err := c.ShouldBindJSON(&updatedReport); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	update := bson.M{"$set": updatedReport}
	_, err = db.ReportCollection.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report updated successfully"})
}

// DeleteReport deletes a report by its ID
func DeleteReport(c *gin.Context) {
	idParam := c.Param("id")
	id, err := primitive.ObjectIDFromHex(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid report ID"})
		return
	}

	_, err = db.ReportCollection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete report"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Report deleted successfully"})
}
