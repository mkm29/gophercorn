package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/mkm29/gophercorn/graph/generated"
	"github.com/mkm29/gophercorn/graph/model"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input model.NewPost) (*model.Post, error) {
	post := &model.Post{
		Title:  input.Title,
		Body:   input.Body,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID, // fix this line
	}
	r.posts = append(r.posts, post)
	return post, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Address:     input.Address,
		PhoneNumber: input.PhoneNumber,
	}
	r.user = user
	return user, nil
}

func (r *postResolver) User(ctx context.Context, obj *model.Post) (*model.User, error) {
	return &model.User{ID: obj.UserID, Name: "user " + obj.UserID}, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.posts, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return r.users, nil
}

func (r *queryResolver) User(ctx context.Context, userID string) (*model.User, error) {
	return &model.User{ID: userID}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Post returns generated.PostResolver implementation.
func (r *Resolver) Post() generated.PostResolver { return &postResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
