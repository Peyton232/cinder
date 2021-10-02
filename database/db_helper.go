package database

import "cinderella-meetup/graph/model"

func (DB *DB) matches(seeker *model.User, matcher *model.User) bool {
	if contains(seeker.BlockMatches, matcher.UserID) || contains(matcher.BlockMatches, seeker.UserID) {
		return false
	}

	// check preferences

	//else
	return true
}

// Check if str is inside the list s
func contains(s []*string, str string) bool {
	for _, v := range s {
		if *v == str {
			return true
		}
	}

	return false
}

func (db *DB) removeUserElement(list []*model.User, ID string) []*model.User {
	// find index of item in list
	index := -1
	for i, n := range list {
		if ID == n.UserID {
			index = i
		}
	}

	if index != -1 {
		// remove element at that spot
		list = append(list[:index], list[index+1:]...)
	}

	return list
}
