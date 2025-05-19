package controllers

import (
	"MTBS/db"
	"MTBS/models"
	"MTBS/services"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookTicket(c *gin.Context) {
	var ticket models.Ticket
	if err := c.ShouldBindJSON(&ticket); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	
	scheduleID, err := primitive.ObjectIDFromHex(ticket.ScheduleID.Hex())
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid schedule ID"})
		return
	}

	var schedule models.Schedule
	err = db.ScheduleCollection.FindOne(context.TODO(), bson.M{"_id": scheduleID}).Decode(&schedule)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Schedule not found"})
		return
	}

	availableSeats := schedule.Seats - schedule.BookedSeats
	if availableSeats <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No seats available"})
		return
	}


	ticket.ID = primitive.NewObjectID()
	ticket.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	_, err = db.TicketCollection.InsertOne(context.TODO(), ticket)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Booking failed"})
		return
	}

	// Update booked seats in schedule
	seatsToBook := 1

	_, err = db.ScheduleCollection.UpdateOne(
		context.TODO(),
		bson.M{"_id": scheduleID},
		bson.M{"$inc": bson.M{"booked_seats": seatsToBook}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update schedule"})
		return
	}
	report := models.Report{
		ID:            primitive.NewObjectID(),
		TicketID:      ticket.TicketID,
		MovieTitle:    schedule.MovieTitle,
		ShowTime:      schedule.StartTime,
		SeatNumber:    ticket.SeatNumber,
		TicketType:    ticket.TicketType,
		Price:         ticket.Price,
		PaymentStatus: "Completed", // Assuming payment is completed
		AmountPaid:    ticket.AmountPaid,
	}

	InsertedID,ok := services.SaveReport(context.TODO(), report)
	if ok != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save report"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "Ticket booked and Report Created successfully",
		"ticket_id":    ticket.ID.Hex(),
		"seats_booked": seatsToBook,
		"schedule_id":  scheduleID.Hex(),
		"report_id":    InsertedID,

	})
}
