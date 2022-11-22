package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/kirankkirankumar/gqlgen-ddk/graph/generated"
	"github.com/kirankkirankumar/gqlgen-ddk/graph/model"
)

// CreateTweet is the resolver for the createTweet field.
func (r *mutationResolver) CreateTweet(ctx context.Context, body *string) (*model.Tweet, error) {

	panic(fmt.Errorf("not implemented: CreateTweet - createTweet"))

}

// DeleteTweet is the resolver for the deleteTweet field.
func (r *mutationResolver) DeleteTweet(ctx context.Context, id int) (*model.Tweet, error) {

	panic(fmt.Errorf("not implemented: DeleteTweet - deleteTweet"))

}

// MarkTweetRead is the resolver for the markTweetRead field.
func (r *mutationResolver) MarkTweetRead(ctx context.Context, id int) (*bool, error) {

	panic(fmt.Errorf("not implemented: MarkTweetRead - markTweetRead"))

}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, id int) (*model.Event, error) {

	panic(fmt.Errorf("not implemented: DeleteEvent - deleteEvent"))

}

// Tweet is the resolver for the Tweet field.
func (r *queryResolver) Tweet(ctx context.Context, id int) (*model.Tweet, error) {

	var data *model.Tweet
	r.Repo.GetData(&data)
	return data, nil

}

// Tweets is the resolver for the Tweets field.
func (r *queryResolver) Tweets(ctx context.Context, limit *int, skip *int, sortField *string, sortOrder *string) ([]*model.Tweet, error) {

	var data []*model.Tweet
	r.Repo.GetData(&data)
	return data, nil

}

// TweetsMeta is the resolver for the TweetsMeta field.
func (r *queryResolver) TweetsMeta(ctx context.Context) (*model.Meta, error) {

	var data *model.Meta
	r.Repo.GetData(&data)
	return data, nil

}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {

	var data *model.User
	r.Repo.GetData(&data)
	return data, nil

}

// Notifications is the resolver for the Notifications field.
func (r *queryResolver) Notifications(ctx context.Context, limit *int) ([]*model.Notification, error) {

	var data []*model.Notification
	r.Repo.GetData(&data)
	return data, nil

}

// NotificationsMeta is the resolver for the NotificationsMeta field.
func (r *queryResolver) NotificationsMeta(ctx context.Context) (*model.Meta, error) {

	var data *model.Meta
	r.Repo.GetData(&data)
	return data, nil

}

// Events is the resolver for the Events field.
func (r *queryResolver) Events(ctx context.Context) (*model.Event, error) {

	var data *model.Event
	r.Repo.GetData(&data)
	return data, nil

}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
