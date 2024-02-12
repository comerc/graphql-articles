package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.43

import (
	"context"
	"errors"

	"github.com/acelot/articles/internal/feature/tag"
	"github.com/acelot/articles/internal/gql/model"
	"go.uber.org/zap"
)

// Update is the resolver for the update field.
func (r *tagMutationResolver) Update(ctx context.Context, obj *model.TagMutation, input model.TagUpdateInput) (model.TagUpdateResult, error) {
	updatedTag, err := r.env.Services.TagService.UpdateTag(input)

	if errors.Is(err, tag.NotFoundError) {
		return model.TagNotFoundProblem{Message: "tag not found"}, nil
	}

	if errors.Is(err, tag.VersionMismatchError) {
		return NewVersionMismatchProblem(), nil
	}

	if err != nil {
		r.env.Logger.Error("tag.Service.UpdateTag", zap.Error(err))

		return NewInternalErrorProblem(), nil
	}

	return model.TagUpdateOk{
		Tag: tag.MapOneToGqlModel(*updatedTag),
	}, nil
}
