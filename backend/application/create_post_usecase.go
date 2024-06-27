package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/repository"

	"github.com/samber/do"
)

type CreatePostUsecase interface {
	Execute(ctx context.Context, input CreatePostUsecaseInput) (CreatePostUsecaseOutput, error)
}

type CreatePostUsecaseInput struct {
	UserID int
	Title  string
	Body   string
}

type CreatePostUsecaseOutput struct {
	PostID int
}

func NewCreatePostUsecase(i *do.Injector) (CreatePostUsecase, error) {
	postRepo := do.MustInvoke[repository.PostRepository](i)
	return &createPostUsecaseInteractor{
		postRepository: postRepo,
	}, nil
}

type createPostUsecaseInteractor struct {
	postRepository repository.PostRepository
}

// Execute implements CreatePostUsecase.
func (c *createPostUsecaseInteractor) Execute(ctx context.Context, input CreatePostUsecaseInput) (CreatePostUsecaseOutput, error) {
	postID, err := c.postRepository.Create(ctx, input.UserID, input.Title, input.Body)
	if err != nil {
		return CreatePostUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "新しい投稿の作成に失敗しました。詳細は運営までお問い合わせください。")
	}

	return CreatePostUsecaseOutput{
		PostID: postID,
	}, nil
}
