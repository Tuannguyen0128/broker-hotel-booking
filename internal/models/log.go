package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type LogData struct {
	ID             primitive.ObjectID `json:"id" bson:"ID,omitempty"`
	RequestDetail  string             `json:"request_detail" bson:"request_detail,omitempty"`
	ResponseDetail string             `json:"response_detail" bson:"response_detail,omitempty"`
	CreatedDate    time.Time          `json:"created_date" bson:"created_date,omitempty"`
}
