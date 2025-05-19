package services

import (
	"MTBS/db"
	"MTBS/models"
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func SaveReport(ctx context.Context, report models.Report) (primitive.ObjectID, error) {

	result, err := db.ReportCollection.InsertOne(ctx, report)
	if err != nil {
		return primitive.NilObjectID, err
	}

	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, mongo.ErrInvalidIndexValue

	}

	return insertedID, nil

}
