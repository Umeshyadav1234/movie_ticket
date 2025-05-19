package controllers

import (
	"context"
	"net/http"
	"time"

	"MTBS/db"
	"MTBS/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie.ID = primitive.NewObjectID()
	movie.CreatedAt = primitive.NewDateTimeFromTime(time.Now())
	movie.UpdatedAt = movie.CreatedAt

	_, err := db.MovieCollection.InsertOne(context.TODO(), movie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Movie creation failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Movie added", "movie": movie})
}

func GetAllMovies(c *gin.Context) {
	cursor, err := db.MovieCollection.Find(context.TODO(), primitive.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch movies"})
		return
	}
	defer cursor.Close(context.TODO())

	var movies []models.Movie
	if err := cursor.All(context.TODO(), &movies); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse movies"})
		return
	}

	c.JSON(http.StatusOK, movies)
}
func GetMovieByID(c *gin.Context) {
    id := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var movie models.Movie
    err = db.MovieCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&movie)
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
        return
    }

    c.JSON(http.StatusOK, movie)
}
func UpdateMovieByID(c *gin.Context) {
    id := c.Param("id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var updateData models.Movie
    if err := c.ShouldBindJSON(&updateData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    update := bson.M{
        "$set": bson.M{
            "title":       updateData.Title,
            "description": updateData.Description,
            "duration":    updateData.Duration,
            "genre":       updateData.Genre,
            "language":    updateData.Language,
            "release_date": updateData.ReleaseDate,
        },
    }

    _, err = db.MovieCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update movie"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Movie updated successfully"})
}

