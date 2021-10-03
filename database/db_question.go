package database

import (
	"cinderella-meetup/graph/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (db *DB) CreateQuestion(question string) *model.DailyQuestion {
	collection := db.question
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quest := &model.DailyQuestion{}

	quest.Question = question

	res, err := collection.InsertOne(ctx, quest)
	if err != nil || res == nil {
		log.Print(err)
		log.Print("\nunable to insert user into DB in database package\n")
		return nil
	}
	return quest
}

func (db *DB) ReadQuestion() *model.DailyQuestion {
	collection := db.question
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	quest := &model.DailyQuestion{}

	collection.FindOne(ctx, bson.D{}).Decode(&quest)

	return quest
}
