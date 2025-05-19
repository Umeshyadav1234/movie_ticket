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

func RegisterCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return

	}
	customer.ID = primitive.NewObjectID()
	cus, err := db.CustomerCollection.InsertOne(context.Background(), customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create customer"})
		return

	}
	c.JSON(http.StatusOK, cus)

}
func Getcustomers(c *gin.Context) {
	res, err := db.CustomerCollection.Find(context.Background(), bson.M{})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get customers"})
		return

	}
	var customers []models.Customer
	for res.Next(context.Background()) {
		var custom models.Customer
		if err := res.Decode(&custom); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to Decode"})
			return
		}
		customers = append(customers, custom)
	}
	c.JSON(http.StatusOK, customers)
	
}

//get by id
func GetCusbyId(c *gin.Context){
	idparam:=c.Param("id")
	id,err:=primitive.ObjectIDFromHex(idparam)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}
	var cust models.Customer
	err=db.CustomerCollection.FindOne(context.Background(),bson.M{"_id":id}).Decode(&cust)
	if err!=nil{
		c.JSON(http.StatusNotFound, gin.H{"error": "cust not found"})
		return
	}
	c.JSON(http.StatusOK,cust)
	
}

//update by id

func UpadatebyId(c *gin.Context){
	idparam:=c.Param("id")
	id,err:=primitive.ObjectIDFromHex(idparam)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}
	var updatecus models.Customer
	if err:=c.ShouldBindJSON(&updatecus);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

    update := bson.M{
        "$set": updatecus,
    }
	_,err=db.CustomerCollection.UpdateOne(context.Background(),bson.M{"_id":id},update)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"Failed"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer updated successfully"})

}
func DeleteCustomer(c *gin.Context){
	idParam := c.Param("id")
    id, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
	_, err = db.CustomerCollection.DeleteOne(context.Background(),bson.M{"_id": id})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete Customer"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Customer deleted successfully"})
}
