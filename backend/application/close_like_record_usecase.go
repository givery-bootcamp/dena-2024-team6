package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type CloseLikeRecordUsecase interface {
	Execute(ctx context.Context) error
}

func NewCloseLikeRecordUsecase(i *do.Injector) (CloseLikeRecordUsecase, error) {
	likeRepo := do.MustInvoke[repository.LikeRepository](i)
	return &CloseLikeRecordUsecaseInteractor{
		likeRepository: likeRepo,
	}, nil
}

type CloseLikeRecordUsecaseInteractor struct {
	likeRepository repository.LikeRepository
}

// Execute implements CloseLikeRecordUsecase.
func (c *CloseLikeRecordUsecaseInteractor) Execute(ctx context.Context) error {
	err := c.likeRepository.Close(ctx)
	if err != nil {
		return apperror.New(apperror.CodeInternalServer, "failed to close like record")
	}
	return nil
}
