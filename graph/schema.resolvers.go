package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cinderella-meetup/graph/generated"
	"cinderella-meetup/graph/model"
	"context"
)

func (r *mutationResolver) CreateNewUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.CreateUser(&input), nil
}

func (r *mutationResolver) ChangeQuestion(ctx context.Context, question *string) (*model.DailyQuestion, error) {
	return db.CreateQuestion(*question), nil
}

func (r *mutationResolver) FindTodaysMatches(ctx context.Context) (bool, error) {
	return db.FindTodaysMatches(), nil
}

func (r *mutationResolver) ClockStrikesMidnight(ctx context.Context) (bool, error) {
	return db.ClockStrikesTwelve(), nil
}

func (r *mutationResolver) MatchWith(ctx context.Context, userid string, matchesID string) (*model.User, error) {
	return db.MatchWith(userid, matchesID), nil
}

func (r *mutationResolver) UnmatchWith(ctx context.Context, userid string, matchesID string) (*model.User, error) {
	return db.UnMatchWith(userid, matchesID), nil
}

func (r *mutationResolver) BlockPerson(ctx context.Context, userID string, blockedID string) (*model.User, error) {
	return db.BlockPerson(userID, blockedID), nil
}

func (r *mutationResolver) ChangePreferences(ctx context.Context, userID string, input *model.NewPref) (*model.User, error) {
	return db.ChangePref(userID, *input), nil
}

func (r *mutationResolver) SendDailyAnswer(ctx context.Context, userID string, answer string) (*model.User, error) {
	return db.SendDailyAnswer(userID, answer), nil
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	return db.AllUsers(), nil
}

func (r *queryResolver) UserByID(ctx context.Context, userID string) (*model.User, error) {
	return db.FindUserByID(userID), nil
}

func (r *queryResolver) WhatisQuestion(ctx context.Context) (*model.DailyQuestion, error) {
	return db.ReadQuestion(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
