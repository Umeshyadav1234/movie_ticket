package controllers

import (
	"MTBS/models"
	"context"
	"net/http"
	"MTBS/db"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ViewAvailableSeatsByMovie(c *gin.Context) {
	movieIDStr := c.Param("movie_id")
	movieID, err := primitive.ObjectIDFromHex(movieIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	cursor, err := db.ScheduleCollection.Find(context.TODO(), bson.M{"movie_id": movieID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
		return
	}
	defer cursor.Close(context.TODO())

	var results []gin.H
	for cursor.Next(context.TODO()) {
		var schedule models.Schedule
		if err := cursor.Decode(&schedule); err == nil {
			availableSeats := schedule.Seats - schedule.BookedSeats
			results = append(results, gin.H{
				"schedule_id":     schedule.ID.Hex(),
				"start_time":      schedule.StartTime,
				"available_seats": availableSeats,
			})
		}
	}

	c.JSON(http.StatusOK, results)
}
