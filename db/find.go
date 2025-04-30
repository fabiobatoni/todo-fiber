package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func Find(collection string, documents any) error {
	client, ctx := getConnection()
	defer client.Disconnect(ctx)

	c := client.Database(dbname).Collection(collection)

	cursor, err := c.Find(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	return cursor.All(context.Background(), documents)
}
