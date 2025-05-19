package db

import "go.mongodb.org/mongo-driver/mongo"

var (
	CustomerCollection    *mongo.Collection
	AdminCollection       *mongo.Collection
	TicketCollection      *mongo.Collection
	ReservationCollection *mongo.Collection
	TransactionCollection *mongo.Collection
	ReportCollection      *mongo.Collection
	MovieCollection 	  *mongo.Collection
	ScheduleCollection     *mongo.Collection
)

func Collections() {
	CustomerCollection = DB.Collection("Customer")
	AdminCollection = DB.Collection("Admin")
	TicketCollection = DB.Collection("Ticket")
	ReservationCollection = DB.Collection("Reservation")
	TransactionCollection = DB.Collection("Transaction")
	ReportCollection = DB.Collection("Report")
	MovieCollection=DB.Collection("Movie" )
	ScheduleCollection=DB.Collection("Schedule")

}
