package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/SRsawaguchi/graphql-hanson-server/graph/generated"
	"github.com/SRsawaguchi/graphql-hanson-server/graph/model"
	"github.com/SRsawaguchi/graphql-hanson-server/internal/auth"
	"github.com/SRsawaguchi/graphql-hanson-server/internal/links"
	"github.com/SRsawaguchi/graphql-hanson-server/internal/users"
	"github.com/SRsawaguchi/graphql-hanson-server/pkg/jwt"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return &model.Link{}, fmt.Errorf("access denied")
	}

	link := links.Link{
		Title:   input.Title,
		Address: input.Address,
		User:    user,
	}
	_, err := link.Save(ctx, r.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return &model.Link{
		ID:      strconv.Itoa(int(link.ID)),
		Title:   link.Title,
		Address: link.Address,
		User: &model.User{
			ID:   strconv.Itoa(user.ID),
			Name: user.Username,
		},
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	_, err := user.Create(ctx, r.DB)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	user := users.User{
		Username: input.Username,
		Password: input.Password,
	}
	correct, err := user.Authenticate(ctx, r.DB)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	if !correct {
		err := &users.WrongUsernameOrPasswordError{}
		log.Println(err)
		return "", err
	}

	token, err := jwt.GenerateToken(user.Username)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	username, err := jwt.ParseToken(input.Token)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	token, err := jwt.GenerateToken(username)
	if err != nil {
		log.Println(err.Error())
		return "", err
	}
	return token, nil
}

func (r *queryResolver) Links(ctx context.Context) ([]*model.Link, error) {
	resultLinks, err := links.GetAll(ctx, r.DB)
	if err != nil {
		log.Println(err.Error())
		return []*model.Link{}, err
	}

	links := []*model.Link{}
	for _, link := range resultLinks {
		links = append(links, &model.Link{
			ID:      strconv.Itoa(int(link.ID)),
			Title:   link.Title,
			Address: link.Address,
			User: &model.User{
				ID:   strconv.Itoa(link.User.ID),
				Name: link.User.Username,
			},
		})
	}

	return links, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
