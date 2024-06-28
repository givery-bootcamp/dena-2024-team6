package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type GetLikeRecordUsecase interface {
	Execute(ctx context.Context, input GetLikeRecordUsecaseInput) (int, error)
}

type GetLikeRecordUsecaseInput struct {
	PostID int
}

func NewGetLikeRecordUsecase(i *do.Injector) (GetLikeRecordUsecase, error) {
	likeRepo := do.MustInvoke[repository.LikeRepository](i)
	return &GetLikeRecordUsecaseInteractor{
		likeRepository: likeRepo,
	}, nil
}

type GetLikeRecordUsecaseInteractor struct {
	likeRepository repository.LikeRepository
}

// Execute implements GetLikeRecordUsecase.
func (c *GetLikeRecordUsecaseInteractor) Execute(ctx context.Context, input GetLikeRecordUsecaseInput) (int, error) {
	likes, err := c.likeRepository.Get(ctx, input.PostID)
	if err != nil {
		return 0, apperror.New(apperror.CodeInternalServer, "failed to fetch likes record")
	}

	return likes, nil
}
