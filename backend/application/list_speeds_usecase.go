package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type ListSpeedsUsecase interface {
	Execute(ctx context.Context) (ListSpeedsUsecaseOutput, error)
}

type ListSpeedsUsecaseOutput struct {
	Speeds []model.Speed
}

func NewListSpeedsUsecase(i *do.Injector) (ListSpeedsUsecase, error) {
	speedRepo := do.MustInvoke[repository.SpeedRepository](i)
	return &listSpeedsUsecaseInteractor{
		speedRepository: speedRepo,
	}, nil
}

type listSpeedsUsecaseInteractor struct {
	speedRepository repository.SpeedRepository
}

// Execute implements ListSpeedsUsecase.
func (l *listSpeedsUsecaseInteractor) Execute(ctx context.Context) (ListSpeedsUsecaseOutput, error) {
	speeds, err := l.speedRepository.List(ctx)
	if err != nil {
		return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get speeds")
	}
	if len(speeds) == 0 {
		return ListSpeedsUsecaseOutput{}, apperror.New(apperror.CodeNotFound, "there is no speed")
	}

	return ListSpeedsUsecaseOutput{
		Speeds: speeds,
	}, nil
}
