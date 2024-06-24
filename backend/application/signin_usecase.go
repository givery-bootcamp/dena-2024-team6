package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"
	"myapp/domain/service"
	"strconv"

	"github.com/samber/do"
)

type SigninUsecase interface {
	Execute(ctx context.Context, input SigninUsecaseInput) (SigninUsecaseOutput, error)
}

type SigninUsecaseInput struct {
	UserName string
	Password string
}

type SigninUsecaseOutput struct {
	User  model.User
	Token string
}

func NewSigninUsecase(i *do.Injector) (SigninUsecase, error) {
	idTokenService := do.MustInvoke[service.IdtokenService](i)
	userRepositry := do.MustInvoke[repository.UserRepository](i)
	return &signinUsecaseInteractor{
		idTokenService: idTokenService,
		userRepository: userRepositry,
	}, nil
}

type signinUsecaseInteractor struct {
	idTokenService service.IdtokenService
	userRepository repository.UserRepository
}

// Execute implements SigninUsecase.
func (s *signinUsecaseInteractor) Execute(ctx context.Context, input SigninUsecaseInput) (SigninUsecaseOutput, error) {
	user, err := s.userRepository.GetByUserNameAndPassword(ctx, input.UserName, input.Password)
	if err != nil {
		return SigninUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to verify credential")
	}
	if user.IsEmpty() {
		return SigninUsecaseOutput{}, apperror.New(apperror.CodeUnauthorized, "unauthorized user")
	}

	// ユーザIDをstringに変換してからJWTを生成する
	strUID := strconv.Itoa(user.ID)
	token, err := s.idTokenService.Generate(strUID)
	if err != nil {
		return SigninUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to generate credentials")
	}

	return SigninUsecaseOutput{
		User:  user,
		Token: token,
	}, nil
}
