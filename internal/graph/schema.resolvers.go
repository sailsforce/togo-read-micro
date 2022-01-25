package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/sailsforce/togo-read-micro/internal/graph/generated"
	"github.com/sailsforce/togo-read-micro/internal/graph/model"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) Users(ctx context.Context, params *model.ListUsersInput) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context, limit *int, sort *string) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
