package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Insert(collection string, data any) (primitive.ObjectID, error) {
	client, ctx := getConnection()
	defer client.Connect(ctx)

	c := client.Database(dbname).Collection(collection)

	resp, err := c.InsertOne(context.Background(), data)

	if err != nil {
		return primitive.NilObjectID, err
	}

	return resp.InsertedID.(primitive.ObjectID), nil
}
