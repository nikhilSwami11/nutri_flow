package pantry

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	collection *mongo.Collection
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		collection: db.Collection("pantry_items"),
	}
}

func (r *Repository) GetAll(userID string) ([]PantryItem, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{"user_id": userID})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var items []PantryItem
	if err := cursor.All(ctx, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func (r *Repository) Create(item *PantryItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item.ID = primitive.NewObjectID()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, item)
	return err
}

func (r *Repository) Update(item *PantryItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	item.UpdatedAt = time.Now()

	after := options.After
	err := r.collection.FindOneAndUpdate(
		ctx,
		bson.M{"_id": item.ID, "user_id": item.UserID},
		bson.M{"$set": bson.M{
			"name":       item.Name,
			"quantity":   item.Quantity,
			"unit":       item.Unit,
			"updated_at": item.UpdatedAt,
		}},
		&options.FindOneAndUpdateOptions{ReturnDocument: &after},
	).Decode(item)

	return err
}

func (r *Repository) Delete(item *PantryItem) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := r.collection.DeleteOne(
		ctx,
		bson.M{"_id": item.ID, "user_id": item.UserID},
	)
	return err
}
