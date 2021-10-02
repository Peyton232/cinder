package database

import (
	"cinderella-meetup/graph/model"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (DB *DB) FindTodaysMatches() bool {
	collection := DB.users
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		log.Print("\nunable to find any users from DB in database package\n")
		return false
	}

	var users []*model.User
	// create working list
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Print(err)
			log.Print("\nunable to read user model in database package\n")
			return false
		}
		// if user.Preference.FindMatchToday != nil && *user.Preference.FindMatchToday {
		// 	users = append(users, user)
		// }
		users = append(users, user)
	}

	// diminishing list
	for i := len(users); i > 0; {

		itr := users[0]

		for j := 1; j < len(users); j++ {
			if DB.matches(itr, users[j]) {
				// match them
				collection.FindOneAndUpdate(ctx, bson.M{"userid": itr.UserID}, bson.M{"$set": bson.M{"todaysMatch": users[j].UserID}})
				collection.FindOneAndUpdate(ctx, bson.M{"userid": users[j].UserID}, bson.M{"$set": bson.M{"todaysMatch": itr.UserID}})

				// remove from working list
				users = DB.removeUserElement(users, users[j].UserID)
				users = DB.removeUserElement(users, itr.UserID)
				break
			}

			if j == len(users)-1 {
				// no match for u
				users = DB.removeUserElement(users, itr.UserID)
			}
		}

		i = len(users)
	}

	return true
}

func (DB *DB) ClockStrikesTwelve() bool {
	collection := DB.users
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Print(err)
		log.Print("\nunable to find any users from DB in database package\n")
		return false
	}

	// create working list
	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Print(err)
			log.Print("\nunable to read user model in database package\n")
			return false
		}

		collection.FindOneAndUpdate(ctx, bson.M{"userid": user.UserID}, bson.M{"$set": bson.M{"todaysMatch": nil}})
	}

	return true
}
