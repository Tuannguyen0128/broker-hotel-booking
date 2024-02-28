package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogData struct {
	ID          primitive.ObjectID `json:"id" bson:"ID,omitempty"`
	Content     string             `json:"content" bson:"content,omitempty"`
	CreatedDate time.Time          `json:"created_date" bson:"created_date,omitempty"`
}
