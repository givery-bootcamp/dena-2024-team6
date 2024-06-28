package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type LikePostUsecase interface {
	Execute(ctx context.Context, input LikePostUsecaseInput) error
}

type LikePostUsecaseInput struct {
	PostID int
	Value  int
}

func NewLikePostUsecase(i *do.Injector) (LikePostUsecase, error) {
	likeRepo := do.MustInvoke[repository.LikeRepository](i)
	return &LikePostUsecaseInteractor{
		likeRepository: likeRepo,
	}, nil
}

type LikePostUsecaseInteractor struct {
	likeRepository repository.LikeRepository
}

// Execute implements LikePostUsecase.
func (c *LikePostUsecaseInteractor) Execute(ctx context.Context, input LikePostUsecaseInput) error {
	likes, err := c.likeRepository.Get(ctx, input.PostID)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to fetch likes record")
	}
	likes += input.Value
	err = c.likeRepository.Update(ctx, input.PostID, likes)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to Update comment")
	}

	return nil
}
