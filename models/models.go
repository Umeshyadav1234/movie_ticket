package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Customer struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	StudID     int                `json:"stud_id" bson:"stud_id"`
	FName      string             `json:"fname" bson:"fname"`
	LName      string             `json:"lname" bson:"lname"`
	Gender     int                `json:"gender" bson:"gender"`
	Age        int                `json:"age" bson:"age"`
	ContactAdd int                `json:"contact_add" bson:"contact_add"`
	Email      string             `json:"email" bson:"email"`
	CustPass   string             `json:"cust_pass" bson:"cust_pass"`
	Role       string             `json:"Role"  bson:"Role"`
}

type Admin struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	AdminID    int                `json:"admin_id" bson:"admin_id"`
	FName      string             `json:"fname" bson:"fname"`
	LName      string             `json:"lname" bson:"lname"`
	Gender     int                `json:"gender" bson:"gender"`
	Age        int                `json:"age" bson:"age"`
	ContactAdd int                `json:"contact_add" bson:"contact_add"`
	Email      string             `json:"email" bson:"email"`
	AdminPass  string             `json:"admin_pass" bson:"admin_pass"`
	Role       string             `json:"Role"  bson:"Role"`
}

type Ticket struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TicketID     int                `json:"ticketid" bson:"ticketid"`
	CustomerID   primitive.ObjectID `bson:"customer_id" json:"customer_id"`
	CustomerMail string             `bson:"customermail"  json:"customermail"`
	TicketNumber int                `json:"ticket_number" bson:"ticket_number"`
	AccomTime    string             `json:"accom_time" bson:"accom_time"`
	TicketType   string             `json:"ticket_type" bson:"ticket_type"`
	Price        float64            `json:"prize" bson:"prize"`
	SeatNumber   int                `json:"seat_number" bson:"seat_number"`
	CreatedAt    primitive.DateTime `bson:"created_at" json:"created_at"`
	ScheduleID   primitive.ObjectID `json:"schedule_id,omitempty" bson:"schedule_id,omitempty"`
	AmountPaid   float64            `json:"amountpaid"      bson:"amountpaid"`
}

type Reservation struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ResID     int                `json:"res_id" bson:"res_id"`
	TicketID  int                `json:"ticket_id" bson:"ticket_id"`
	CustID    int                `json:"cust_id" bson:"cust_id"`
	AdminID   int                `json:"admin_id" bson:"admin_id"`
	ResDate   int                `json:"res_date" bson:"res_date"`
	Duration  string             `json:"duration" bson:"duration"`
	TimeStart string             `json:"time_start" bson:"time_start"`
	TimeEnd   string             `json:"time_end" bson:"time_end"`
	Payment   string             `json:"payment" bson:"payment"`
}

type Transaction struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	TransID   int                `json:"trans_id" bson:"trans_id"`
	TransName string             `json:"trans_name" bson:"trans_name"`
	TicketID  int                `json:"ticket_id" bson:"ticket_id"`
	ResID     int                `json:"res_id" bson:"res_id"`
	CustID    int                `json:"cust_id" bson:"cust_id"`
	AdminID   int                `json:"admin_id" bson:"admin_id"`
	TransDate string             `json:"trans_date" bson:"trans_date"`
}

type Report struct {
	ID            primitive.ObjectID `json:"report_id" bson:"report_id"`
	TicketID      int                `json:"ticket_id" bson:"ticket_id"`
	MovieTitle    string             `json:"movie_title" bson:"movie_title"`
	ShowTime      string             `json:"show_time" bson:"show_time"`
	SeatNumber    int                `json:"seat_number" bson:"seat_number"`
	TicketType    string             `json:"ticket_type" bson:"ticket_type"`
	Price         float64            `json:"prize" bson:"prize"`
	PaymentStatus string             `json:"payment_status" bson:"payment_status"` 
	AmountPaid    float64            `json:"amount_paid" bson:"amount_paid"`
}

type Movie struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Genre       string             `bson:"genre" json:"genre"`
	Duration    int                `bson:"duration" json:"duration"` // in minutes
	Language    string             `bson:"language" json:"language"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`
	ReleaseDate string             `bson:"releasedate"  json:"releasedate"`
}
type Schedule struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	MovieID     primitive.ObjectID `bson:"movie_id" json:"movie_id"`
	MovieTitle  string             `bson:"movietitle"    json:"movietitle"`
	StartTime   string             `bson:"starttime" json:"starttime"`
	Screen      string             `bson:"screen" json:"screen"`
	Price       float64            `bson:"price" json:"price"` // ticket price
	Seats       int                `bson:"seats" json:"seats"` // total seats
	BookedSeats int                `bson:"BookedSeats" json:"BookedSeats"`
	CreatedAt   primitive.DateTime `bson:"created_at" json:"created_at"`
	UpdatedAt   primitive.DateTime `bson:"updated_at" json:"updated_at"`
}

type Login struct {
	Email    string `json:"email" binding:"required"`
	Role     string `json:"role" binding:"required"`
}
