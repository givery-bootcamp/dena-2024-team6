//go:generate mockgen -source=get_post_detail_usecase.go -destination=get_post_detail_usecase_mock.go -package=application
package application

import (
	"context"
	apperror "myapp/domain/apperror"
	"myapp/domain/model"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type GetPostDetailUsecase interface {
	Execute(ctx context.Context, input GetPostDetailUsecaseInput) (GetPostDetailUsecaseOutput, error)
}

type GetPostDetailUsecaseInput struct {
	ID int
}

type GetPostDetailUsecaseOutput struct {
	Post model.PostDetail
}

func NewGetPostDetailUsecase(i *do.Injector) (GetPostDetailUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	return &getPostDetailUsecaseInteractor{
		postRepository: postRepo,
	}, nil
}

type getPostDetailUsecaseInteractor struct {
	postRepository repository.PostRepository
}

// Execute implements GetPostDetailUsecase.
func (g getPostDetailUsecaseInteractor) Execute(ctx context.Context, input GetPostDetailUsecaseInput) (GetPostDetailUsecaseOutput, error) {
	post, err := g.postRepository.GetDetail(ctx, input.ID)
	if err != nil {
		return GetPostDetailUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to get post detail")
	}
	if post.IsEmpty() {
		return GetPostDetailUsecaseOutput{}, apperror.New(apperror.CodeNotFound, "post is not found")
	}

	return GetPostDetailUsecaseOutput{
		Post: post,
	}, nil
}
