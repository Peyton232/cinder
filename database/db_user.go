package database

import (
	"context"
	"log"
	"time"

	"cinderella-meetup/graph/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *DB) CreateUser(input *model.NewUser) *model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	newID := primitive.NewObjectID().Hex()
	user := &model.User{
		UserID:  newID,
		Profile: (*model.Profile)(input.Profile),
	}

	res, err := collection.InsertOne(ctx, user)
	if err != nil || res == nil {
		log.Print(err)
		log.Print("\nunable to insert user into DB in database package\n")
		return nil
	}
	return user
}

func (db *DB) AllUsers() []*model.User {
	collection := db.users
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		log.Print("\nunable to find any users from DB in database package\n")
		return nil
	}
	var users []*model.User
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Print(err)
			log.Print("\nunable to read user model in database package\n")
			return nil
		}
		users = append(users, user)
	}
	return users
}

func (DB *DB) FindUserByID(ID string) *model.User {
	collection := DB.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	collection.FindOne(ctx, bson.M{"userid": ID}).Decode(&user)
	return &user
}

func (DB *DB) MatchWith(userID string, matchesID string) *model.User {
	collection := DB.users
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	user := model.User{}
	// if other person passed, then also pass
	//else
	// move from daily match to pastMatches
}
