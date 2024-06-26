package application

import (
	"context"
	"myapp/domain/apperror"
	"myapp/domain/model"
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
	Post model.Post
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
	post, err := c.postRepository.Create(ctx, input.UserID, input.Title, input.Body)
	if err != nil {
		return CreatePostUsecaseOutput{}, apperror.New(apperror.CodeInternalServer, "failed to create post")
	}

	return CreatePostUsecaseOutput{
		Post: post,
	}, nil
}
