package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"
	"myapp/domain/service"
	"strconv"
	"unicode/utf8"

	"github.com/samber/do"
)

type SignupUsecase interface {
	Execute(ctx context.Context, input SignupUsecaseInput) (SignupUsecaseOutput, error)
}

type SignupUsecaseInput struct {
	UserName string
	Password string
}

type SignupUsecaseOutput struct {
	User  model.User
	Token string
}

func NewSignupUsecase(i *do.Injector) (SignupUsecase, error) {
	idTokenService := do.MustInvoke[service.IdtokenService](i)
	userRepository := do.MustInvoke[repository.UserRepository](i)
	return &signupUsecaseInteractor{
		idTokenService: idTokenService,
		userRepository: userRepository,
	}, nil
}

type signupUsecaseInteractor struct {
	idTokenService service.IdtokenService
	userRepository repository.UserRepository
}

// Execute implements SignupUsecase.
func (s *signupUsecaseInteractor) Execute(ctx context.Context, input SignupUsecaseInput) (SignupUsecaseOutput, error) {
	// The user name must be 1-13 characters long, and only alphanumeric and numeric characters are allowed.
	if utf8.RuneCountInString(input.UserName) < 1 || utf8.RuneCountInString(input.UserName) > 13 {
		return SignupUsecaseOutput{}, apperror.New(apperror.CodeInvalidArgument, "ユーザ名は1文字以上13文字以下である必要があります")
	}

	// The password must be 12-100 characters long, and code points other than ASCII characters are allowed.
	if len(input.Password) < 12 || len(input.Password) > 100 {
		return SignupUsecaseOutput{}, apperror.New(apperror.CodeInvalidArgument, "パスワードは12文字以上100文字以下である必要があります")
	}
	if !isASCII(input.Password) {
		return SignupUsecaseOutput{}, apperror.New(apperror.CodeInvalidArgument, "パスワードはASCII文字である必要があります")
	}

	exists, err := s.userRepository.Exists(ctx, input.UserName)
	if err != nil {
		return SignupUsecaseOutput{}, err
	}
	if exists {
		return SignupUsecaseOutput{}, apperror.New(apperror.CodeConflict, "ユーザ名が既に存在しています")
	}

	user, err := s.userRepository.Create(ctx, input.UserName, input.Password)
	if err != nil {
		return SignupUsecaseOutput{}, err
	}

	// ユーザIDをstringに変換してからJWTを生成する
	strUID := strconv.Itoa(user.ID)
	token, err := s.idTokenService.Generate(strUID)
	if err != nil {
		return SignupUsecaseOutput{}, err
	}

	return SignupUsecaseOutput{
		User:  user,
		Token: token,
	}, nil
}

func isASCII(s string) bool {
	for _, c := range s {
		if c > 127 {
			return false
		}
	}
	return true
}
