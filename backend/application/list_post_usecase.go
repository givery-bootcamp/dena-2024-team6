//go:generate mockgen -source=list_post_usecase.go -destination=list_post_usecase_mock.go -package=application
package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type ListPostUsecase interface {
	Execute(ctx context.Context) (ListPostUsecaseOutput, error)
}

type ListPostUsecaseOutput struct {
	Posts []model.Post
}

func NewListPostUsecase(i *do.Injector) (ListPostUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	return &listPostUsecaseInteractor{
		postRepository: postRepo,
	}, nil
}

type listPostUsecaseInteractor struct {
	postRepository repository.PostRepository
}

// Execute implements ListPostUsecase.
func (l *listPostUsecaseInteractor) Execute(ctx context.Context) (ListPostUsecaseOutput, error) {
	posts, err := l.postRepository.List(ctx)
	if err != nil {
		return ListPostUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get posts")
	}
	if len(posts) == 0 {
		return ListPostUsecaseOutput{}, apperror.New(apperror.CodeNotFound, "there is no post")
	}

	return ListPostUsecaseOutput{
		Posts: posts,
	}, nil
}
