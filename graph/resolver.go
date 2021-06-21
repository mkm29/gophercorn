package graph

//go:generate go run github.com/99designs/gqlgen
import "github.com/mkm29/gophercorn/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	posts []*model.Post
	user  *model.User
	users []*model.User
}
