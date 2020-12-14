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
	"github.com/SRsawaguchi/graphql-hanson-server/internal/links"
)

func (r *mutationResolver) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	link := links.Link{
		Title:   input.Title,
		Address: input.Address,
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
	}, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, input model.Login) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	panic(fmt.Errorf("not implemented"))
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
