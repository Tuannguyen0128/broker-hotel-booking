package log

import (
	"broker-hotel-booking/internal/models"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type LogRepository struct {
	db *mongo.Database
}

func New(db *mongo.Database) *LogRepository {
	return &LogRepository{db: db}
}

func (r *LogRepository) InsertRequest(req *models.LogData) *models.ErrInfo {
	log.Println(*req)
	_, err := r.db.Collection("log").InsertOne(context.TODO(), req)

	if err != nil {
		err.Error()
	}

	return nil
}
