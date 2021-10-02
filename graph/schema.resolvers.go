package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"cinderella-meetup/graph/generated"
	"cinderella-meetup/graph/model"
	"context"
	"fmt"
)

func (r *mutationResolver) CreateNewUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	return db.CreateUser(&input), nil
}

func (r *mutationResolver) FindTodaysMatches(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) MatchWith(ctx context.Context, userid string, matchesID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UnmatchWith(ctx context.Context, userid string, matchesID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) BlockPerson(ctx context.Context, userID string, blockedID string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ChangePreferences(ctx context.Context, userID string, input *model.NewPref) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) SendDailyAnswer(ctx context.Context, userID string, answer string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) AllUsers(ctx context.Context) ([]*model.User, error) {
	return db.AllUsers(), nil
}

func (r *queryResolver) UserByID(ctx context.Context, userID string) (*model.User, error) {
	return db.FindUserByID(userID), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
