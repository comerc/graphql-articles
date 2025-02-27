package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

// Article is the resolver for the article field.
func (r *mutationResolver) Article(ctx context.Context) (*model.ArticleMutation, error) {
	return &model.ArticleMutation{}, nil
}

// ArticleMutation returns runtime.ArticleMutationResolver implementation.
func (r *Resolver) ArticleMutation() runtime.ArticleMutationResolver {
	return &articleMutationResolver{r}
}

type articleMutationResolver struct{ *Resolver }
