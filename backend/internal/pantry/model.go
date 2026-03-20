package pantry

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PantryItem struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"user_id" bson:"user_id"`
	Name      string             `json:"name" bson:"name"`
	Quantity  float64            `json:"quantity" bson:"quantity"`
	Unit      string             `json:"unit" bson:"unit"`
	ImageURL  string             `json:"image_url,omitempty" bson:"image_url,omitempty"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
}
