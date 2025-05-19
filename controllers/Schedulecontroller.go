package controllers

import (
	"context"
	"net/http"
	"time"

	"MTBS/db"
	"MTBS/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AddSchedule(c *gin.Context) {
	var schedule models.Schedule
	if err := c.ShouldBindJSON(&schedule); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate MovieID exists
	var movie models.Movie
	err := db.MovieCollection.FindOne(context.TODO(), primitive.M{"_id": schedule.MovieID}).Decode(&movie)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Movie ID"})
		return
	}

	schedule.ID = primitive.NewObjectID()
	schedule.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	schedule.UpdatedAt = schedule.CreatedAt

	_, err = db.ScheduleCollection.InsertOne(context.TODO(), schedule)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Schedule creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Schedule added", "schedule": schedule})
}

func GetSchedulesByMovie(c *gin.Context) {
	movieID := c.Param("movie_id")
	objID, err := primitive.ObjectIDFromHex(movieID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid movie ID"})
		return
	}

	cursor, err := db.ScheduleCollection.Find(context.TODO(), primitive.M{"movie_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching schedules"})
		return
	}
	defer cursor.Close(context.TODO())

	var schedules []models.Schedule
	if err := cursor.All(context.TODO(), &schedules); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error parsing schedules"})
		return
	}

	c.JSON(http.StatusOK, schedules)
}
