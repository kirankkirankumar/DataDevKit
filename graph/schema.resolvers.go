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
func (r *mutationResolver) DeleteTweet(ctx context.Context, id string) (*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: DeleteTweet - deleteTweet"))
}

// MarkTweetRead is the resolver for the markTweetRead field.
func (r *mutationResolver) MarkTweetRead(ctx context.Context, id string) (*bool, error) {
	panic(fmt.Errorf("not implemented: MarkTweetRead - markTweetRead"))
}

// Tweet is the resolver for the Tweet field.
func (r *queryResolver) Tweet(ctx context.Context, id string) (*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: Tweet - Tweet"))
}

// Tweets is the resolver for the Tweets field.
func (r *queryResolver) Tweets(ctx context.Context, limit *int, skip *int, sortField *string, sortOrder *string) ([]*model.Tweet, error) {
	panic(fmt.Errorf("not implemented: Tweets - Tweets"))
}

// TweetsMeta is the resolver for the TweetsMeta field.
func (r *queryResolver) TweetsMeta(ctx context.Context) (*model.Meta, error) {
	panic(fmt.Errorf("not implemented: TweetsMeta - TweetsMeta"))
}

// User is the resolver for the User field.
func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented: User - User"))
}

// Notifications is the resolver for the Notifications field.
func (r *queryResolver) Notifications(ctx context.Context, limit *int) ([]*model.Notification, error) {
	panic(fmt.Errorf("not implemented: Notifications - Notifications"))
}

// NotificationsMeta is the resolver for the NotificationsMeta field.
func (r *queryResolver) NotificationsMeta(ctx context.Context) (*model.Meta, error) {
	panic(fmt.Errorf("not implemented: NotificationsMeta - NotificationsMeta"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//   - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//     it when you're done.
//   - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
