package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sailsforce/togo-read-micro/internal/graph/generated"
	"github.com/sailsforce/togo-read-micro/internal/graph/model"
)

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *mutationResolver) UserLogin(ctx context.Context, input model.UserCredentials) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	id := uuid.Must(uuid.NewV4())
	user := &model.User{
		ID:        id.String(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Role:      "USER",
		Email:     input.Email,
		Phone:     input.Phone,
		Name:      input.Name,
		Password:  input.Password,
		UserColor: "000000",
		UserImg:   "https://testimage.com",
	}
	r.users = append(r.users, user)

	return user, nil
}

func (r *queryResolver) Me(ctx context.Context, id *string) (*model.User, error) {
	me := &model.User{}
	for _, v := range r.users {
		if v.ID != *id || v.DeletedAt != nil {
			continue
		}

		me.ID = v.ID
		me.CreatedAt = v.CreatedAt
		me.UpdatedAt = v.CreatedAt
		me.Role = v.Role
		me.Email = v.Email
		me.Phone = v.Phone
		me.Name = v.Name
		me.UserColor = v.UserColor
		me.UserImg = v.UserImg
		me.DogProfiles = v.DogProfiles
	}

	return me, nil
}

func (r *queryResolver) Articles(ctx context.Context, limit *int, sort *string) ([]*model.Article, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }
