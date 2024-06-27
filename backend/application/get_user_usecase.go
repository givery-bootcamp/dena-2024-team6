package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type GetUserUsecase interface {
	Execute(ctx context.Context, input GetUserUsecaseInput) (GetUserUsecaseOutput, error)
}

type GetUserUsecaseInput struct {
	ID int
}

type GetUserUsecaseOutput struct {
	User model.User
}

func NewGetUserUsecase(i *do.Injector) (GetUserUsecase, error) {
	userRepository := do.MustInvoke[repository.UserRepository](i)
	return &getUserUsecaseInteractor{
		userRepository: userRepository,
	}, nil
}

type getUserUsecaseInteractor struct {
	userRepository repository.UserRepository
}

// Execute implements GetUserUsecase.
func (g *getUserUsecaseInteractor) Execute(ctx context.Context, input GetUserUsecaseInput) (GetUserUsecaseOutput, error) {
	user, err := g.userRepository.GetByID(ctx, input.ID)
	if err != nil {
		return GetUserUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get user")
	}
	if user.IsEmpty() {
		return GetUserUsecaseOutput{}, apperror.New(apperror.CodeNotFound, "user is not found")
	}

	return GetUserUsecaseOutput{
		User: user,
	}, nil
}
