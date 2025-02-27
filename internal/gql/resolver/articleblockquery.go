package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"

	"github.com/acelot/articles/internal/gql/model"
	"github.com/acelot/articles/internal/gql/runtime"
)

// ArticleBlock is the resolver for the articleBlock field.
func (r *queryResolver) ArticleBlock(ctx context.Context) (*model.ArticleBlockQuery, error) {
	return &model.ArticleBlockQuery{}, nil
}

// ArticleBlockQuery returns runtime.ArticleBlockQueryResolver implementation.
func (r *Resolver) ArticleBlockQuery() runtime.ArticleBlockQueryResolver {
	return &articleBlockQueryResolver{r}
}

type articleBlockQueryResolver struct{ *Resolver }
